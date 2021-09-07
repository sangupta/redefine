package com.sangupta.redefine.ast;

import java.util.ArrayList;
import java.util.List;

public class TypeReference extends AstNode {

	public AstObject typeName;
	
	public AstObject type;
	
	public final List<LiteralType> types = new ArrayList<>();
	
	public final List<Parameter> parameters = new ArrayList<>();
	

}
