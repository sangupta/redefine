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
	"sort"

	ast "sangupta.com/redefine/ast"
	"sangupta.com/redefine/model"
)

// Simple struct that acts as a wrapper for various
// ways the application can be invoked
type RedefineApp struct {
	Config *RedefineConfig
}

// Value object to define how the component JSON
// should be written to disk and/or served for client
type jsonPayload struct {

	// the title that is recognized from the package.json
	// file or a user supplied string
	Title string `json:"title"`

	// the extracted components
	Components []model.Component `json:"components"`
}

func (app *RedefineApp) ExtractAndWriteComponents() {
	config := app.Config

	// scan the base folder for all files present
	files, err := config.scanFolder()
	if err != nil {
		log.Fatal(err)
		return
	}

	// parse AST for each file
	astMap, syntaxKind := ast.BuildAstForFiles(files)

	// extract components
	components := model.GetComponents(astMap, syntaxKind)

	// sort components
	sort.SliceStable(components, func(i, j int) bool {
		return components[i].Name < components[j].Name
	})

	// write the JSON file
	payload := jsonPayload{
		Title:      config.Title,
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

func (app *RedefineApp) PrintComponentsFromSingleFile(absoluteFilePath string) {
	files := []string{absoluteFilePath}
	astMap, syntaxKind := ast.BuildAstForFiles(files)
	components := model.GetComponents(astMap, syntaxKind)
	jsonStr, _ := json.MarshalIndent(components, "", "  ")
	fmt.Println(string(jsonStr))
}
