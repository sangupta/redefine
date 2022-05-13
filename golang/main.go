/**
 * Redefine
 *
 * MIT License.
 * Copyright (c) 2022, Sandeep Gupta.
 *
 * Use of this source code is governed by a MIT style license
 * that can be found in LICENSE file in the code repository.
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	ast "sangupta.com/redefine/ast"
	"sangupta.com/redefine/model"
)

type RedefineConfig struct {
	baseFolder string
	includes   []string
}

func main() {
	ts := "/Users/sangupta/git/sangupta/bedrock/src/components/asset/AssetBrowser.tsx"
	list := []string{ts}
	ast.BuildAstForFiles(list)
}

func main1() {
	config := RedefineConfig{
		baseFolder: "/Users/sangupta/git/sangupta/bedrock/src",
		includes:   []string{"*.ts", "*.tsx"},
	}

	// scan the base folder for all files present
	files, err := scanFolder(config.baseFolder, config.includes)
	if err != nil {
		log.Fatal(err)
		return
	}

	// parse AST for each file
	astMap := ast.BuildAstForFiles(files)

	// extract components
	components := model.GetComponents(astMap)

	// generate a component dictionary
	// componentFileMap := make(map[string]model.Component, 0)
	// for _, component := range components {
	// 	componentFileMap[component.SourcePath+"$"+component.Name] = component
	// }

	// write the JSON file
	payload := jsonPayload{
		Title:      "React components",
		Components: components,
	}

	jsonStr, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		log.Fatal(err)
		return
	}

	// start the server as needed
	fmt.Println(jsonStr)
}

type jsonPayload struct {
	Title      string
	Components []model.Component
}

/**
 * Scan a folder for all files that match the pattern given
 */
func scanFolder(baseFolder string, includes []string) ([]string, error) {
	totalFiles := []string{}

	filepath.Walk(baseFolder, func(path string, fileInfo os.FileInfo, err error) error {
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

		if !isFileIncluded(absPath, includes) {
			return nil
		}

		// add file to list
		totalFiles = append(totalFiles, absPath)
		return nil
	})

	return totalFiles, nil
}

/**
 * Check if a file should be included in the list
 * of files to be parsed
 */
func isFileIncluded(path string, includes []string) bool {
	for _, pattern := range includes {
		if isWildCardMatch(path, pattern) {
			return true
		}
	}

	return false
}

/**
 * Check if this is a wildcard match
 */
func isWildCardMatch(str string, pattern string) bool {
	i := 0
	j := 0
	starIndex := -1
	iIndex := -1

	strLen := len(str)
	patternLen := len(pattern)

	for i < strLen {
		if j < patternLen && (string(pattern[j]) == "?" || string(pattern[j]) == string(str[i])) {
			i++
			j++
		} else if j < patternLen && string(pattern[j]) == "*" {
			starIndex = j
			iIndex = i
			j++
		} else if starIndex != -1 {
			j = starIndex + 1
			i = iIndex + 1
			iIndex++
		} else {
			return false
		}
	}

	for j < patternLen && string(pattern[j]) == "*" {
		j++
	}

	return j == patternLen
}
