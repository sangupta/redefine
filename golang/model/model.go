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
 *
 * @param fileAstMap a `map` of file paths v/s their `SourceFile`
 *		instance.
 *
 * @param syntaxKind the `SyntaxKind` object as extracted from
 *		the typescript compiler.
 */
func GetComponents(fileAstMap map[string]ast.SourceFile, syntaxKind *ast.SyntaxKind) []Component {
	// setup syntax so anyone can use it within the package
	Syntax = syntaxKind

	// start timing
	start := time.Now()

	// create component list
	list := make([]Component, 0)

	// if there are no files, just skip
	if len(fileAstMap) == 0 {
		return list
	}

	// process for each file path and AST
	for file, sourceFile := range fileAstMap {
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
 * Extract a list of components from a given source file.
 *
 * @param name the name of the file defining component.
 *
 * @param path the absolute path to the file
 *
 * @param sourceFile the `SourceFile` instance describing the AST
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

	propDefaultValueMap := getPropsDefaultValuesIfAvailable(&classDeclStatement)

	// find component props
	if len(componentTypeWrapper.Type.TypeArguments) > 0 {
		// the first argument specifies the props
		typeReference := componentTypeWrapper.Type.TypeArguments[0]

		// find all members of the interface from the source file
		members := source.GetMembersOfType(typeReference.TypeName.EscapedText)

		// document all the members as this components props
		if len(members) > 0 {
			for _, member := range members {
				componentDef.Props = append(componentDef.Props, *getComponentProp(member, propDefaultValueMap))
			}
		}
	}

	return &componentDef
}

/**
 * Build a map of default values for all props for this
 * component. The key is the name of the prop, and value
 * the default value of prop. If there is no default value
 * for a prop, the key for that prop is not present in the
 * map.
 */
func getPropsDefaultValuesIfAvailable(classDeclStatement *ast.Statement) map[string]string {
	defaultValueMap := make(map[string]string, 0)

	// for all these prop members, see if there is a default value specified or not
	defaultProps := findDefaultPropsMember(classDeclStatement)
	if defaultProps == nil {
		return defaultValueMap
	}

	// check if initializer and properties exist
	if defaultProps.Initializer == nil && len(defaultProps.Initializer.Properties) == 0 {
		return defaultValueMap
	}

	for _, property := range defaultProps.Initializer.Properties {
		propName := property.Name.EscapedText
		propValue := extractPropValue(property)

		defaultValueMap[propName] = propValue
	}

	return defaultValueMap
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

/**
 * Given a class definition, find the member that
 * is named `defaultProps` and is `static` defined.
 */
func findDefaultPropsMember(classDeclStatement *ast.Statement) *ast.Member {
	if len(classDeclStatement.Members) == 0 {
		return nil
	}

	for _, member := range classDeclStatement.Members {
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

/**
 * Generate component prop definition from a class
 * component member, defined in its interface.
 *
 * @param member `Member` instance of the `interface`
 * 		implementing the props.
 *
 * @param propDefaultValueMap a `map` of default values
 * 		for props as read from the `static defaultProps`
 * 		read from the component.
 */
func getComponentProp(member ast.Member, propDefaultValueMap map[string]string) *PropDef {
	def := PropDef{
		Name:        member.Name.EscapedText,
		Description: ast.GetJsDoc(member.JsDoc),
	}

	// check if prop is required or not
	if member.QuestionToken != nil {
		def.Required = false
	} else {
		def.Required = true
	}

	// get the prop type if available
	if member.TypeReference != nil && member.TypeReference.TypeName != nil {
		def.PropType = member.TypeReference.TypeName.EscapedText
	} else {
		memberType := Syntax.GetType(member.TypeReference)
		if !Syntax.IsUnknownType(memberType) && !Syntax.IsFunctionType(member.TypeReference) {
			def.PropType = memberType
		} else {
			if Syntax.IsUnionType(member.TypeReference) {
				def.PropType = "$enum"
			} else if Syntax.IsFunctionType(member.TypeReference) {
				def.PropType = "$function"

				if member.TypeReference.Parameters != nil {
					def.Params = make([]ParamDef, 0)

					// build the type using definitions
					for _, param := range member.TypeReference.Parameters {
						def.Params = append(def.Params, ParamDef{
							Name:      param.Name.EscapedText,
							ParamType: Syntax.GetType(param.TypeReference),
						})
					}

					// set return type of function
					def.ReturnType = Syntax.GetType(member.TypeReference.TypeValue)
				}
			}
		}
	}

	// set default value if applicable
	def.DefaultValue = propDefaultValueMap[def.Name]

	return &def
}
