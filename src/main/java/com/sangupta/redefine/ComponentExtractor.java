package com.sangupta.redefine;

import java.util.ArrayList;
import java.util.List;

import com.sangupta.redefine.ComponentUtils.ComponentTypeWrapper;
import com.sangupta.redefine.ast.Member;
import com.sangupta.redefine.ast.SourceFile;
import com.sangupta.redefine.ast.Statement;
import com.sangupta.redefine.ast.TypeReference;
import com.sangupta.redefine.model.ComponentDef;
import com.sangupta.redefine.model.PropDef;

public class ComponentExtractor {

	public static List<ComponentDef> extactComponents(String name, String path, SourceFile ast) {
		final List<ComponentDef> components = new ArrayList<>();
		
		// find all classes that are probable component candidates
		for(Statement statement : ast.statements) {
			
			// detect class based components first
			if(TypescriptUtils.isClassDeclaration(statement)) {
				// case 1: has export keyword, and extend react.component or just component from both react library
				if(statement.hasExportModifier() && statement.hasHeritageClauses() && TypescriptUtils.hasMethod(statement, "render")) {
					ComponentTypeWrapper componentTypeWrapper = ComponentUtils.detectComponentType(ast, statement);
					if(componentTypeWrapper == null || componentTypeWrapper.componentType == null) {
						continue;
					}
					
					ComponentDef def = new ComponentDef(statement.getClassName(), path, componentTypeWrapper.componentType);
					
					// add component description
					def.description = TypescriptUtils.getJsDoc(statement.jsDoc);
					
					// find component props
					if(!componentTypeWrapper.type.typeArguments.isEmpty()) {
						// the first argument specifies the props
						TypeReference typeReference = componentTypeWrapper.type.typeArguments.get(0);
						
						// find all members of the interface from the source file
						List<Member> members = ast.getMembersOfType(typeReference.typeName.escapedText);
						
						// document all the members as this components props
						if(members != null) {
							for(Member member : members) {
								def.props.add(getComponentProp(member));
							}
						}
					}
					
					// add component to definition list
					components.add(def);
				}
				
				continue;
			}
		}
		
		return components;
	}

	private static PropDef getComponentProp(Member member) {
		PropDef def = new PropDef();
		
		def.name = member.name.escapedText;
		
		if(member.type.typeName != null) {
			def.type = member.type.typeName.escapedText;
		} else {
			if(TypescriptUtils.isUnionType(member.type)) {
				def.type = "$enum";
			}
		}
		
		if(member.questionToken != null) {
			def.required = false;
		}
		def.description = TypescriptUtils.getJsDoc(member.jsDoc);
		
		return def;
	}

}
