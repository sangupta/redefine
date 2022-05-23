/*

Redefine - UI component documentation

MIT License.
Copyright (c) 2022, Sandeep Gupta.

Use of this source code is governed by a MIT style license
that can be found in LICENSE file in the code repository.

*/

package main

import (
	core "sangupta.com/redefine/core"
)

func main() {
	config := core.GetRedefineConfig()

	app := core.RedefineApp{
		Config: config,
	}

	app.ExtractAndWriteComponents()
	// app.PrintComponentsFromSingleFile("/Users/")
}
