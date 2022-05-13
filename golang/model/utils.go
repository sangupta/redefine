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
	"strings"

	"sangupta.com/redefine/ast"
)

type ComponentTypeWrapper struct {
	ComponentType *ComponentType
	Clause        *ast.HeritageClause
	Type          *ast.TypeValue
}

func detectComponentType(ast ast.SourceFile, st ast.Statement) *ComponentTypeWrapper {
	if len(st.HeritageClauses) == 0 {
		return nil
	}

	for _, clause := range st.HeritageClauses {
		for _, typ := range clause.Types {
			expr := typ.Expression.Expression.EscapedText
			name := typ.Expression.Name.EscapedText

			if (name == "Component" || name == "PureComponent") && isReactImport(ast, expr) {
				return &ComponentTypeWrapper{}
			}
		}
	}

	return nil
}

func isReactImport(sourceFile ast.SourceFile, name string) bool {
	for _, st := range sourceFile.Statements {
		if !ast.IsImportDeclaration(&st) {
			continue
		}

		// check in imports
		if strings.EqualFold(sourceFile.GetImportPath(name), "react") {
			return true
		}
	}

	return false
}
