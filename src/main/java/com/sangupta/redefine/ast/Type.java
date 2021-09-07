package com.sangupta.redefine.ast;

import java.util.ArrayList;
import java.util.List;

public class Type extends AstNode {
	
	public Expression expression;
	
	public final List<TypeReference> typeArguments = new ArrayList<>();

}
