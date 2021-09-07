package com.sangupta.redefine.ast;

public class Property extends AstNode {
	
	public AstObject name;
	
	public AstObject initializer;
	
	@Override
	public String toString() {
		return "[Property: " + this.name + "; Kind: " + this.kind + "]";
	}

}
