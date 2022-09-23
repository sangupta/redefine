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
	"os"
	"path"
	"sort"
	"strings"

	ast "sangupta.com/redefine/ast"
	"sangupta.com/redefine/model"
)

// Simple struct that acts as a wrapper for various
// ways the application can be invoked
type RedefineApp struct {
	RunMode    string
	Config     *RedefineConfig
	BaseFolder string
}

func (app *RedefineApp) IsBuildMode() bool {
	return strings.EqualFold("build", app.RunMode)
}

func (app *RedefineApp) IsPublishMode() bool {
	return strings.EqualFold("publish", app.RunMode)
}

func (app *RedefineApp) IsServeMode() bool {
	return !(app.IsBuildMode() || app.IsPublishMode())
}

// Value object to define how the component JSON
// should be written to disk and/or served for client
type jsonPayload struct {
	Title       string            `json:"title"`       // the title that is recognized from the package.json file or a user supplied string
	Favicon     string            `json:"favicon"`     // favicon defined in redefine config
	Description string            `json:"description"` // description read from package.json file
	Index       string            `json:"libDocs"`     // index file as defined in docs folder
	Version     string            `json:"version"`     // version read from package.json file
	HomePage    string            `json:"homePage"`    // home page read from package.json file
	Author      PackageAuthor     `json:"author"`      // author read from package.json file
	License     string            `json:"license"`     // license read from package.json file
	Components  []model.Component `json:"components"`  // the extracted components
	CustomCss   string            `json:"customCSS"`   // custom css that needs to be included in page
	Lib         string            `json:"library"`     // the actual component library JS
}

func (app *RedefineApp) ExtractAndWriteComponents() ([]byte, error) {
	config := app.Config

	// scan the base folder for all files present
	files, err := config.scanFolder()
	if err != nil {
		return nil, err
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
		baseLen := len(config.SrcFolder.Root)

		for index := range components {
			if strings.HasPrefix(components[index].SourcePath, config.SrcFolder.Root) {
				components[index].SourcePath = components[index].SourcePath[baseLen+1:]
			}
		}
	}

	// add documentation if available to component
	if config.DocsFolder != nil && config.DocsFolder.Root != "" {
		for index := range components {
			// build path to doc file
			docFile := path.Join(config.DocsFolder.Root, components[index].SourcePath, components[index].Name)
			ext := path.Ext(docFile)
			docFile = docFile[0:len(docFile)-len(ext)] + ".md"

			// see if file exists
			if !FileExists(docFile) {
				continue
			}

			// read the doc file
			mdFile, err := os.ReadFile(docFile)
			if err != nil {
				panic(err)
			}

			components[index].Docs = string(mdFile)
		}
	}

	return app.writeFinalJsonFile(components)
}

// Function responsible to write the final components.json
// file to where it needs to be.
//
// This method also returns the generated JSON string back.
func (app *RedefineApp) writeFinalJsonFile(components []model.Component) ([]byte, error) {
	// basic sanity
	config := app.Config
	pkgJson := config.packageJson
	if pkgJson == nil {
		pj := PackageJson{}
		pkgJson = &pj
	}

	// read index.md file if it exists
	indexMdPath := config.DocsFolder.Index
	var libDocs []byte
	if FileExists(indexMdPath) {
		libDocs, _ = os.ReadFile(indexMdPath)
	}

	// read all custom css files
	var builder strings.Builder
	if len(config.Build.CssFiles) > 0 {
		for _, css := range config.Build.CssFiles {
			fmt.Println("Reading custom CSS file from: " + css)
			cssData, err := os.ReadFile(css)
			if err != nil {
				continue
			}

			builder.WriteString(string(cssData))
			builder.WriteRune('\n')
		}
	}

	// write the JSON file
	payload := jsonPayload{
		Title:       config.Template.Title,
		Favicon:     config.Template.FavIcon,
		Index:       string(libDocs),
		Components:  components,
		Description: pkgJson.Description,
		HomePage:    pkgJson.HomePage,
		Version:     pkgJson.Version,
		Author:      pkgJson.Author,
		License:     pkgJson.License,
		CustomCss:   builder.String(),
		Lib:         config.Build.Lib,
	}

	// create JSON byte array
	jsonStr, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return nil, err
	}

	// find the output folder
	if app.IsBuildMode() {
		// only write the file when we are in build mode
		var outFolder string
		if pkgJson.MainFile != "" {
			outFolder = path.Dir(pkgJson.MainFile)
			outFolder = path.Join(app.BaseFolder, outFolder)
		} else {
			outFolder = app.BaseFolder
		}

		// write the file to disk
		jsonFile := path.Join(outFolder, "components.json")
		fmt.Println("Components JSON written to: " + jsonFile)
		os.WriteFile(jsonFile, jsonStr, 0644)
	}

	return jsonStr, nil
}

func (app *RedefineApp) PrintComponentsFromSingleFile(absoluteFilePath string) {
	files := []string{absoluteFilePath}
	astMap, syntaxKind := ast.BuildAstForFiles(files)
	components := model.GetComponents(astMap, syntaxKind)
	jsonStr, _ := json.MarshalIndent(components, "", "  ")
	fmt.Println(string(jsonStr))
}
