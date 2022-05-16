/*
Redefine - UI component documentation

MIT License.
Copyright (c) 2022, Sandeep Gupta.

Use of this source code is governed by a MIT style license
that can be found in LICENSE file in the code repository.
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast "sangupta.com/redefine/ast"
	"sangupta.com/redefine/model"
)

func TestEmptySourceFile(t *testing.T) {
	code := ""

	components := getComponents(code)
	assert.True(t, len(components) == 0, "Test when source file is empty")
}

func TestClassComponentWithNoPropsAndNoJsDoc(t *testing.T) {
	code := `import React from 'react';
	
	export default class HelloWorld extends React.Component {

		render() {
			return <div>Hello World</div>
		}

	}
	`

	components := getComponents(code)
	assert.True(t, len(components) == 1, "Test when source file is empty")

	component := components[0]
	assert.Equal(t, "", component.Description)
	assert.Equal(t, model.REACT_CLASS_COMPONENT, component.ComponentType)
	assert.Equal(t, 0, len(component.Props))
}

func TestClassComponentWithPropsAndWithJsDoc(t *testing.T) {
	code := `import React from 'react';

	interface HelloWorldProps {
		/**
		 * I am param string.
		 */
		paramString:string;

		/**
		 * I am param bool.
		 */
		paramBool?:bool;

		/**
		 * I am param any.
		 */
		paramAny:any;

		/**
		 * I am param number.
		 */
		paramNumber:number;

		/**
		 * I am param object.
		 */
		paramObject:object;

		/**
		 * I am param function.
		 */
		paramFunction:Function;

		/**
		 * I am param arrow function.
		 */
		paramEmptyArrowFunction:() => void;

		/**
		 * I am param arrow function with args.
		 */
		paramArrowFunction: (str:string, num:number) => object;
	}
	
	/**
	 * This is a hello world component
	 */
	export default class HelloWorld extends React.Component<HelloWorldProps> {

		static defaultProps = {
			paramString: "hello",
			paramBool:false,
			paramAny: { name: "Redefine" },
			paramNumber: 256,
			paramObject: { hello : "world" },
			paramFunction: () => {},
		}

		render() {
			return <div>Hello World</div>
		}

	}
	`

	components := getComponents(code)
	assert.True(t, len(components) == 1, "Test when source file is empty")

	component := components[0]
	assert.Equal(t, "This is a hello world component", component.Description)
	assert.Equal(t, model.REACT_CLASS_COMPONENT, component.ComponentType)
	assert.Equal(t, 8, len(component.Props))

	// first param
	param := component.Props[0]
	assert.Equal(t, "paramString", param.Name)
	assert.Equal(t, "string", param.PropType)
	assert.Equal(t, true, param.Required)
	assert.Equal(t, "hello", param.DefaultValue)
	assert.Equal(t, "I am param string.", param.Description)
	assert.Equal(t, "", param.ReturnType)
	assert.Nil(t, param.Params)
	assert.Nil(t, param.EnumTypes)

	// second param
	param = component.Props[1]
	assert.Equal(t, "paramBool", param.Name)
	assert.Equal(t, "bool", param.PropType)
	assert.Equal(t, false, param.Required)
	assert.Equal(t, "false", param.DefaultValue)
	assert.Equal(t, "I am param bool.", param.Description)
	assert.Equal(t, "", param.ReturnType)
	assert.Nil(t, param.Params)
	assert.Nil(t, param.EnumTypes)

	// third param
	param = component.Props[2]
	assert.Equal(t, "paramAny", param.Name)
	assert.Equal(t, "any", param.PropType)
	assert.Equal(t, true, param.Required)
	assert.Equal(t, "", param.DefaultValue)
	assert.Equal(t, "I am param any.", param.Description)
	assert.Equal(t, "", param.ReturnType)
	assert.Nil(t, param.Params)
	assert.Nil(t, param.EnumTypes)

	// fourth param
	param = component.Props[3]
	assert.Equal(t, "paramNumber", param.Name)
	assert.Equal(t, "number", param.PropType)
	assert.Equal(t, true, param.Required)
	assert.Equal(t, "256", param.DefaultValue)
	assert.Equal(t, "I am param number.", param.Description)
	assert.Equal(t, "", param.ReturnType)
	assert.Nil(t, param.Params)
	assert.Nil(t, param.EnumTypes)

	// fifth param
	param = component.Props[4]
	assert.Equal(t, "paramObject", param.Name)
	assert.Equal(t, "object", param.PropType)
	assert.Equal(t, true, param.Required)
	assert.Equal(t, "", param.DefaultValue)
	assert.Equal(t, "I am param object.", param.Description)
	assert.Equal(t, "", param.ReturnType)
	assert.Nil(t, param.Params)
	assert.Nil(t, param.EnumTypes)

	// sixth param
	param = component.Props[5]
	assert.Equal(t, "paramFunction", param.Name)
	assert.Equal(t, "Function", param.PropType)
	assert.Equal(t, true, param.Required)
	assert.Equal(t, "", param.DefaultValue)
	assert.Equal(t, "I am param function.", param.Description)
	assert.Equal(t, "", param.ReturnType)
	assert.Nil(t, param.Params)
	assert.Nil(t, param.EnumTypes)

	// seventh param
	param = component.Props[6]
	assert.Equal(t, "paramEmptyArrowFunction", param.Name)
	assert.Equal(t, "$function", param.PropType)
	assert.Equal(t, true, param.Required)
	assert.Equal(t, "", param.DefaultValue)
	assert.Equal(t, "I am param arrow function.", param.Description)
	assert.Equal(t, "void", param.ReturnType)
	assert.Equal(t, 0, len(param.Params))
	assert.Nil(t, param.EnumTypes)

	// eighth param
	param = component.Props[7]
	assert.Equal(t, "paramArrowFunction", param.Name)
	assert.Equal(t, "$function", param.PropType)
	assert.Equal(t, true, param.Required)
	assert.Equal(t, "", param.DefaultValue)
	assert.Equal(t, "I am param arrow function with args.", param.Description)
	assert.Equal(t, "object", param.ReturnType)
	assert.Equal(t, 2, len(param.Params))
	assert.Nil(t, param.EnumTypes)

	assert.Equal(t, "str", param.Params[0].Name)
	assert.Equal(t, "string", param.Params[0].ParamType)

	assert.Equal(t, "num", param.Params[1].Name)
	assert.Equal(t, "number", param.Params[1].ParamType)
}

func TestSimpleFunctionComponent(t *testing.T) {
	code := `
	/**
	 * Simple hello world component
	 */
	export function HelloWorld() {
		return <div>Hello World</div>
	}
	`

	components := getComponents(code)
	assert.True(t, len(components) == 1)

	component := components[0]
	assert.Equal(t, "HelloWorld", component.Name)
	assert.Equal(t, "Simple hello world component", component.Description)
	assert.Equal(t, model.REACT_FUNCTION_COMPONENT, component.ComponentType)
	assert.Equal(t, 0, len(component.Props))
}

func getComponents(code string) []model.Component {
	sourceFile, syntaxKind := ast.GetAstForFileContents(code)
	components := model.GetComponentsFromSourceFile(sourceFile, syntaxKind, "testComponent.go", "in-memory/testing")

	// jsonStr, _ := json.MarshalIndent(components, "", "  ")
	// fmt.Println(string(jsonStr))

	return components
}
