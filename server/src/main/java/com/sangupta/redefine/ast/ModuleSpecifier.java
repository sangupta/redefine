package com.sangupta.redefine.ast;

public class ModuleSpecifier extends AstNode {
	
	public String text;
	
	@Override
	public String toString() {
		return "[ModuleSpecifier: " + this.text + "; Kind: " + this.kind + "]";
	}
	
}
