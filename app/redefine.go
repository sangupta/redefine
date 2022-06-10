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
	"os"

	core "sangupta.com/redefine/core"
)

func main() {
	app := parseOsArguments()

	config := core.GetRedefineConfig(app.BaseFolder)

	// `nil` config comes in case when we have an error
	// the error was written to console
	if config == nil {
		return
	}

	// setup config
	app.Config = config

	// run extraction
	app.ExtractAndWriteComponents()
	fmt.Println("Done.")
	fmt.Println()

	// app.PrintComponentsFromSingleFile("/Users/sangupta/git/sangupta/bedrock/src/components/asset/AssetBrowser.tsx")
}

func parseOsArguments() *core.RedefineApp {
	var baseFolder string

	// check for os arguments
	if len(os.Args) > 1 {
		baseFolder = os.Args[1]
	} else {
		// we will pick the current folder
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("No path was specified and error reading current directory")
		}

		baseFolder = cwd
	}

	app := core.RedefineApp{
		BaseFolder: baseFolder,
	}

	return &app
}
