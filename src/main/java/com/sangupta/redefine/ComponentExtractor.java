package com.sangupta.redefine;

import java.util.ArrayList;
import java.util.List;

import com.sangupta.redefine.ComponentUtils.ComponentTypeWrapper;
import com.sangupta.redefine.ast.Member;
import com.sangupta.redefine.ast.Parameter;
import com.sangupta.redefine.ast.SourceFile;
import com.sangupta.redefine.ast.Statement;
import com.sangupta.redefine.ast.TypeReference;
import com.sangupta.redefine.model.ComponentDef;
import com.sangupta.redefine.model.ComponentType;
import com.sangupta.redefine.model.ParamDef;
import com.sangupta.redefine.model.PropDef;

public class ComponentExtractor {
	
//	private static final Gson GSON = new GsonBuilder().setPrettyPrinting().create();

	public static List<ComponentDef> extactComponents(String name, String path, SourceFile ast) {
		final List<ComponentDef> components = new ArrayList<>();
		
		// find all classes/functions that are probable component candidates
		for(Statement statement : ast.statements) {			
			// detect class based components first
			if(TypescriptUtils.isClassDeclaration(statement)) {
				ComponentDef def = extractClassBasedComponents(path, ast, statement);
				if(def != null) {
					// add component to definition list
					components.add(def);
				}
				continue;
			}
			
			// detect function based components
			if(TypescriptUtils.isFunctionDeclaration(statement)) {
				ComponentDef def = extractFunctionBasedComponent(path, ast, statement);
				if(def != null) {
					// add component to definition list
					components.add(def);
				}
				continue;
			}
		}
		
		return components;
	}
	
	/**
	 * Extract function based components
	 * 
	 * @param path
	 * @param ast
	 * @param statement
	 * @param components
	 */
	private static ComponentDef extractFunctionBasedComponent(String path, SourceFile ast, Statement statement) {
		if(!statement.hasExportModifier()) {
			return null;
		}
		
		if(!TypescriptUtils.returnsJsxFragement(statement.body)) {
			return null;
		}
		
		ComponentDef def = new ComponentDef(statement.name.escapedText, path, ComponentType.REACT_FUNCTION_COMPONENT);
		
		return def;
	}

	/**
	 * Extract class based components
	 * 
	 * @param path
	 * @param ast
	 * @param statement
	 * @param components
	 */
	private static ComponentDef extractClassBasedComponents(String path, SourceFile ast, Statement statement) {
		if(!statement.hasExportModifier()) {
			return null;
		}
		
		if(!statement.hasHeritageClauses()) {
			return null;
		}
		
		// case 1: has export keyword, and extend react.component or just component from both react library
		if(!TypescriptUtils.hasMethod(statement, "render")) {
			return null;
		}
		
		// this is a class based component
		ComponentTypeWrapper componentTypeWrapper = ComponentUtils.detectComponentType(ast, statement);
		if(componentTypeWrapper == null || componentTypeWrapper.componentType == null) {
			return null;
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

		return def;
	}

	private static PropDef getComponentProp(Member member) {
		PropDef def = new PropDef();
		
		def.name = member.name.escapedText;
		
		System.out.println(member.type);
		
		if(member.type.typeName != null) {
			def.type = member.type.typeName.escapedText;
		} else {
			String memberType = TypescriptUtils.getType(member.type);
			if(!TypescriptUtils.UNKNOWN.equals(memberType) && !TypescriptUtils.isFunctionType(member.type)) {
				def.type = memberType;
			} else {
				if(TypescriptUtils.isUnionType(member.type)) {
					def.type = "$enum";
				} else if(TypescriptUtils.isFunctionType(member.type)) {
					def.type = "$function";
					
					// build the type using definitions
					for(Parameter param: member.type.parameters) {
						def.params.add(new ParamDef(param.name.escapedText, TypescriptUtils.getType(param.type)));
					}
					def.returnType = TypescriptUtils.getType(member.type.type);
				}
			}

		}
		
		if(member.questionToken != null) {
			def.required = false;
		}
		def.description = TypescriptUtils.getJsDoc(member.jsDoc);
		
		return def;
	}

}
