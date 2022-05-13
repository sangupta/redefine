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

import (
	"fmt"

	"github.com/quickjs-go/quickjs-go"
)

type Typescript struct {
	syntaxKind SyntaxKind
}

func validateObject(result quickjs.Value) {
	names, _ := result.PropertyNames()
	fmt.Println()
	// kind := result.Get("kind").String()
	// fmt.Println("kind:: " + kind)
	fmt.Println(names)
	fmt.Println()
}

func (ts *Typescript) getSourceFile(result quickjs.Value) *SourceFile {
	defer result.Free()

	sourceFile := SourceFile{}
	if result.IsNull() {
		return &sourceFile
	}

	kind := getInt(result.Get("kind"))
	if kind != ts.syntaxKind.SourceFile {
		return nil
	}

	sourceFile.Kind = kind
	sourceFile.Statements = getArray(result.Get("statements"), sourceFile.Statements, ts.getStatement)

	return &sourceFile
}

func getArray[T any](result quickjs.Value, slice []T, converter func(quickjs.Value) *T) []T {
	defer result.Free()

	if result.IsNull() {
		return slice
	}

	if !result.IsArray() {
		return slice
	}

	length := int(result.Len())
	for i := 0; i < length; i++ {
		item := result.GetByUint32(uint32(i))
		stmt := converter(item)
		if stmt != nil {
			slice = append(slice, *stmt)
		}
	}

	return slice
}

func (ts *Typescript) getStatement(result quickjs.Value) *Statement {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	validateObject(result)

	stmt := Statement{
		ImportClause:    ts.getImportClause(result.Get("ImportClause")),
		ModuleSpecifier: ts.getModuleSpecifier(result.Get("moduleSpecifier")),
		Name:            ts.getAstObject(result.Get("name")),
		Body:            ts.getBlock(result.Get("Block")),
		Expression:      ts.getExpression(result.Get("Expression")),
		Kind:            getInt(result.Get("kind")),
	}

	stmt.HeritageClauses = getArray(result.Get("heritageClauses"), stmt.HeritageClauses, ts.getHeritageClause)
	stmt.Modifiers = getArray(result.Get("modifiers"), stmt.Modifiers, ts.getAstObject)
	stmt.Members = getArray(result.Get("members"), stmt.Members, ts.getMember)
	stmt.JsDoc = getArray(result.Get("jsDoc"), stmt.JsDoc, ts.getAstObject)
	stmt.Parameters = getArray(result.Get("parameters"), stmt.Parameters, ts.getParameter)

	return &stmt
}

func (ts *Typescript) getMember(result quickjs.Value) *Member {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	validateObject(result)

	member := Member{
		Name:          ts.getAstObject(result.Get("name")),
		TypeReference: ts.getTypeReference(result.Get("type")),
		QuestionToken: ts.getAstObject(result.Get("questionToken")),
		Initializer:   ts.getInitializer(result.Get("initializer")),
		Kind:          getInt(result.Get("kind")),
	}

	member.JsDoc = getArray(result.Get("jsDoc"), member.JsDoc, ts.getAstObject)
	member.Modifiers = getArray(result.Get("modifiers"), member.Modifiers, ts.getAstObject)

	return &member
}

func (ts *Typescript) getHeritageClause(result quickjs.Value) *HeritageClause {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	kind := getInt(result.Get("kind"))
	if kind != ts.syntaxKind.HeritageClause {
		return nil
	}

	hc := HeritageClause{
		Kind: getInt(result.Get("kind")),
	}
	hc.Types = getArray(result.Get("types"), hc.Types, ts.getTypeValue)

	return &hc
}

func (ts *Typescript) getTypeReference(result quickjs.Value) *TypeReference {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	validateObject(result)

	tr := TypeReference{
		Kind:      getInt(result.Get("kind")),
		TypeName:  ts.getAstObject(result.Get("typeName")),
		TypeValue: ts.getAstObject(result.Get("typeValue")),
	}
	tr.Types = getArray(result.Get("types"), tr.Types, ts.getLiteralType)
	tr.Parameters = getArray(result.Get("parameters"), tr.Parameters, ts.getParameter)

	return &tr
}

func (ts *Typescript) getLiteralType(result quickjs.Value) *LiteralType {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	validateObject(result)

	lt := LiteralType{
		Kind:    getInt(result.Get("kind")),
		Literal: ts.getAstObject(result.Get("literal")),
	}

	return &lt
}

func (ts *Typescript) getInitializer(result quickjs.Value) *Initializer {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	in := Initializer{
		Kind: getInt(result.Get("kind")),
	}
	in.Properties = getArray(result.Get("properties"), in.Properties, ts.getProperty)

	return &in
}

func (ts *Typescript) getProperty(result quickjs.Value) *Property {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	validateObject(result)

	pr := Property{
		Kind:        getInt(result.Get("kind")),
		Initializer: ts.getAstObject(result.Get("initializer")),
		Name:        ts.getAstObject(result.Get("name")),
	}

	return &pr
}

func (ts *Typescript) getParameter(result quickjs.Value) *Parameter {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	validateObject(result)

	p := Parameter{
		Name:          ts.getAstObject(result.Get("name")),
		TypeReference: ts.getTypeReference(result.Get("type")),
		Kind:          getInt(result.Get("kind")),
	}

	return &p
}

func (ts *Typescript) getBlock(result quickjs.Value) *Block {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	validateObject(result)

	bl := Block{
		Kind: getInt(result.Get("kind")),
	}
	bl.Statements = getArray(result.Get("statements"), bl.Statements, ts.getStatement)

	return &bl
}

func (ts *Typescript) getModuleSpecifier(result quickjs.Value) *ModuleSpecifier {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	validateObject(result)

	ms := ModuleSpecifier{
		Text: getString(result.Get("text")),
		Kind: getInt(result.Get("kind")),
	}

	return &ms
}

func (ts *Typescript) getExpression(result quickjs.Value) *Expression {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	validateObject(result)

	ex := Expression{
		Name: ts.getAstObject(result.Get("name")),
	}

	return &ex
}

func (ts *Typescript) getImportClause(result quickjs.Value) *ImportClause {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	validateObject(result)

	kind := getInt(result.Get("kind"))
	if kind != ts.syntaxKind.ImportDeclaration {
		return nil
	}

	// isTypeOnly
	ic := ImportClause{
		IsTypeOnly:    getBool(result.Get("isTypeOnly")),
		Name:          ts.getAstObject(result.Get("name")),
		Kind:          kind,
		NamedBindings: ts.getNameBindings(result.Get("namedBindings")),
	}

	return &ic
}

func (ts *Typescript) getAstObject(result quickjs.Value) *AstObject {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	kind := getInt(result.Get("kind"))
	if kind != ts.syntaxKind.ImportDeclaration {
		return nil
	}

	obj := AstObject{
		EscapedText:              getString(result.Get("escapedText")),
		Comment:                  getString(result.Get("comment")),
		Text:                     getString(result.Get("text")),
		HasExtendedUnicodeEscape: getBool(result.Get("hasExtendedUnicodeEscape")),
		Kind:                     getInt(result.Get("kind")),
	}

	return &obj
}

func (ts *Typescript) getNameBindings(result quickjs.Value) *NamedBindings {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	validateObject(result)

	nb := NamedBindings{
		Name: ts.getAstObject(result.Get("name")),
		Kind: getInt(result.Get("kind")),
	}

	return &nb
}

func (ts *Typescript) getTypeValue(result quickjs.Value) *TypeValue {
	defer result.Free()

	if result.IsNull() {
		return nil
	}

	validateObject(result)

	tv := TypeValue{
		Expression: ts.getExpression(result.Get("expression")),
		Kind:       getInt(result.Get("kind")),
	}
	tv.TypeArguments = getArray(result.Get("typeArguments"), tv.TypeArguments, ts.getTypeReference)

	return &tv
}

// PARSE PRIMITIVE VALUES

func getString(result quickjs.Value) string {
	defer result.Free()

	if result.IsNull() {
		return ""
	}

	if result.IsString() {
		return result.String()
	}

	panic("Object is not a string")
}

func getInt(result quickjs.Value) int {
	defer result.Free()

	if !result.IsNumber() {
		fmt.Println(result.IsBigDecimal())
		fmt.Println(result.IsBigFloat())
		fmt.Println(result.IsBigInt())
		fmt.Println(result.IsString())
		fmt.Println(result.IsArray())
		fmt.Println(result.String())
		fmt.Println(result.PropertyNames())
		return 0
	}

	return int(result.Int32())
}

func getBool(result quickjs.Value) bool {
	defer result.Free()

	if result.IsNull() {
		return false
	}

	if result.IsBool() {
		return result.Bool()
	}

	return false
}
