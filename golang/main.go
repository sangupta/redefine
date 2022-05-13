/**
 * Redefine - UI component documentation
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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	ast "sangupta.com/redefine/ast"
	"sangupta.com/redefine/model"
)

/**
 * The user provided configuration as to where
 * to look for code.
 */
type RedefineConfig struct {
	baseFolder string
	includes   []string
}

/**
 * The main entry point to the application.
 */
func main() {
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
	astMap, syntaxKind := ast.BuildAstForFiles(files)

	// extract components
	components := model.GetComponents(astMap, syntaxKind)

	// generate a component dictionary
	componentFileMap := make(map[string]model.Component, 0)
	for _, component := range components {
		componentFileMap[component.SourcePath+"$"+component.Name] = component
	}

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

	// write the file to disk
	ioutil.WriteFile("components.json", jsonStr, 0644)
}

/**
 * Value object to define how the component JSON
 * should be written to disk and/or served for client
 */
type jsonPayload struct {
	Title      string            `json:"title"`
	Components []model.Component `json:"components"`
}

/**
 * Scan a folder for all files that match a given pattern.
 *
 * @param baseFolder the base location that needs to be scanned
 *
 * @param includes an array of wildcard patterns that select files
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
