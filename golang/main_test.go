package main

import (
	"encoding/json"
	"fmt"
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
		paramAny:any;
		paramNumber:number;
		paramObject:object;
		paramFunction:Function;

		paramEmptyArrowFunction:() => void;
		paramArrowFunction: (str:string, num:number) => object;
	}
	
	/**
	 * This is a hello world component
	 */
	export default class HelloWorld extends React.Component<HelloWorldProps> {

		static defaultProps = {
			paramString: "hello"
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
	assert.Equal(t, "", param.DefaultValue)
	assert.Equal(t, "I am param bool.", param.Description)
	assert.Equal(t, "", param.ReturnType)
	assert.Nil(t, param.Params)
	assert.Nil(t, param.EnumTypes)
}

func getComponents(code string) []model.Component {
	sourceFile, syntaxKind := ast.GetAstForFileContents(code)
	components := model.GetComponentsFromSourceFile(sourceFile, syntaxKind, "testComponent.go", "in-memory/testing")

	jsonStr, _ := json.MarshalIndent(components, "", "  ")
	fmt.Println(string(jsonStr))

	return components
}
