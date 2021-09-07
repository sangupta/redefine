package com.sangupta.redefine.ast;

public class Element extends AstNode {

	public AstObject name;
	
	public AstObject propertyName;
	
	@Override
	public String toString() {
		return "[Element: " + this.name + "; Kind: " + this.kind + "]";
	}
	
}
