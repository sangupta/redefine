/*

Redefine - UI component documentation

MIT License.
Copyright (c) 2022, Sandeep Gupta.

Use of this source code is governed by a MIT style license
that can be found in LICENSE file in the code repository.

*/

package model

type Component struct {
	Name          string        `json:"name"`
	SourcePath    string        `json:"sourcePath"`
	ComponentType ComponentType `json:"componentType"`
	Description   string        `json:"description"`
	Props         []PropDef     `json:"props"`
}

type PropDef struct {
	Name         string     `json:"name"`
	PropType     string     `json:"type"`
	EnumTypes    []ParamDef `json:"enumOf"`
	Required     bool       `json:"required"`
	DefaultValue string     `json:"defaultValue"`
	Description  string     `json:"description"`
	ReturnType   string     `json:"returnType"`
	Params       []ParamDef `json:"params"`
}

type ParamDef struct {
	Name      string `json:"name"`
	ParamType string `json:"type"`
}

type ComponentType int64

const (
	REACT_CLASS_COMPONENT ComponentType = iota
	REACT_FUNCTION_COMPONENT
)
