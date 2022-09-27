/*

Redefine - UI component documentation

MIT License.
Copyright (c) 2022, Sandeep Gupta.

Use of this source code is governed by a MIT style license
that can be found in LICENSE file in the code repository.

*/

package core

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/google/uuid"
)

// Structure format for the folder configuration
// Same struct is used for source, docs, dist etc
type ConfigFolder struct {
	Root           string   `json:"root"`           // root folder relative to base folder
	Includes       []string `json:"includes"`       // what files are included
	Index          string   `json:"index"`          // the index file, if applicable
	HasFrontMatter bool     `json:"hasFrontMatter"` // whether the documentation has front matter or not
}

// Attributes that the developer can customize to
// be displayed in the redefine UI
type ConfigTemplate struct {
	Title   string `json:"title"`   // title of the page
	FavIcon string `json:"favicon"` // favicon to be displayed
}

type BuildConfig struct {
	Dist      string   `json:"dist"`          // where json is written during build action
	Publish   string   `json:"publishFolder"` // where final published files are written
	CssFiles  []string `json:"css"`           // css files to load
	FontFiles []string `json:"fonts"`         // the font files to be loaded
	JsFiles   []string `json:"js"`            // the JS files to be loaded before we start the app
	Lib       string   `json:"lib"`           // the actual component library to be used
}

// The user provided configuration as to where
// to look for components, the type to detect,
// and other user supplied configuration when
// invoking the redefine app.
type RedefineConfig struct {
	baseFolder  string            // the folder where redefine was run
	packageJson *PackageJson      // the final package json that is read
	libraryMap  map[string]string // map which stores the final library paths
	SrcFolder   *ConfigFolder     `json:"src"`      // the base folder from where all components are read
	DocsFolder  *ConfigFolder     `json:"docs"`     // folder from where docs are to be read
	Build       *BuildConfig      `json:"build"`    // folder where output is written
	Template    *ConfigTemplate   `json:"template"` // template configuration for view page
}

func (config *RedefineConfig) HasLibraryFile(id string) bool {
	_, exists := config.libraryMap[id]

	return exists
}

func (config *RedefineConfig) GetLibraryBytes(id string) []byte {
	if id == "" {
		return nil
	}

	val := config.libraryMap[id]
	if val == "" {
		return nil
	}

	bytes, err := os.ReadFile(val)
	if err != nil {
		return nil
	}

	return bytes
}

// Extract redefine configuration params using the
// OS arguments and/or redefine.config file present
// in the current folder
func GetRedefineConfig(baseFolder string) *RedefineConfig {
	// check if we have a package.json file in there
	packageJsonFilePath := path.Join(baseFolder, "package.json")
	fmt.Println("Reading package.json from: " + packageJsonFilePath)
	packageJsonExists := FileExists(packageJsonFilePath)

	// this is where we store all our configuration
	var config *RedefineConfig

	// read package.json file
	var packageJson PackageJson
	if packageJsonExists {
		packageJsonFileContents, err := os.ReadFile(packageJsonFilePath)

		// error is eaten as we can work on defaults
		if err == nil {
			json.Unmarshal(packageJsonFileContents, &packageJson)

			// read redefine configuration from here
			config = packageJson.Redefine
		}
	}

	// if nothing is present in package.json file
	// let's check if we have `redefine.config.json` file
	if config == nil {
		config = readRedefineConfig(baseFolder)
	}

	// will this happen?
	if config == nil {
		fmt.Println("No configuration available, will use defaults")
		config = &RedefineConfig{}
	}

	// setup base folder
	config.packageJson = &packageJson
	config.baseFolder = baseFolder

	// normalize configuration
	normalizeConfiguration(config, &packageJson)

	// all done
	return config
}

// Read `redefine.config.json` from the given folder, if present.
// If file is not present, return a simple default structure so
// that we can populate same with package.json and other sensible
// defaults.
func readRedefineConfig(baseFolder string) *RedefineConfig {
	fmt.Println("No redefine config was found inside package.json, looking for redefine.config.json...")
	config := RedefineConfig{}

	// check if the path passed is to a folder containing redefine.config.json
	// file or to the place that we need to scan
	configFilePath := path.Join(baseFolder, "redefine.config.json")
	configFile := FileExists(configFilePath)

	if configFile {
		fmt.Println("Found redefine.config.json at: " + configFilePath)
		// read the JSON file and populate the structure
		configFileContents, err := os.ReadFile(configFilePath)
		if err != nil {
			fmt.Println("redefine.config.json file present, unable to read file.")
			return nil
		}

		// unmarshal the file
		fmt.Println("Init using redefine.config.json...")
		json.Unmarshal(configFileContents, &config)
	}

	return &config
}

// this method normalizes configuration based
// on values specified in the package.json or redefine.config.json
// TL;DR: setup defaults
func normalizeConfiguration(config *RedefineConfig, packageJson *PackageJson) {
	// -----------------------------------------------
	// details about src folder
	if config.SrcFolder == nil {
		config.SrcFolder = &ConfigFolder{}
	}

	// normalize the base folder path
	if config.SrcFolder.Root == "" {
		// see if we have 'src', 'lib', or 'packages' folder existing
		// if yes, we will default to it
		if FileExists(config.NormalizeFolderPath("src")) {
			config.SrcFolder.Root = "src"
		} else if FileExists(config.NormalizeFolderPath("lib")) {
			config.SrcFolder.Root = "src"
		} else if FileExists(config.NormalizeFolderPath("packages")) {
			config.SrcFolder.Root = "src"
		}
	}
	config.SrcFolder.Root = config.NormalizeFolderPath(config.SrcFolder.Root)

	// for includes
	if len(config.SrcFolder.Includes) == 0 {
		config.SrcFolder.Includes = []string{"*.ts", "*.tsx", "*.js", "*.jsx"}
	}

	// -----------------------------------------------
	// normalize the docs folder path
	if config.DocsFolder == nil {
		config.DocsFolder = &ConfigFolder{
			HasFrontMatter: true,
		}
	}
	if config.DocsFolder.Root == "" {
		config.DocsFolder.Root = "docs"
	}

	// base path
	config.DocsFolder.Root = config.NormalizeFolderPath(config.DocsFolder.Root)

	// the index file
	if config.DocsFolder.Index == "" {
		config.DocsFolder.Index = "index.md"
	}

	// what constitutes documentation file
	if len(config.DocsFolder.Includes) == 0 {
		config.DocsFolder.Includes = []string{"*.md"}
	}

	// -----------------------------------------------
	// normalize build folder
	if config.Build == nil {
		config.Build = &BuildConfig{}
	}

	if config.Build.Dist == "" {
		config.Build.Dist = "dist"
	}
	config.Build.Dist = config.NormalizeFolderPath(config.Build.Dist)

	if config.Build.Publish == "" {
		config.Build.Publish = "publish"
	}
	config.Build.Publish = config.NormalizeFolderPath(config.Build.Publish)

	// normalize css files
	if config.Build.CssFiles == nil {
		config.Build.CssFiles = []string{}
	}
	if len(config.Build.CssFiles) > 0 {
		for index, css := range config.Build.CssFiles {
			config.Build.CssFiles[index] = config.NormalizeFolderPath(css)
		}
	}

	// normalize js files
	// JS file paths need not be normalized
	if config.Build.JsFiles == nil {
		config.Build.JsFiles = []string{}
	}

	// normalize the library by replacing the JS name with a UUID
	if config.Build.Lib == "" {
		config.Build.Lib = packageJson.MainFile
	}
	if config.Build.Lib != "" {
		config.libraryMap = make(map[string]string)
		uuid := uuid.New().String() + ".js"

		fullJsPath := config.NormalizeFolderPath(config.Build.Lib)
		config.libraryMap[uuid] = fullJsPath
		config.Build.Lib = uuid

		// let's also add the JS map file
		fullMapPath := fullJsPath + ".map"
		if FileExists(fullMapPath) {
			config.libraryMap[filepath.Base(fullMapPath)] = fullMapPath
		}
	}

	// -----------------------------------------------
	// normalize template details
	if config.Template == nil {
		config.Template = &ConfigTemplate{}
	}

	if config.Template.Title == "" {
		config.Template.Title = packageJson.Name

		if config.Template.Title == "" {
			config.Template.Title = filepath.Base(config.SrcFolder.Root)
		}
	}
}

// Function to normalize a folder path by prefixing
// the base folder path and joining if with the given
// child path
func (config *RedefineConfig) NormalizeFolderPath(childPath string) string {
	if childPath == "" {
		return config.baseFolder
	}

	folder := path.Join(config.baseFolder, childPath)
	folder, _ = filepath.Abs(folder)
	return folder
}

// Scan a folder for all files that match a given pattern.
// Returns an array of absolute paths, along with an `error`
// object which is `nil` if successful.
//
// @param baseFolder the base location that needs to be scanned
//
// @param includes an array of wildcard patterns that select files
func (config *RedefineConfig) scanFolder() ([]string, error) {
	totalFiles := []string{}

	filepath.Walk(config.SrcFolder.Root, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("Unable to read files from path: " + path)
			return nil
		}

		// skip if this is a folder
		if fileInfo.IsDir() {
			return nil
		}

		// get absolute path for path
		absPath, err := filepath.Abs(path)
		if err != nil {
			log.Fatal(err)
		}

		// check if the file is included or not
		if !isFileIncluded(absPath, config.SrcFolder.Includes) {
			return nil
		}

		// add file to list
		totalFiles = append(totalFiles, absPath)
		return nil
	})

	return totalFiles, nil
}

// Simple debug function to print information
// regarding what is being used
func (config *RedefineConfig) PrintInfo() {
	fmt.Println("\nUsing following configuration:")
	fmt.Println("    Src folder: " + config.SrcFolder.Root)
	fmt.Printf("    Src includes: %v\n", config.SrcFolder.Includes)
	fmt.Println("    Lib file: " + config.Build.Lib)
	fmt.Println("    Docs folder: " + config.DocsFolder.Root)
	fmt.Println("    Docs index: " + config.DocsFolder.Index)
	fmt.Println()
}
