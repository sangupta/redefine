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
	"time"

	core "sangupta.com/redefine/core"
)

// the generated components.json file represented
// as bytes
var componentsJson []byte

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
		componentsJson = jsonBytes
		serveBuildOverHttp()
	}
}

// This method serves the generated components.json over
// HTTP. Optionally, any built files that are defined
// in package.json (including any folder) are also served
func serveBuildOverHttp() {
	http.HandleFunc("/", httpHandler)
	fmt.Println("Starting HTTP server on http://localhost:1309 ...")
	err := http.ListenAndServe(":1309", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// use basic http handler to serve all files
func httpHandler(writer http.ResponseWriter, request *http.Request) {
	uriPath := request.URL.Path

	if uriPath == "/" {
		uriPath = "/index.html"
	}

	fmt.Println("Serving request: " + uriPath)
	if uriPath == "/components.json" {
		writer.Header().Add("Content-Type", "application/json")
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Methods", "GET")
		writer.Header().Add("Access-Control-Max-Age", "86400")
		writer.WriteHeader(http.StatusOK)
		writer.Write(componentsJson)
		return
	}

	// if the request was not served, find the file as was
	// created in the dist folder for the library

	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte("Not found"))
}

func parseOsArguments() *core.RedefineApp {
	var baseFolder string

	// check for os arguments
	numArgs := len(os.Args)
	runMode := "build"
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
