/**
 * Redefine
 *
 * MIT License.
 * Copyright (c) 2022, Sandeep Gupta.
 *
 * Use of this source code is governed by a MIT style license
 * that can be found in LICENSE file in the code repository.
 */

package ast

import "strings"

const (
	KIND_VOID_KEYWORD         = 113
	KIND_ANY_KEYWORD          = 128
	KIND_BOOLEAN_KEYWORD      = 131
	KIND_NEVER_KEYWORD        = 141
	KIND_NUMBER_KEYWORD       = 144
	KIND_STRING_KEYWORD       = 147
	KIND_UNDEFINED_KEYWORD    = 150
	KIND_ARROW_METHOD         = 164
	KIND_SIMPLE_METHOD        = 166
	KIND_FUNCTION_TYPE        = 175
	KIND_UNION_TYPE           = 183
	KIND_NULL_KEYWORD         = 192
	KIND_PROPERTY_ACCESS_EXPR = 198
	KIND_PARANTHESIS_EXPR     = 208
	KIND_EXPR_WITH_TYPE_ARGS  = 224
	KIND_RETURN_STATEMENT     = 243
	KIND_CLASS_DECL           = 253
	KIND_FUNCTION_DECL        = 252
	KIND_INTERFACE_DECL       = 254
	KIND_IMPORT_DECL          = 262
	KIND_JSX_ELEMENT          = 274
	KIND_JSX_FRAGMENT         = 278
	KIND_HERITAGE_CLAUSE      = 287
	UNKNOWN                   = "$unknown"
)

func getNodeType(node AstNode) string {
	switch node.GetKind() {
	case 253:
		return "ClassDeclaration"
	case 252:
		return "FunctionDeclaration"
	case 254:
		return "InterfaceDeclaration"
	case 262:
		return "ImportDeclaration"
	}

	return "Unknown"
}

func IsReturnStatement(node AstNode) bool {
	return node.GetKind() == KIND_RETURN_STATEMENT
}

func IsClassDeclaration(node AstNode) bool {
	return node.GetKind() == KIND_CLASS_DECL
}

func IsArrowMethodDeclaration(node AstNode) bool {
	return node.GetKind() == KIND_ARROW_METHOD
}

func IsSimpleMethodDeclaration(node AstNode) bool {
	return node.GetKind() == KIND_SIMPLE_METHOD
}

func IsInterfaceDeclaration(node AstNode) bool {
	return node.GetKind() == KIND_INTERFACE_DECL
}

func IsImportDeclaration(node AstNode) bool {
	return node.GetKind() == KIND_IMPORT_DECL
}

func IsPropertyAccessExpression(node AstNode) bool {
	return node.GetKind() == KIND_PROPERTY_ACCESS_EXPR
}

func IsExpressionWithTypeArguments(node AstNode) bool {
	return node.GetKind() == KIND_EXPR_WITH_TYPE_ARGS
}

func IsHeritageClause(node AstNode) bool {
	return node.GetKind() == KIND_HERITAGE_CLAUSE
}

func IsUnionType(node AstNode) bool {
	return node.GetKind() == KIND_UNION_TYPE
}

func IsFunctionDeclaration(node AstNode) bool {
	return node.GetKind() == KIND_FUNCTION_DECL
}

func IsMethodDeclaration(node AstNode) bool {
	return IsSimpleMethodDeclaration(node) || IsArrowMethodDeclaration(node)
}

func IsFunctionType(node AstNode) bool {
	return node.GetKind() == KIND_FUNCTION_TYPE
}

func HasMethodOfName(statement Statement, methodName string) bool {
	if IsClassDeclaration(&statement) {
		return false
	}

	for _, member := range statement.Members {
		if IsMethodDeclaration(&member) && member.Name.EscapedText == methodName {
			return true
		}
	}

	return false
}

func GetJsDoc(jsDoc []AstObject) string {
	if len(jsDoc) == 0 {
		return ""
	}

	if len(jsDoc) == 1 {
		return jsDoc[0].Comment
	}

	var sb strings.Builder
	for _, doc := range jsDoc {
		sb.WriteString(doc.Comment)
		sb.WriteRune('\n')
	}

	return sb.String()
}

func DoesMethodReturnsJsxFragement(body Block) bool {
	if len(body.Statements) == 0 {
		return false
	}

	for _, st := range body.Statements {
		if IsReturnStatement(&st) && IsJsxElement(st.Expression) {
			return true
		}
	}

	return false
}

func IsJsxElement(expr *Expression) bool {
	if expr == nil {
		return false
	}

	// direct JsxElement - return <component> ... </component>
	if expr.Kind == KIND_JSX_ELEMENT {
		return true
	}

	// jsx-frgament - multiple nodes - return <> ... </>
	if expr.Kind == KIND_JSX_FRAGMENT {
		return true
	}

	// inside a paranthesis - return ( <> ... </>)
	if expr.Kind == KIND_PARANTHESIS_EXPR && IsJsxElement(expr.Expression) {
		return true
	}

	return false
}

func GetType(node AstNode) string {
	switch node.GetKind() {
	case KIND_NUMBER_KEYWORD:
		return "number"
	case KIND_STRING_KEYWORD:
		return "string"
	case KIND_BOOLEAN_KEYWORD:
		return "boolean"
	case KIND_VOID_KEYWORD:
		return "void"
	case KIND_FUNCTION_TYPE:
		return "Function"
	case KIND_ANY_KEYWORD:
		return "any"
	case KIND_NULL_KEYWORD:
		return "null"
	case KIND_UNDEFINED_KEYWORD:
		return "undefined"
	case KIND_NEVER_KEYWORD:
		return "never"
	}

	return UNKNOWN
}
