package com.sangupta.redefine.ast;

import java.util.ArrayList;
import java.util.List;

public class Member extends AstNode {
	
	public AstObject name;
	
	public TypeReference type;
	
	public AstObject questionToken;
	
	public final List<AstObject> jsDoc = new ArrayList<>();

	public final List<AstObject> modifiers = new ArrayList<>();
	
	public Initializer initializer;
	
	@Override
	public String toString() {
		return "[Member: " + this.name + "; Kind: " + this.kind + "]";
	}
}
