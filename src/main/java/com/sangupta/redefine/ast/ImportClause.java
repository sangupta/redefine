package com.sangupta.redefine.ast;

public class ImportClause extends AstNode {

	public boolean isTypeOnly;
	
	public AstObject name;
	
	public NamedBindings namedBindings;
}
