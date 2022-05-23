/*

Redefine - UI component documentation

MIT License.
Copyright (c) 2022, Sandeep Gupta.

Use of this source code is governed by a MIT style license
that can be found in LICENSE file in the code repository.

*/

package model

import (
	"strings"

	"sangupta.com/redefine/ast"
)

type ComponentTypeWrapper struct {
	ComponentType ComponentType
	Clause        *ast.HeritageClause
	ClauseType    *ast.TypeValue
	Detected      bool
}

// This method detects the component type, its heritage
// clause (read the interface implementing the props)
// and the applicable `HeritageClause.Type`
func detectComponentType(sourceFile ast.SourceFile, classDeclStatement ast.Statement) *ComponentTypeWrapper {
	if len(classDeclStatement.HeritageClauses) == 0 {
		return nil
	}

	for _, clause := range classDeclStatement.HeritageClauses {
		for _, clauseType := range clause.Types {
			expr := clauseType.Expression

			var exprText string

			if expr != nil {
				if Syntax.IsPropertyAccessExpression(expr) {
					exprText = expr.Expression.EscapedText
					name := expr.Name.EscapedText

					if !(name == "Component" || name == "PureComponent") {
						continue
					}
				}

				if Syntax.IsIdentifier(expr) {
					exprText = expr.EscapedText
				}

				if exprText == "" {
					continue
				}

				if isReactImport(sourceFile, exprText) {
					return &ComponentTypeWrapper{
						ComponentType: REACT_CLASS_COMPONENT,
						Clause:        &clause,
						ClauseType:    &clauseType,
						Detected:      true,
					}
				}
			}
		}
	}

	return nil
}

// Check if the name of the import associated with
// the class `extends` keyword is coming in from `react`.
// To find it out, we run a check against all import
// statements in the source file, and see if the import
// object is coming in from a package called `react`.
func isReactImport(sourceFile ast.SourceFile, name string) bool {
	for _, st := range sourceFile.Statements {
		if !Syntax.IsImportDeclaration(&st) {
			continue
		}

		// check in imports
		if strings.EqualFold(sourceFile.GetImportPath(name), "react") {
			return true
		}
	}

	return false
}
