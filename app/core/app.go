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
	"path"
	"sort"
	"strings"

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

	// fix source path in components
	if len(components) > 0 {
		baseLen := len(config.SrcFolder)

		for index := range components {
			if strings.HasPrefix(components[index].SourcePath, config.SrcFolder) {
				components[index].SourcePath = components[index].SourcePath[baseLen+1:]
			}
		}
	}

	// add documentation if available to component
	if config.DocsFolder != "" {
		for index := range components {
			// build path to doc file
			docFile := path.Join(config.DocsFolder, components[index].SourcePath, components[index].Name)
			ext := path.Ext(docFile)
			docFile = docFile[0:len(docFile)-len(ext)] + ".md"

			// see if file exists
			if !FileExists(docFile) {
				continue
			}

			// read the doc file
			mdFile, err := ioutil.ReadFile(docFile)
			if err != nil {
				panic(err)
			}

			components[index].Docs = string(mdFile)
		}
	}

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
