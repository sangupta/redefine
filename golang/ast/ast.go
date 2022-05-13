/**
 * Redefine
 *
 * MIT License.
 * Copyright (c) 2022, Sandeep Gupta.
 *
 * Use of this source code is governed by a MIT style license
 * that can be found in LICENSE file in the code repository.
 */

package ast

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	stdruntime "runtime"
	"time"

	"github.com/quickjs-go/quickjs-go"
)

/**
 * Create a map of AST's by parsing each file
 */
func BuildAstForFiles(files []string) map[string]SourceFile {
	start := time.Now()
	astMap := make(map[string]SourceFile, len(files))

	doWork := func(context *quickjs.Context, globals quickjs.Value, parseCode quickjs.Value, syntaxKind SyntaxKind) {
		for _, file := range files {
			sourceFile := parseSingleFile(file, context, globals, parseCode, syntaxKind)
			if sourceFile != nil {
				astMap[file] = *sourceFile
			}
		}
	}

	// do job
	runInQuickJS(doWork)

	// get time spent
	duration := time.Since(start)
	fmt.Println("Total time in parsing files: " + duration.String())

	return astMap
}

/**
 * Parse a single file by reading it from disk
 */
func parseSingleFile(file string, context *quickjs.Context, globals quickjs.Value, parseCode quickjs.Value, syntaxKind SyntaxKind) *SourceFile {
	// read the source code file from disk
	sourceCode, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	// create argument list to call the method
	args := make([]quickjs.Value, 4)
	args[0] = context.String("index.ts")
	args[1] = context.String(string(sourceCode))
	args[2] = context.String("")
	args[3] = context.Bool(true)

	// invoke the "createSourceFile" method
	result, err := context.Call(globals, parseCode, args)
	// defer result.Free()

	// check for error, and free up the result
	check(err)

	// now convert the "result" represented as AST in QJS objects
	// to the pure objects that we require
	fmt.Println("Fetching object for " + file)

	typescript := Typescript{
		syntaxKind: syntaxKind,
	}

	sourceFile := typescript.getSourceFile(result)
	return sourceFile
}

func runInQuickJS(doWork func(context *quickjs.Context, globals quickjs.Value, parseCode quickjs.Value, syntaxKind SyntaxKind)) SyntaxKind {
	// read typescript code to be used
	typeScript, err := ioutil.ReadFile("/Users/sangupta/git/sangupta/bedrock/node_modules/typescript/lib/typescript.js")
	if err != nil {
		panic(err)
	}

	// all processing for QJS happens in same thread
	stdruntime.LockOSThread()

	// build quick js runtime
	runtime := quickjs.NewRuntime()
	defer runtime.Free()

	context := runtime.NewContext()
	defer context.Free()

	// load TS source code
	result, err := context.EvalFile(string(typeScript), 0, "typescript.js")
	check(err)
	defer result.Free()

	// never free this - throws cgo error at app termination
	globals := context.Globals()

	ts := globals.Get("ts")
	defer ts.Free()

	// read syntax kind
	sk := ts.Get("SyntaxKind")
	jsJson := globals.Get("JSON")
	stringify := jsJson.Get("stringify")
	stringifyArgs := make([]quickjs.Value, 1)
	stringifyArgs[0] = sk

	defer sk.Free()
	defer jsJson.Free()
	defer stringify.Free()

	syntaxKind := SyntaxKind{}
	syntaxKindJson, err := context.Call(globals, stringify, stringifyArgs)
	if err != nil {
		_ = json.Unmarshal([]byte(syntaxKindJson.String()), &syntaxKind)
	}

	// read script target
	scriptTarget := ts.Get("ScriptTarget")
	defer scriptTarget.Free()

	system := scriptTarget.Get("Latest")
	defer system.Free()

	// read parsing function
	parseCode := ts.Get("createSourceFile")
	defer parseCode.Free()

	// do the actual work
	doWork(context, globals, parseCode, syntaxKind)

	// remove OS thread lock
	stdruntime.UnlockOSThread()

	return syntaxKind
}

/**
 * Check and print the QuickJS error if any
 */
func check(err error) {
	if err != nil {
		var evalErr *quickjs.Error
		if errors.As(err, &evalErr) {
			fmt.Println(evalErr.Cause)
			fmt.Println(evalErr.Stack)
		}
		panic(err)
	}
}
