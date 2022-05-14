/**
 * Redefine - UI component documentation
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

var Syntax *ast.SyntaxKind

/**
 * Get a map of all components against the file that they
 * are present in. This uses the `SourceFile` instance and
 * scans for all classes/functions to figure out candidate
 * components.
 */
func GetComponents(astMap map[string]ast.SourceFile, syntaxKind *ast.SyntaxKind) []Component {
	// setup syntax so anyone can use it
	Syntax = syntaxKind

	// start timing
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
	fmt.Println("Extracting components from: " + path + "/" + name)

	cl := make([]Component, 0)

	if len(sourceFile.Statements) == 0 {
		return cl
	}

	for _, statement := range sourceFile.Statements {
		// detect class based components
		if Syntax.IsClassDeclaration(&statement) {
			component := extractClassBasedComponents(path, sourceFile, statement)
			if component != nil {
				cl = append(cl, *component)
			}
			continue
		}

		// detect function based components
		if Syntax.IsFunctionDeclaration(&statement) {
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
func extractClassBasedComponents(path string, source ast.SourceFile, classDeclStatement ast.Statement) *Component {
	// skip if there is no export modifier - we only document
	// public components
	if !classDeclStatement.HasExportModifier() {
		return nil
	}

	// check if the class has any heritage clause - means it extends
	// another clause. A class component must extend React.Component
	// to be a component
	if !classDeclStatement.HasHeritageClauses() {
		return nil
	}

	// case 1: has export keyword, and extend react.component or just component from both react library
	// the class must have a method called "render" to be a component
	if !Syntax.HasMethodOfName(classDeclStatement, "render") {
		return nil
	}

	// all checks pass - this is a class based component
	// verify if it extends from React or not
	componentTypeWrapper := detectComponentType(source, classDeclStatement)
	if componentTypeWrapper == nil || !componentTypeWrapper.Detected {
		return nil
	}

	// class extends and is definitely a react component
	componentDef := Component{
		Name:          classDeclStatement.GetClassName(),
		SourcePath:    path,
		ComponentType: componentTypeWrapper.ComponentType,
		Description:   ast.GetJsDoc(classDeclStatement.JsDoc),
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
				componentDef.Props = append(componentDef.Props, *getComponentProp(member))
			}
		}
	}

	// if there were props detected find their default values
	if len(componentDef.Props) > 0 {
		// for all these prop members, see if there is a default value specified or not
		defaultProps := findDefaultPropsMember(path, source, classDeclStatement)
		if defaultProps != nil {

			// check if initializer and properties exist
			if defaultProps.Initializer != nil && len(defaultProps.Initializer.Properties) > 0 {
				for _, property := range defaultProps.Initializer.Properties {
					propName := property.Name.EscapedText
					propValue := extractPropValue(property)

					fmt.Println("  found default value for: " + propName + " as: " + propValue)

					// set this value as the default value for
					// the correct prop
					for _, comProp := range componentDef.Props {
						if comProp.Name == propName {
							fmt.Println("       set value for: " + comProp.Name + " as: " + propValue)
							comProp.DefaultValue = propValue
							// break so that outer loop can run
							// break
						}
					}
				}
			}
		}
	}

	return &componentDef
}

/**
 * This function extracts the property value using
 * the initializer of the property. This is only used
 * when reading properties from `static defaultProps`
 * member of the class based component.
 */
func extractPropValue(property ast.Property) string {
	if property.Initializer.Kind == Syntax.TrueKeyword {
		return "true"
	}

	if property.Initializer.Kind == Syntax.FalseKeyword {
		return "false"
	}

	return property.Initializer.EscapedText
}

func findDefaultPropsMember(path string, source ast.SourceFile, st ast.Statement) *ast.Member {
	if len(st.Members) == 0 {
		return nil
	}

	for _, member := range st.Members {
		if member.Name != nil && member.Name.EscapedText == "defaultProps" && member.HasStaticModifier() {
			return &member
		}
	}

	return nil
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

	return &def
}
