/**
 * Redefine
 *
 * MIT License.
 * Copyright (c) 2022, Sandeep Gupta.
 *
 * Use of this source code is governed by a MIT style license
 * that can be found in LICENSE file in the code repository.
 */

package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"sangupta.com/redefine/ast"
)

/**
 * Get a map of all components against the file that they
 * are present in.
 */
func GetComponents(astMap map[string]ast.SourceFile) []Component {
	start := time.Now()

	// create component list
	list := make([]Component, 0)

	if len(astMap) == 0 {
		return list
	}

	for file, sourceFile := range astMap {
		name, path := getNameAndPath(file)

		components := extractComponents(name, path, sourceFile)
		list = append(list, components...)
	}

	// get time spent
	duration := time.Since(start)
	fmt.Println("Total number of components extracted: " + strconv.Itoa(len(list)))
	fmt.Println("Total time in extracting components: " + duration.String())

	return list
}

/**
 * Extract a list of components from a given source file
 */
func extractComponents(name string, path string, sourceFile ast.SourceFile) []Component {
	cl := make([]Component, 0)

	if len(sourceFile.Statements) == 0 {
		return cl
	}

	for _, statement := range sourceFile.Statements {
		// detect class based components
		if ast.IsClassDeclaration(&statement) {
			component := extractClassBasedComponents(path, sourceFile, statement)
			if component != nil {
				cl = append(cl, *component)
			}
			continue
		}

		// detect function based components
		if ast.IsFunctionDeclaration(&statement) {
			component := extractFunctionBasedComponent(path, sourceFile, statement)
			if component != nil {
				cl = append(cl, *component)
			}
			continue
		}
	}

	return cl
}

/**
 * Extract name and path from a complete full absolute path
 */
func getNameAndPath(str string) (string, string) {
	if string(str[len(str)-1]) == "/" {
		str = str[0 : len(str)-1]
	}

	lastSlash := strings.LastIndex(str, "/")
	if lastSlash < 0 {
		return str, ""
	}

	return str[lastSlash+1:], str[0:lastSlash]
}

/**
 * Extract a class based component (if applicable) from the given statement
 */
func extractClassBasedComponents(path string, source ast.SourceFile, st ast.Statement) *Component {
	// skip if there is no export modifier - we only document
	// public components
	if !st.HasExportModifier() {
		return nil
	}

	// check if the class has any heritage clause - means it extends
	// another clause. A class component must extend React.Component
	// to be a component
	if !st.HasHeritageClauses() {
		return nil
	}

	// case 1: has export keyword, and extend react.component or just component from both react library
	// the class must have a method called "render" to be a component
	if !ast.HasMethodOfName(st, "render") {
		return nil
	}

	// all checks pass - this is a class based component
	// verify if it extends from React or not
	componentTypeWrapper := detectComponentType(source, st)
	if componentTypeWrapper == nil || componentTypeWrapper.ComponentType == nil {
		return nil
	}

	// class extends and is definitely a react component
	def := Component{
		Name:          st.GetClassName(),
		SourcePath:    path,
		ComponentType: *componentTypeWrapper.ComponentType,
		Description:   ast.GetJsDoc(st.JsDoc),
		Props:         make([]PropDef, 0),
	}

	// find component props
	if len(componentTypeWrapper.Type.TypeArguments) > 0 {
		// the first argument specifies the props
		typeReference := componentTypeWrapper.Type.TypeArguments[0]

		// find all members of the interface from the source file
		members := source.GetMembersOfType(typeReference.TypeName.EscapedText)

		// document all the members as this components props
		if len(members) > 0 {
			for _, member := range members {
				def.Props = append(def.Props, *getComponentProp(member))
			}
		}
	}

	// // for all these prop members, see if there is a default value specified or not
	// Object x = ComponentUtils.getDefaultPropsStaticClassMember(statement);

	return &def
}

/**
 * Extract a function based component (if applicable) from the given statement
 */
func extractFunctionBasedComponent(path string, source ast.SourceFile, st ast.Statement) *Component {
	return nil
}

func getComponentProp(member ast.Member) *PropDef {
	def := PropDef{
		Name:        member.Name.EscapedText,
		Description: ast.GetJsDoc(member.JsDoc),
	}

	// if member.Type.TypeName != nil {

	// }

	return &def
}
