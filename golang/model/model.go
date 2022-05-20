/*
Redefine - UI component documentation

MIT License.
Copyright (c) 2022, Sandeep Gupta.

Use of this source code is governed by a MIT style license
that can be found in LICENSE file in the code repository.
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

		components := extractComponentsFromSourceFile(name, path, sourceFile)
		list = append(list, components...)
	}

	// get time spent
	duration := time.Since(start)
	fmt.Println("Total number of components extracted: " + strconv.Itoa(len(list)))
	fmt.Println("Total time in extracting components: " + duration.String())

	return list
}

// Get components as defined in a single source file.
// This is useful for testing
func GetComponentsFromSourceFile(sourceFile *ast.SourceFile, syntaxKind *ast.SyntaxKind, name string, path string) []Component {
	// setup syntax so anyone can use it within the package
	Syntax = syntaxKind

	// convert and return
	return extractComponentsFromSourceFile(name, path, *sourceFile)
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
func extractComponentsFromSourceFile(name string, path string, sourceFile ast.SourceFile) []Component {
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

// Extract name and path from a complete full absolute path.
// Returns the name as the first part and path as the second
// part in the return values.
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
	if !(classDeclStatement.HasExportModifier() || source.IsNameExported(classDeclStatement.Name.EscapedText)) {
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

	// read and build a map (if available) of default values
	// of the component props. We build it before reading the
	// props themselves, so that we can assign the default
	// values within the same loop
	propDefaultValueMap := getPropsDefaultValuesIfAvailable(&classDeclStatement)

	// find component props and their types
	if len(componentTypeWrapper.ClauseType.TypeArguments) > 0 {
		// the first argument specifies the props
		// this is the interface as specified as the first
		// argument in the heritage clause
		typeReference := componentTypeWrapper.ClauseType.TypeArguments[0]

		// find all members of the interface from the source file
		// we just read all members of the interface
		// TODO: we need to find and read all members of any super type
		// as well here, so that we can create a single list of all
		// properties
		members := source.GetMembersOfType(typeReference.TypeName.EscapedText)

		// document all the members as thi components props of this
		// component. We create a value object for each member we found
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

	// iterate over all properties
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
	switch property.Initializer.Kind {
	case Syntax.TrueKeyword:
		return "true"

	case Syntax.FalseKeyword:
		return "false"

	case Syntax.StringLiteral:
		return property.Initializer.Text

	case Syntax.NumericLiteral:
		return property.Initializer.Text

	case Syntax.Identifier:
		return property.Initializer.EscapedText

	case Syntax.NullKeyword:
		return "null"
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
func extractFunctionBasedComponent(path string, source ast.SourceFile, functionStatement ast.Statement) *Component {
	// skip if there is no export modifier - we only document
	// public components
	if !(functionStatement.HasExportModifier() || source.IsNameExported(functionStatement.Name.EscapedText)) {
		return nil
	}

	// check if the function body
	// has a return type of Jsx
	if functionStatement.Body == nil {
		return nil
	}

	// there must be atleast one statement in the function
	// because we need to return a JSX value
	if len(functionStatement.Body.Statements) == 0 {
		return nil
	}

	// check if this function has parameters
	// function parameters are what define the function
	// props (directly, or via destructuring)
	if len(functionStatement.Parameters) > 0 {
		// TODO
	}

	// check each statement in the body
	// if there is a return statement that returns any kind
	// of JSX expression, we will consider this as a component.
	// Similarly, if this function body contains any JSX value
	// in there, we will consider this to be a function component
	for _, statement := range functionStatement.Body.Statements {
		// this is a return statement
		// this takes care of something `return <MyComponent />`
		if Syntax.IsReturnStatement(&statement) && Syntax.IsJsxElement(statement.Expression) {
			// (Syntax.IsParenthesizedExpression(statement.Expression) && Syntax.IsJsxElement(statement.Expression.Expression))) {
			return createFunctionComponentDef(path, functionStatement)
		}

		// body is ParenthesizedExpression with JSX
		// const NewComponent = () => <MyComponent />
		if Syntax.IsParenthesizedExpression(&statement) && Syntax.IsJsxElement(statement.Expression) {
			return createFunctionComponentDef(path, functionStatement)
		}
	}

	return nil
}

func createFunctionComponentDef(path string, functionStatement ast.Statement) *Component {
	componentDef := Component{
		Name:          functionStatement.Name.EscapedText,
		SourcePath:    path,
		ComponentType: REACT_FUNCTION_COMPONENT,
		Description:   ast.GetJsDoc(functionStatement.JsDoc),
	}

	return &componentDef
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
	// create a prop definition for the member
	propDefintion := PropDef{
		Name:        member.Name.EscapedText,
		Description: ast.GetJsDoc(member.JsDoc),
	}

	// check if prop is required or not
	// this is done by checking the question mark token
	if member.QuestionToken != nil {
		propDefintion.Required = false
	} else {
		propDefintion.Required = true
	}

	// get the prop type if available
	// this is a tricky place. The prop type may not have
	// been explicitly defined.
	if member.TypeReference != nil && member.TypeReference.TypeName != nil {
		propDefintion.PropType = member.TypeReference.TypeName.EscapedText
	} else {
		memberType := Syntax.GetType(member.TypeReference)
		if !Syntax.IsUnknownType(memberType) && !Syntax.IsFunctionType(member.TypeReference) {
			propDefintion.PropType = memberType
		} else {
			// Is this a union type? for example `myProp: string | bool`
			if Syntax.IsUnionType(member.TypeReference) {
				propDefintion.PropType = "$enum"

				// check under type.types - it carries a list
				// of all types of which this value is a union of
				propDefintion.EnumTypes = make([]ParamDef, len(member.TypeReference.Types))

				// iterate and add
				for index, individualType := range member.TypeReference.Types {
					if individualType.Kind == Syntax.TypeReference {
						// we read the value from typeName.escapedText
						propDefintion.EnumTypes[index] = ParamDef{
							Name:      individualType.TypeName.EscapedText,
							ParamType: "",
						}
					} else if individualType.Kind == Syntax.LiteralType {
						// we read the value from literal.text
						propDefintion.EnumTypes[index] = ParamDef{
							Name:      individualType.Literal.Text,
							ParamType: Syntax.GetType(individualType.Literal),
						}
					}
				}
			} else if Syntax.IsFunctionType(member.TypeReference) {
				propDefintion.PropType = "$function"

				if member.TypeReference.Parameters != nil {
					propDefintion.Params = make([]ParamDef, 0)

					// build the type using definitions
					for _, param := range member.TypeReference.Parameters {
						propDefintion.Params = append(propDefintion.Params, ParamDef{
							Name:      param.Name.EscapedText,
							ParamType: Syntax.GetType(param.TypeReference),
						})
					}

					// set return type of function
					propDefintion.ReturnType = Syntax.GetType(member.TypeReference.TypeValue)
				}
			}
		}
	}

	// set default value if applicable
	propDefintion.DefaultValue = propDefaultValueMap[propDefintion.Name]

	return &propDefintion
}
