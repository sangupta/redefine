package com.sangupta.redefine.ast;

import java.util.ArrayList;
import java.util.List;

import com.sangupta.redefine.TypescriptUtils;

public class Statement extends AstNode {
	
	public ImportClause importClause;
	
	public ModuleSpecifier moduleSpecifier;
	
	public AstObject name;
	
	public Block body;
	
	public Expression expression;
	
	public final List<HeritageClause> heritageClauses = new ArrayList<>();

	public final List<AstObject> modifiers = new ArrayList<>();
	
	public final List<Member> members = new ArrayList<>();
	
	public final List<AstObject> jsDoc = new ArrayList<>();
	
	public final List<Parameter> parameters = new ArrayList<>();
	
	public String getClassName() {
		if(!TypescriptUtils.isClassDeclaration(this)) {
			throw new RuntimeException("Expected a class declaration");
		}
		
		return this.name.escapedText;
	}
	
	public boolean hasExportModifier() {
		for(AstObject modifier : this.modifiers) {
			if(modifier.kind == 92) {
				return true;
			}
		}
		
		return false;
	}
	
	public boolean hasDefaultModifier() {
		for(AstObject modifier : this.modifiers) {
			if(modifier.kind == 87) {
				return true;
			}
		}
		
		return false;		
	}
	
	public boolean hasHeritageClauses() {
		return !this.heritageClauses.isEmpty();
	}
	
	@Override
	public String toString() {
		return "[" + TypescriptUtils.getNodeType(this) + "]";
	}
	
}
