package com.sangupta.redefine.ast;

public abstract class AstNode {
	
	public int kind;
	
	@Override
	public String toString() {
		return "[" + this.getClass().getName() + "; Kind=" + this.kind + "]";
	}

}
