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

// MAIN TYPES

type AstNode interface {
	GetKind() int
}

type AstObject struct {
	EscapedText              string
	Comment                  string
	Text                     string
	HasExtendedUnicodeEscape bool
	Kind                     int
}

type AstType struct {
	Expression AstObject
	Kind       int
}

type Block struct {
	Statements []Statement
	Kind       int
}

type Element struct {
	Name         AstObject
	PropertyName AstObject
	Kind         int
}

type Expression struct {
	Expression               *Expression
	Name                     *AstObject
	EscapedText              string
	Comment                  string
	Text                     string
	HasExtendedUnicodeEscape bool
	Kind                     int
}

type HeritageClause struct {
	Types []TypeValue
	Kind  int
}

type ImportClause struct {
	IsTypeOnly    bool
	Name          *AstObject
	NamedBindings *NamedBindings
	Kind          int
}

type Initializer struct {
	Properties []Property
	Kind       int
}

type LiteralType struct {
	Literal *AstObject
	Kind    int
}

type Member struct {
	Name          *AstObject
	TypeReference *TypeReference
	QuestionToken *AstObject
	JsDoc         []AstObject
	Modifiers     []AstObject
	Initializer   *Initializer
	Kind          int
}

type ModuleSpecifier struct {
	Text string
	Kind int
}

type NamedBindings struct {
	Name     *AstObject
	Elements []Element
	Kind     int
}

type Parameter struct {
	Name          *AstObject
	TypeReference *TypeReference
	Kind          int
}

type Property struct {
	Name        *AstObject
	Initializer *AstObject
	Kind        int
}

type SourceFile struct {
	Statements []Statement
	Kind       int

	importsResolved bool
	imports         map[string]string
}

type Statement struct {
	ImportClause    *ImportClause
	ModuleSpecifier *ModuleSpecifier
	Name            *AstObject
	Body            *Block
	Expression      *Expression
	HeritageClauses []HeritageClause
	Modifiers       []AstObject
	Members         []Member
	JsDoc           []AstObject
	Parameters      []Parameter
	Kind            int
}

type TypeValue struct {
	Expression    *Expression
	TypeArguments []TypeReference
	Kind          int
}

type TypeReference struct {
	TypeName   *AstObject
	TypeValue  *AstObject
	Types      []LiteralType
	Parameters []Parameter
	Kind       int
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
	if !IsClassDeclaration(st) {
		panic("Expected a class declaration")
	}

	return st.Name.EscapedText
}

func (st *Statement) HasExportModifier() bool {
	for _, modifier := range st.Modifiers {
		if modifier.Kind == 92 {
			return true
		}
	}

	return false
}

func (st *Statement) HasDefaultModifier() bool {
	for _, modifier := range st.Modifiers {
		if modifier.Kind == 87 {
			return true
		}
	}

	return false
}

func (st *Statement) HasHeritageClauses() bool {
	return len(st.HeritageClauses) > 0
}

func (sf *SourceFile) GetImportPath(key string) string {
	if sf.importsResolved {
		sf.resolveImports()
	}

	return sf.imports[key]
}

func (sf *SourceFile) resolveImports() {
	if len(sf.Statements) == 0 {
		return
	}

	for _, st := range sf.Statements {
		if !IsImportDeclaration(&st) {
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

func (sf *SourceFile) HasClassDeclaration() bool {
	for _, statement := range sf.Statements {
		if IsClassDeclaration(&statement) {
			return true
		}
	}

	return false
}

func (sf *SourceFile) GetMembersOfType(typeName string) []Member {
	// is this an imported typeName
	importLibrary := sf.GetImportPath(typeName)
	if len(importLibrary) > 0 {
		return sf.GetMembersOfTypeFromLibrary(importLibrary, typeName)
	}

	for _, statement := range sf.Statements {
		if IsInterfaceDeclaration(&statement) {
			if statement.Name != nil && typeName == statement.Name.EscapedText {
				return statement.Members
			}
		}
	}

	return nil
}

func (sf *SourceFile) GetMembersOfTypeFromLibrary(importLibrary string, typeName string) []Member {
	return nil
}
