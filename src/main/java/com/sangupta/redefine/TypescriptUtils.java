package com.sangupta.redefine;

import java.util.List;

import com.sangupta.jerry.util.AssertUtils;
import com.sangupta.jerry.util.StringUtils;
import com.sangupta.redefine.ast.AstNode;
import com.sangupta.redefine.ast.AstObject;
import com.sangupta.redefine.ast.Block;
import com.sangupta.redefine.ast.Expression;
import com.sangupta.redefine.ast.Member;
import com.sangupta.redefine.ast.Statement;

public class TypescriptUtils {
	
	public static final String UNKNOWN = "$unknown";
	
	public static String getNodeType(AstNode node) {
		switch(node.kind) {
			case 253:
				return "ClassDeclaration";
			case 252:
				return "FunctionDeclaration";
			case 254:
				return "InterfaceDeclaration";
			case 262:
				return "ImportDeclaration";
		}
		
		return "Unknown";
	}
	
	private static boolean isReturnStatement(AstNode node) {
		return node.kind == 243;
	}
	
	public static boolean isClassDeclaration(AstNode node) {
		return node.kind == 253;
	}
	
	public static boolean isMethodDeclaration(AstNode node) {
		return isSimpleMethodDeclaration(node) || isArrowMethodDeclaration(node);
	}

	public static boolean isArrowMethodDeclaration(AstNode node) {
		return node.kind == 164;
	}
	
	public static boolean isSimpleMethodDeclaration(AstNode node) {
		return node.kind == 166;
	}

	public static boolean isInterfaceDeclaration(AstNode node) {
		return node.kind == 254;
	}

	public static boolean isImportDeclaration(AstNode node) {
		return node.kind == 262;
	}

	public static boolean isPropertyAccessExpression(AstNode node) {
		return node.kind == 198;
	}

	public static boolean isExpressionWithTypeArguments(AstNode node) {
		return node.kind == 224;
	}

	public static boolean isHeritageClause(AstNode node) {
		return node.kind == 287;
	}
	
	public static boolean isUnionType(AstNode node) {
		return node.kind == 183;
	}

	public static boolean isFunctionDeclaration(AstNode node) {
		return node.kind == 252;
	}

	public static boolean hasMethod(Statement statement, String methodName) {
		if(!isClassDeclaration(statement)) {
			return false;
		}
		
		for(Member member : statement.members) {
			if(isMethodDeclaration(member) && member.name.escapedText.equals(methodName)) {
				return true;
			}
		}
		
		return false;
	}

	public static String getJsDoc(List<AstObject> jsDoc) {
		if(AssertUtils.isEmpty(jsDoc)) {
			return StringUtils.EMPTY_STRING;
		}
		
		if(jsDoc.size() == 1) {
			return jsDoc.get(0).comment;
		}
		
		StringBuilder builder = new StringBuilder(1024);
		for(AstObject doc : jsDoc) {
			builder.append(doc.comment);
			builder.append("\n");
		}
		
		return builder.toString();
	}

	public static boolean returnsJsxFragement(Block body) {
		for(Statement statement : body.statements) {
			if(isReturnStatement(statement) && isJsxElement(statement.expression)) {
				return true;
			}
		}
		
		return false;
	}

	/**
	 * Check if given expression returns a JsxElement or a JsxFragment.
	 * 
	 * @param expression
	 * @return
	 */
	private static boolean isJsxElement(Expression expression) {
		if(expression == null) {
			return false;
		}
		
		// direct JsxElement - return <component> ... </component>
		if(expression.kind == 274) {
			return true;
		}
		
		// jsx-frgament - multiple nodes - return <> ... </>
		if(expression.kind == 278) {
			return true;
		}
		
		// inside a paranthesis - return ( <> ... </>)
		if(expression.kind == 208 && isJsxElement(expression.expression)) {
			return true;
		}
		return false;
	}

	public static boolean isNumberKeyword(AstNode node) {
		return node.kind == 144;
	}
	
	public static boolean isStringKeyword(AstNode node) {
		return node.kind == 147;
	}
	
	public static boolean isBooleanKeyword(AstNode node) {
		return node.kind == 131;
	}
	
	public static boolean isVoidKeyword(AstNode node) {
		return node.kind == 113;
	}
	
	public static boolean isAnyKeyword(AstNode node) {
		return node.kind == 128;
	}
	
	public static boolean isFunctionType(AstNode node) {
		return node.kind == 175;
	}

	public static String getType(AstNode node) {
		if(node == null) {
			return null;
		}
		
		switch(node.kind) {
			case 144:
				return "number";
			case 147:
				return "string";
			case 131:
				return "boolean";
			case 113:
				return "void";
			case 175:
				return "Function";
			case 128:
				return "any";
			case 192:
				return "null";
			case 150:
				return "undefined";
			case 141:
				return "never";
		}
		
		return UNKNOWN;
	}
	
}
