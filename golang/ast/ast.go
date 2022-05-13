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

type tsParser struct {
	runtime    *quickjs.Runtime
	context    *quickjs.Context
	globals    *quickjs.Value
	codeParser *quickjs.Value
	syntaxKind *SyntaxKind
}

/**
 * Create a map of AST's by parsing each file
 */
func BuildAstForFiles(files []string) map[string]SourceFile {
	start := time.Now()
	astMap := make(map[string]SourceFile, len(files))

	// all processing for QJS happens in same thread
	stdruntime.LockOSThread()

	parser := tsParser{}
	parser.init()
	defer parser.free()

	// do job
	doWork(files, &parser, astMap)

	// get time spent
	duration := time.Since(start)
	fmt.Println("Total time in parsing files: " + duration.String())

	// remove OS thread lock
	stdruntime.UnlockOSThread()

	return astMap
}

func doWork(files []string, parser *tsParser, astMap map[string]SourceFile) {
	for _, file := range files {
		sourceFile := parseSingleFile(file, parser)
		if sourceFile != nil {
			astMap[file] = *sourceFile
		}
	}
}

/**
 * Parse a single file by reading it from disk
 */
func parseSingleFile(file string, parser *tsParser) *SourceFile {
	// read the source code file from disk
	sourceCode, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	// create argument list to call the method
	args := make([]quickjs.Value, 4)
	args[0] = parser.context.String("index.ts")
	args[1] = parser.context.String(string(sourceCode))
	args[2] = parser.context.String("")
	args[3] = parser.context.Bool(true)

	// invoke the "createSourceFile" method
	result, err := parser.context.Call(*parser.globals, *parser.codeParser, args)
	// defer result.Free()

	// check for error, and free up the result
	check(err)

	// now convert the "result" represented as AST in QJS objects
	// to the pure objects that we require
	fmt.Println("Fetching object for " + file)

	typescript := Typescript{
		syntaxKind: *parser.syntaxKind,
	}

	sourceFile := typescript.getSourceFile(result)
	return sourceFile
}

func (parser *tsParser) free() {
	parser.context.Free()
	parser.codeParser.Free()

	// finally free the runtime
	defer parser.runtime.Free()
}

func (parser *tsParser) init() {
	// read typescript code to be used
	typeScript, err := ioutil.ReadFile("/Users/sangupta/git/sangupta/bedrock/node_modules/typescript/lib/typescript.js")
	if err != nil {
		panic(err)
	}

	// build quick js runtime
	runtime := quickjs.NewRuntime()
	parser.runtime = &runtime

	context := runtime.NewContext()
	parser.context = context

	// load TS source code
	result, err := context.EvalFile(string(typeScript), 0, "typescript.js")
	check(err)
	defer result.Free()

	// never free this - throws cgo error at app termination
	globals := context.Globals()
	parser.globals = &globals

	ts := globals.Get("ts")
	defer ts.Free()

	// read syntax kind
	sk := ts.Get("SyntaxKind")
	defer sk.Free()

	// get JSON.stringify function
	jsJson := globals.Get("JSON")
	defer jsJson.Free()

	stringify := jsJson.Get("stringify")
	defer stringify.Free()

	stringifyArgs := make([]quickjs.Value, 1)
	stringifyArgs[0] = sk

	syntaxKind := SyntaxKind{}
	syntaxKindJson, err := context.Call(globals, stringify, stringifyArgs)
	if err != nil {
		_ = json.Unmarshal([]byte(syntaxKindJson.String()), &syntaxKind)
	}

	parser.syntaxKind = &syntaxKind

	// read script target
	scriptTarget := ts.Get("ScriptTarget")
	defer scriptTarget.Free()

	system := scriptTarget.Get("Latest")
	defer system.Free()

	// read parsing function
	parseCode := ts.Get("createSourceFile")
	parser.codeParser = &parseCode
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
