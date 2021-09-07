package com.sangupta.redefine;

import java.util.List;

import com.sangupta.jerry.util.AssertUtils;
import com.sangupta.jerry.util.StringUtils;
import com.sangupta.redefine.ast.AstNode;
import com.sangupta.redefine.ast.AstObject;
import com.sangupta.redefine.ast.Member;
import com.sangupta.redefine.ast.Statement;

public class TypescriptUtils {
	
	public static String getNodeType(AstNode node) {
		switch(node.kind) {
			case 253:
				return "ClassDeclaration";
			case 254:
				return "InterfaceDeclaration";
			case 262:
				return "ImportDeclaration";
		}
		
		return "Unknown";
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

	
}
