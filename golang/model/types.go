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

type Component struct {
	Name          string
	SourcePath    string
	ComponentType ComponentType
	Description   string
	Props         []PropDef
}

type PropDef struct {
	Name         string
	PropType     string
	Required     bool
	DefaultValue string
	Description  string
	ReturnType   string
	Params       []ParamDef
}

type ParamDef struct {
	Name      string
	ParamType string
}

type ComponentType int64

const (
	REACT_CLASS_COMPONENT ComponentType = iota
	REACT_FUNCTION_COMPONENT
)
