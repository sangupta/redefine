package com.sangupta.redefine.model;

import java.util.ArrayList;
import java.util.List;

public class PropDef {

	public String name;
	
	public String type;
	
	public boolean required;
	
	public String defaultValue;
	
	public String description;
	
	public String returnType;
	
	public final List<ParamDef> params = new ArrayList<>();
	
}
