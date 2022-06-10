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
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

// The user provided configuration as to where
// to look for components, the type to detect,
// and other user supplied configuration when
// invoking the redefine app.
type RedefineConfig struct {
	// the base folder from where all components are read
	SrcFolder string `json:"srcFolder"`

	// type of files to include
	Includes []string `json:"includes"`

	// folder from where docs are to be read
	DocsFolder string `json:"docsFolder"`

	// the documentation mode to use
	DocMode string `json:"docMode"`

	// the title to use when emitting the components.json file
	Title string `json:"title"`

	// the path of the library to use from disk when loading in UI
	LibraryPath string `json:"libraryPath"`

	// the published URL of the library to use when loading the UI
	LibraryUrl string `json:"libraryUrl"`

	PackageJson *PackageJson
}

// Extract redefine configuration params using the
// OS arguments and/or redefine.config file present
// in the current folder
func GetRedefineConfig(baseFolder string) *RedefineConfig {
	// check if we have a package.json file in there
	packageJsonFilePath := path.Join(baseFolder, "package.json")
	packageJsonExists := FileExists(packageJsonFilePath)

	// read package.json file
	var packageJson PackageJson
	if packageJsonExists {
		packageJsonFileContents, err := ioutil.ReadFile(packageJsonFilePath)
		if err == nil {
			json.Unmarshal(packageJsonFileContents, &packageJson)
		}
	}

	// check if the path passed is to a folder containing redefine.config.json
	// file or to the place that we need to scan
	configFilePath := path.Join(baseFolder, "redefine.config.json")
	configFile := FileExists(configFilePath)

	// create the app config object
	config := RedefineConfig{
		PackageJson: &packageJson,
	}

	if configFile {
		// read the JSON file and populate the structure
		configFileContents, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			fmt.Println("redefine.config.json file present, unable to read file.")
			return nil
		}

		// unmarshal the file
		fmt.Println("Init using redefine.config.json...")
		json.Unmarshal(configFileContents, &config)

		config.SrcFolder = normalizeFolderPath(baseFolder, config.SrcFolder)   // normalize the base folder path
		config.DocsFolder = normalizeFolderPath(baseFolder, config.DocsFolder) // normalize the docs folder path
	} else {
		// config file does not exist
		// we create default configuration
		config = RedefineConfig{}
	}

	// setup defaults
	// for includes
	if len(config.Includes) == 0 {
		config.Includes = []string{"*.ts", "*.tsx", "*.js", "*.jsx"}
	}

	// for title
	if config.Title == "" {
		config.Title = packageJson.Name

		if config.Title == "" {
			config.Title = filepath.Base(config.SrcFolder)
		}
	}

	// for docs folder
	if config.DocsFolder == "" {
		config.DocsFolder = normalizeFolderPath(baseFolder, "docs")
	}

	// for src folder
	if config.SrcFolder == "" {
		folder := normalizeFolderPath(baseFolder, "src")
		if FileExists(folder) {
			config.SrcFolder = folder
		}
	}
	if config.SrcFolder == "" {
		folder := normalizeFolderPath(baseFolder, "lib")
		if FileExists(folder) {
			config.SrcFolder = folder
		}
	}
	if config.SrcFolder == "" {
		folder := normalizeFolderPath(baseFolder, "packages")
		if FileExists(folder) {
			config.SrcFolder = folder
		} else {
			config.SrcFolder = baseFolder
		}
	}

	// all done
	return &config
}

func normalizeFolderPath(baseFolder string, configPath string) string {
	if configPath == "" {
		return ""
	}

	folder := path.Join(baseFolder, configPath)
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
//
func (config *RedefineConfig) scanFolder() ([]string, error) {
	totalFiles := []string{}

	filepath.Walk(config.SrcFolder, func(path string, fileInfo os.FileInfo, err error) error {
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
		if !isFileIncluded(absPath, config.Includes) {
			return nil
		}

		// add file to list
		totalFiles = append(totalFiles, absPath)
		return nil
	})

	return totalFiles, nil
}
