/*

Redefine - UI component documentation

MIT License.
Copyright (c) 2022, Sandeep Gupta.

Use of this source code is governed by a MIT style license
that can be found in LICENSE file in the code repository.

*/

package main

func main() {
	config := getRedefineConfig()

	app := RedefineApp{
		config: config,
	}

	app.extractAndWriteComponents()
	// app.printComponentsFromSingleFile("/Users/")
}
