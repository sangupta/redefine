package com.sangupta.redefine.ast;

import java.util.ArrayList;
import java.util.List;

public class NamedBindings extends AstNode {

	public AstObject name;
	
	public final List<Element> elements = new ArrayList<>();
}
