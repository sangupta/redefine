/*

Redefine - UI component documentation

MIT License.
Copyright (c) 2022, Sandeep Gupta.

Use of this source code is governed by a MIT style license
that can be found in LICENSE file in the code repository.

*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	core "sangupta.com/redefine/core"
)

func main() {
	if len(os.Args) == 1 {
		printHelp()
		return
	}

	// parse OS arguments to fetch action and path
	app := parseOsArguments()

	// read configuration, measuring overall time
	start := time.Now()
	config := core.GetRedefineConfig(app.BaseFolder)

	// `nil` config comes in case when we have an error
	// the error was written to console
	if config == nil {
		return
	}

	// setup config
	app.Config = config

	// print all configuration
	config.PrintInfo()

	// run extraction
	jsonBytes, err := app.ExtractAndWriteComponents()
	duration := time.Since(start)

	// error?
	if err != nil {
		fmt.Println("Ran into issues when extracting components")
		log.Fatal(err)
		return
	}

	// emit time taken in generation
	fmt.Println("Done in " + duration.String())
	fmt.Println()

	// if there was nothing produced, exit quietly
	if jsonBytes == nil {
		return
	}

	// if we are in serve mode, start HTTP server
	if app.IsServeMode() {
		serveBuildOverHttp(jsonBytes, config)
		return
	}

	if app.IsPublishMode() {
		publishApplication(jsonBytes, config)
		return
	}
}

// Publish a static application that can be deployed that
// contains everything this application will need. All static
// files, including components.json, are emitted to disk and
// relatively linked to the generated index.html file.
func publishApplication(jsonBytes []byte, config *core.RedefineConfig) {

}

// This method serves the generated components.json over
// HTTP. Optionally, any built files that are defined
// in package.json (including any folder) are also served
func serveBuildOverHttp(jsonBytes []byte, config *core.RedefineConfig) {
	scanFolders := mapset.NewSet[string]()

	for _, css := range config.Build.CssFiles {
		scanFolders.Add(filepath.Dir(css))
	}
	for _, js := range config.Build.JsFiles {
		scanFolders.Add(filepath.Dir(js))
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		doHttpRequest(writer, request, jsonBytes, config, scanFolders)
	})

	fmt.Println("Starting HTTP server on http://localhost:1309 ...")
	err := http.ListenAndServe(":1309", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sendFile(uri string, writer http.ResponseWriter, bytes []byte) {
	if strings.HasSuffix(uri, ".js") {
		writer.Header().Add("Content-Type", "text/javascript")

		// strip off any import statements that contain '.css";'
		content := string(bytes)
		lines := strings.Split(content, "\n")

		for index, line := range lines {
			if strings.HasPrefix(line, "import \"") && strings.HasSuffix(line, ".css\";") {
				lines[index] = ""
			}
		}

		content = strings.Join(lines, "\n")
		bytes = []byte(content)
	}

	if strings.HasSuffix(uri, ".css") {
		writer.Header().Add("Content-Type", "text/css")
	}

	if strings.HasSuffix(uri, ".js.map") || strings.HasSuffix(uri, ".css.map") || strings.HasSuffix(uri, ".json") {
		writer.Header().Add("Content-Type", "application/json")
	}

	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Methods", "GET")
	writer.Header().Add("Access-Control-Max-Age", "86400")
	writer.WriteHeader(http.StatusOK)
	writer.Write(bytes)
}

// use basic http handler to serve all files
func doHttpRequest(writer http.ResponseWriter, request *http.Request, jsonBytes []byte, config *core.RedefineConfig, scanFolders mapset.Set[string]) {
	uriPath := request.URL.Path

	if uriPath == "/" {
		uriPath = "/index.html"
	}

	fmt.Println("Serving request: " + uriPath)
	if uriPath == "/components.json" {
		sendFile(uriPath, writer, jsonBytes)
		return
	}

	// if the request was not served, find the file as was
	// created in the dist folder for the library
	uriNoSlash := uriPath[1:]
	if config.HasLibraryFile(uriNoSlash) {
		jsFile := config.GetLibraryBytes(uriNoSlash)
		if jsFile == nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("No such library found"))
			return
		}

		// file was read
		sendFile(uriPath, writer, jsFile)
		return
	}

	// find if the file is from the JS files
	localFile := config.NormalizeFolderPath(uriNoSlash)
	if core.FileExists(localFile) {
		fileContents, err := os.ReadFile(localFile)
		if err != nil {
			panic(err)
		}

		sendFile(uriPath, writer, fileContents)
		return
	}

	// nothing was found
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte("Not found"))
}

func parseOsArguments() *core.RedefineApp {
	var baseFolder string

	// check for os arguments
	numArgs := len(os.Args)
	runMode := "serve"
	switch numArgs {
	case 1:
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("No path was specified and error reading current directory")
		}

		baseFolder = cwd

	case 2:
		baseFolder = os.Args[1]

	case 3:
		runMode = os.Args[1]
		baseFolder = os.Args[2]

	default:
		return nil
	}

	app := core.RedefineApp{
		RunMode:    runMode,
		BaseFolder: baseFolder,
	}

	return &app
}

func printHelp() {
	fmt.Println("Redefine: UI component documentation")
	fmt.Println("usage: $ redefine <action> <folder>")
	fmt.Println()
	fmt.Println("    <action>  (optional) specify non-default actions:")
	fmt.Println("              `serve`: run local server to serve documentation")
	fmt.Println("              `build`: export all doc files to an output folder")
	fmt.Println()
	fmt.Println("    <folder>  Root folder where either `package.json` or")
	fmt.Println("              `redefine.config.json` exists.")
	fmt.Println()
	fmt.Println("Detailed instructions at https://redefine.sangupta.com")
	fmt.Println()
}
