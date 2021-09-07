package com.sangupta.redefine;

import java.io.File;
import java.lang.reflect.Field;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.ParameterizedType;
import java.lang.reflect.Type;
import java.util.Collection;
import java.util.List;
import java.util.Map;

import com.eclipsesource.v8.NodeJS;
import com.eclipsesource.v8.V8Object;
import com.eclipsesource.v8.utils.V8ObjectUtils;
import com.sangupta.jerry.util.AssertUtils;
import com.sangupta.jerry.util.ReflectionUtils;
import com.sangupta.redefine.ast.SourceFile;

/**
 * Class that extracts the AST of the code by using the Typescript Compiler executing
 * it inside a J2V8 environment. 
 * 
 * @author sangupta
 *
 */
public class AstExtractor {

	final NodeJS nodeJS;

	final V8Object typescript;

	final V8Object compilerOptions;
	
	final V8Object moduleKind;

	public AstExtractor() {
		this.nodeJS = NodeJS.createNodeJS();
		this.typescript = nodeJS.require(new File("/Users/sangupta/git/sangupta/ts-ast/node_modules/typescript"));
		
		this.moduleKind = typescript.getObject("ScriptTarget");
		final Integer system = moduleKind.getInteger("Latest");
		
		this.compilerOptions = new V8Object(nodeJS.getRuntime());
		this.compilerOptions.add("module", system);
	}
	
	public void release() {
		this.compilerOptions.release();
		this.moduleKind.release();
		this.typescript.release();
		this.nodeJS.release();
	}
	

	public SourceFile getAst(String fileName, String code) {
		V8Object result = (V8Object) typescript.executeJSFunction("createSourceFile", fileName, code, compilerOptions, true);
		Map<String, ? super Object> result2 = V8ObjectUtils.toMap(result);
		while (nodeJS.isRunning()) {
			nodeJS.handleMessage();
		}
		
		result.release();
		
		return convertToAst(result2, SourceFile.class);
	}

	@SuppressWarnings({ "rawtypes", "unchecked" })
	private <T> T convertToAst(Map<String, Object> map, Class<T> clazz) {
		if(map == null) {
			return null;
		}
		
		T instance = newInstance(clazz);

		if(AssertUtils.isEmpty(map)) {
			return instance;
		}
		
		// check fields
		List<Field> fields = ReflectionUtils.getAllFields(clazz);
		if(AssertUtils.isEmpty(fields)) {
			return instance;
		}
		
		// iterate fields
		for(Field field : fields) {
			// find variable name
			String name = field.getName();

			// find value from object map
			Object value = map.get(name);
			if(value == null) {
				// nothing can be done here
				continue;
			}

			// another way to check undefined
			if("class com.eclipsesource.v8.V8Object$Undefined".equals(value.getClass().toString())) {
				continue;
			}
			
			// check for primitives
			if(field.getType().isPrimitive()) {
				setFieldValue(field, instance, value);
				continue;
			}
			
			// check for string
			if(String.class.isAssignableFrom(field.getType())) {
				setFieldValue(field, instance, value.toString());
				continue;
			}
			
			// check for collection
			if(value instanceof Collection) {
				Type type = field.getGenericType();
				Type innerType = null;
				if (type instanceof ParameterizedType) {
					ParameterizedType pt = (ParameterizedType) type;
//					System.out.println("raw type: " + pt.getRawType());
//		            System.out.println("owner type: " + pt.getOwnerType());
//		            System.out.println("actual type args:");
		            innerType = pt.getActualTypeArguments()[0];
		            
		            Collection toPopulate = getCollection(instance, field);
		            Collection<? extends Object> actual = (Collection<?>) value;
		            for(Object item : actual) {
		            	Object arrayItem = convertToAst((Map) item, getClass(innerType));
		            	
		            	toPopulate.add(arrayItem);
		            }
				}
				
				continue;
			}
			
			// this is a pure single object
			setFieldValue(field, instance, convertToAst((Map) value, field.getType()));
		}
		
		return instance;
	}
	
	private Collection<?> getCollection(Object instance, Field field) {
		try {
			return (Collection<?>) field.get(instance);
		} catch (IllegalArgumentException | IllegalAccessException e) {
			throw new RuntimeException(e);
		}
	}

	private Class<?> getClass(Type type) {
		try {
			return Class.forName(type.getTypeName());
		} catch (ClassNotFoundException e) {
			throw new RuntimeException(e);
		}
	}
	
	private <T> T newInstance(Class<T> clazz) {
		try {
			return clazz.getDeclaredConstructor().newInstance();
		} catch (InstantiationException | IllegalAccessException | IllegalArgumentException | InvocationTargetException | NoSuchMethodException | SecurityException e) {
			throw new RuntimeException(e);
		} 
	}

	private void setFieldValue(Field field, Object instance, Object value) {
		try {
			ReflectionUtils.bindValue(field, instance, value);
		} catch (IllegalArgumentException | IllegalAccessException e) {
			throw new RuntimeException(e);
		}			
	}
}
