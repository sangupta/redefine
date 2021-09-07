package com.sangupta.redefine;

import com.sangupta.redefine.ast.HeritageClause;
import com.sangupta.redefine.ast.SourceFile;
import com.sangupta.redefine.ast.Statement;
import com.sangupta.redefine.ast.Type;
import com.sangupta.redefine.model.ComponentType;

public class ComponentUtils {

	/**
	 * Detect the component type from the {@link Statement} that represents
	 * a class. This does not detect function based components.
	 * 
	 * @param ast
	 * @param statement
	 * @return
	 */
	public static ComponentTypeWrapper detectComponentType(SourceFile ast, Statement statement) {
		if(!TypescriptUtils.isClassDeclaration(statement)) {
			throw new RuntimeException("Not a class statement");
		}
		
		for(HeritageClause clause : statement.heritageClauses) {
			for(Type type : clause.types) {
				String expression = type.expression.expression.escapedText;
				String name = type.expression.name.escapedText;
				
				// check for `extends React.Component`
				if("Component".equals(name) && isReactImport(ast, expression)) {
					return new ComponentTypeWrapper(ComponentType.REACT_CLASS_COMPONENT, clause, type);
				}
				
			}
		}
		
		return null;
	}

	/**
	 * Check if the import statement reflects that this `key` comes
	 * from react/preact/stencil library.
	 * 
	 * @param ast
	 * @param name
	 * @return
	 */
	private static boolean isReactImport(SourceFile ast, String name) {
		for(Statement statement : ast.statements) {
			if(!TypescriptUtils.isImportDeclaration(statement)) {
				continue;
			}
			
			// check in imports
			if("react".equalsIgnoreCase(ast.getImportPath(name))) {
				return true;
			}
			
		}
		return false;
	}

	public static class ComponentTypeWrapper {
		public final ComponentType componentType;
		public final HeritageClause heritageClause;
		public final Type type;
		
		public ComponentTypeWrapper(ComponentType componentType, HeritageClause heritageClause, Type type) {
			this.componentType = componentType;
			this.heritageClause = heritageClause;
			this.type = type;
		}
		
	}
}
