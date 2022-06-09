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
	BaseFolder string `json:"baseFolder"`

	// type of files to include
	Includes []string `json:"includes"`

	// folder from where docs are to be read
	DocsFolder string `json:"docsFolder"`

	// the title to use when emitting the components.json file
	Title string `json:"title"`

	// the path of the library to use from disk when loading in UI
	LibraryPath string `json:"libraryPath"`

	// the published URL of the library to use when loading the UI
	LibraryUrl string `json:"libraryUrl"`
}

// Extract redefine configuration params using the
// OS arguments and/or redefine.config file present
// in the current folder
func GetRedefineConfig() *RedefineConfig {
	var baseFolder string

	// check for os arguments
	if len(os.Args) > 1 {
		baseFolder = os.Args[1]
	} else {
		// we will pick the current folder
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("No path was specified and error reading current directory")
			return nil
		}

		baseFolder = cwd
	}

	// check if the path passed is to a folder containing redefine.config.json
	// file or to the place that we need to scan
	configFilePath := path.Join(baseFolder, "redefine.config.json")
	configFile := fileExists(configFilePath)

	var config RedefineConfig

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

		// normalize the base folder path
		baseFolder = path.Join(baseFolder, config.BaseFolder)
		baseFolder, _ = filepath.Abs(baseFolder)
		config.BaseFolder = baseFolder // set the resolved base folder

		// normalize the docs folder path
		if config.DocsFolder != "" {
			docsFolder := path.Join(baseFolder, config.DocsFolder)
			docsFolder, _ = filepath.Abs(docsFolder)
			config.DocsFolder = docsFolder // set the resolved base folder
		}

		// check and fix title as needed
		if config.Title == "" {
			config.Title = filepath.Base(config.BaseFolder)
		}
	} else {
		// config file does not exist
		// we create default configuration
		config = RedefineConfig{
			BaseFolder: baseFolder,
			Includes:   []string{"*.ts", "*.tsx", "*.js", "*.jsx"},
		}
	}

	return &config
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

	filepath.Walk(config.BaseFolder, func(path string, fileInfo os.FileInfo, err error) error {
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
