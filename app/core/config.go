/*

Redefine - UI component documentation

MIT License.
Copyright (c) 2022, Sandeep Gupta.

Use of this source code is governed by a MIT style license
that can be found in LICENSE file in the code repository.

*/

package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// The user provided configuration as to where
// to look for components, the type to detect,
// and other user supplied configuration when
// invoking the redefine app.
type RedefineConfig struct {
	baseFolder string
	includes   []string
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
			os.Exit(0)
			return nil
		}

		baseFolder = cwd
	}

	// check if we have a redefine.json present
	// in the current folder

	config := RedefineConfig{
		baseFolder: baseFolder,
		includes:   []string{"*.ts", "*.tsx", "*.js", "*.jsx"},
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

	filepath.Walk(config.baseFolder, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("Unable to read files from path: " + path)
			return nil
		}

		if fileInfo.IsDir() {
			return nil
		}

		// absFilePath := filepath.Join(path, fileInfo.Name())
		absPath, err := filepath.Abs(path)
		if err != nil {
			log.Fatal(err)
		}

		if !isFileIncluded(absPath, config.includes) {
			return nil
		}

		// add file to list
		totalFiles = append(totalFiles, absPath)
		return nil
	})

	return totalFiles, nil
}
