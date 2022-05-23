/*

Redefine - UI component documentation

MIT License.
Copyright (c) 2022, Sandeep Gupta.

Use of this source code is governed by a MIT style license
that can be found in LICENSE file in the code repository.

*/

package ast

// This file contains all the types
// that are used to convert parsed AST
// into a strongly-typed Go structure.

type AstNode interface {
	GetKind() int
}

type AstObject struct {
	EscapedText              string `json:"escapedText"`
	Comment                  string `json:"comment"`
	Text                     string `json:"text"`
	HasExtendedUnicodeEscape bool   `json:"hasExtendedUnicodeEscape"`
	Kind                     int    `json:"kind"`
}

type AstType struct {
	Expression AstObject `json:"expression"`
	Kind       int       `json:"kind"`
}

type Block struct {
	Statements []Statement `json:"statements"`
	Kind       int         `json:"kind"`
}

type Element struct {
	Name         AstObject `json:"name"`
	PropertyName AstObject `json:"propertyName"`
	Kind         int       `json:"kind"`
}

type Expression struct {
	Expression               *Expression  `json:"expression"`
	Name                     *AstObject   `json:"name"`
	EscapedText              string       `json:"escapedText"`
	Comment                  string       `json:"comment"`
	Text                     string       `json:"text"`
	HasExtendedUnicodeEscape bool         `json:"hasExtendedUnicodeEscape"`
	Kind                     int          `json:"kind"`
	OpeningElement           *JsxElement  `json:"openingElement"`
	Children                 []AstObject  `json:"children"`
	ClosingElement           *JsxElement  `json:"closingElement"`
	Arguments                []Expression `json:"arguments"`
}

type JsxElement struct {
	TagName    *AstObject     `json:"tagName"`
	Attributes []JsxAttribute `json:"attributes"`
	Kind       int            `json:"kind"`
}

type JsxAttribute struct {
	Properties []Member `json:"properties"`
	Kind       int      `json:"kind"`
}

type HeritageClause struct {
	Types []TypeValue `json:"types"`
	Kind  int         `json:"kind"`
}

type ImportClause struct {
	IsTypeOnly    bool           `json:"isTypeOnly"`
	Name          *AstObject     `json:"name"`
	NamedBindings *NamedBindings `json:"namedBindings"`
	Kind          int            `json:"kind"`
}

type Initializer struct {
	Properties []Property `json:"properties"`
	Kind       int        `json:"kind"`
}

type LiteralType struct {
	Literal *AstObject `json:"literal"`
	Kind    int        `json:"kind"`
}

type Member struct {
	Name          *AstObject     `json:"name"`
	TypeReference *TypeReference `json:"type"`
	QuestionToken *AstObject     `json:"questionToken"`
	JsDoc         []AstObject    `json:"jsDoc"`
	Modifiers     []AstObject    `json:"modifiers"`
	Initializer   *Initializer   `json:"initializer"`
	Kind          int            `json:"kind"`
}

type ModuleSpecifier struct {
	Text string `json:"text"`
	Kind int    `json:"kind"`
}

type NamedBindings struct {
	Name     *AstObject `json:"name"`
	Elements []Element  `json:"elements"`
	Kind     int        `json:"kind"`
}

type Parameter struct {
	Name          *AstObject     `json:"name"`
	TypeReference *TypeReference `json:"type"`
	Kind          int            `json:"kind"`
}

type Property struct {
	Name        *AstObject `json:"name"`
	Initializer *AstObject `json:"initializer"`
	Kind        int        `json:"kind"`
}

type SourceFile struct {
	Statements []Statement `json:"statements"`
	Kind       int         `json:"kind"`

	importsResolved bool
	imports         map[string]string
}

type Statement struct {
	ImportClause    *ImportClause    `json:"importClause"`
	ModuleSpecifier *ModuleSpecifier `json:"moduleSpecifier"`
	Name            *AstObject       `json:"name"`
	Body            *Block           `json:"body"`
	Expression      *Expression      `json:"expression"`
	HeritageClauses []HeritageClause `json:"heritageClauses"`
	Modifiers       []AstObject      `json:"modifiers"`
	Members         []Member         `json:"members"`
	JsDoc           []AstObject      `json:"jsDoc"`
	Parameters      []Parameter      `json:"parameters"`
	Kind            int              `json:"kind"`
}

type TypeValue struct {
	Expression    *Expression     `json:"expression"`
	TypeArguments []TypeReference `json:"typeArguments"`
	Kind          int             `json:"kind"`
}

type TypeReference struct {
	TypeName   *AstObject      `json:"typeName"`
	TypeValue  *AstObject      `json:"type"`
	Types      []TypeReference `json:"types"`
	Parameters []Parameter     `json:"parameters"`
	Literal    *AstObject      `json:"literal"`
	Kind       int             `json:"kind"`
}

// implement AstNode interface

func (ast *AstObject) GetKind() int {
	return ast.Kind
}

func (ast *AstType) GetKind() int {
	return ast.Kind
}

func (ast *Block) GetKind() int {
	return ast.Kind
}

func (ast *Element) GetKind() int {
	return ast.Kind
}

func (ast *Expression) GetKind() int {
	return ast.Kind
}

func (ast *ImportClause) GetKind() int {
	return ast.Kind
}

func (ast *HeritageClause) GetKind() int {
	return ast.Kind
}

func (ast *Initializer) GetKind() int {
	return ast.Kind
}

func (ast *LiteralType) GetKind() int {
	return ast.Kind
}

func (ast *Member) GetKind() int {
	return ast.Kind
}

func (ast *ModuleSpecifier) GetKind() int {
	return ast.Kind
}

func (ast *NamedBindings) GetKind() int {
	return ast.Kind
}

func (ast *Parameter) GetKind() int {
	return ast.Kind
}

func (ast *Property) GetKind() int {
	return ast.Kind
}

func (ast *SourceFile) GetKind() int {
	return ast.Kind
}

func (ast *Statement) GetKind() int {
	return ast.Kind
}

func (ast *TypeValue) GetKind() int {
	return ast.Kind
}

func (ast *TypeReference) GetKind() int {
	return ast.Kind
}

// convenience methods

func (st *Statement) GetClassName() string {
	if !Syntax.IsClassDeclaration(st) {
		panic("Expected a class declaration")
	}

	return st.Name.EscapedText
}

func (st *Statement) HasExportModifier() bool {
	for _, modifier := range st.Modifiers {
		if modifier.Kind == Syntax.ExportKeyword {
			return true
		}
	}

	return false
}

func (st *Statement) HasDefaultModifier() bool {
	for _, modifier := range st.Modifiers {
		if modifier.Kind == Syntax.DefaultKeyword {
			return true
		}
	}

	return false
}

func (st *Statement) HasHeritageClauses() bool {
	return len(st.HeritageClauses) > 0
}

func (sf *SourceFile) GetImportPath(key string) string {
	if !sf.importsResolved {
		sf.resolveImports()
	}

	return sf.imports[key]
}

// This method checks if a given name is exported in the file
// for example, `export default <name>`.
func (sf *SourceFile) IsNameExported(name string) bool {
	if len(sf.Statements) == 0 {
		return false
	}

	for _, st := range sf.Statements {
		// check we this is of form `export MyComponent`
		if Syntax.IsExpressionStatement(&st) {
			if st.Expression != nil && Syntax.IsIdentifier(st.Expression) && st.Expression.EscapedText == name {
				return true
			}

			continue
		}

		// check if we have an export assignment
		// of the form is `export injectIntl(MyComponent)`
		if Syntax.IsExportAssignment(&st) {
			if st.Expression != nil {
				// for `export default injectIntl(MyComponent)`
				if Syntax.IsCallExpression(st.Expression) && len(st.Expression.Arguments) > 0 {
					for _, arg := range st.Expression.Arguments {
						if Syntax.IsIdentifier(&arg) && arg.EscapedText == name {
							return true
						}
					}
				}

				// for simple `export MyComponent`
				if Syntax.IsIdentifier(st.Expression) && st.Expression.EscapedText == name {
					return true
				}
			}

			continue
		}
	}

	return false
}

func (sf *SourceFile) resolveImports() {
	if len(sf.Statements) == 0 {
		return
	}

	if sf.imports == nil {
		sf.imports = make(map[string]string, 0)
	}

	for _, st := range sf.Statements {
		if !Syntax.IsImportDeclaration(&st) {
			continue
		}

		// this is an imports clause
		library := st.ModuleSpecifier.Text
		if st.ImportClause.Name != nil {
			sf.imports[st.ImportClause.Name.EscapedText] = library
		}

		if st.ImportClause.NamedBindings != nil {
			if st.ImportClause.NamedBindings.Name != nil {
				sf.imports[st.ImportClause.NamedBindings.Name.EscapedText] = library
			}

			for _, element := range st.ImportClause.NamedBindings.Elements {
				sf.imports[element.Name.EscapedText] = library
			}
		}
	}

	sf.importsResolved = true
}

// Check if the source file has any class declaration. The first
// class declaration found returns `true`.
func (sf *SourceFile) HasClassDeclaration() bool {
	for _, statement := range sf.Statements {
		if Syntax.IsClassDeclaration(&statement) {
			return true
		}
	}

	return false
}

// Given a type name find all members of the type in the source file.
// The type represented here is mostly an interface and specifies
// the props of the React component
func (sf *SourceFile) GetMembersOfType(typeName string) []Member {
	// is this an imported typeName
	importLibrary := sf.GetImportPath(typeName)

	if len(importLibrary) > 0 {
		return sf.GetMembersOfTypeFromLibrary(importLibrary, typeName)
	}

	for _, statement := range sf.Statements {
		if Syntax.IsInterfaceDeclaration(&statement) {
			if statement.Name != nil && typeName == statement.Name.EscapedText {
				return statement.Members
			}
		}
	}

	return nil
}

// Find the members (aka props) of given type from a different library
// or import path. This usually happens when we want to pull props or extend
// props from an interface defined else where in the code.
func (sf *SourceFile) GetMembersOfTypeFromLibrary(importLibrary string, typeName string) []Member {
	return nil
}

// Check if the member has a `static` modifier applied to it or not.
// This is usually checked when reading default props.
func (member *Member) HasStaticModifier() bool {
	for _, modifier := range member.Modifiers {
		if modifier.Kind == Syntax.StaticKeyword {
			return true
		}
	}

	return false
}
