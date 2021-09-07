package com.sangupta.redefine.ast;

import java.util.ArrayList;
import java.util.List;

public class TypeReference extends AstNode {

	public AstObject typeName;
	
	public final List<LiteralType> types = new ArrayList<>();
}
