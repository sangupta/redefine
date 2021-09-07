package com.sangupta.redefine.model;

import java.util.ArrayList;
import java.util.List;

/**
 * Represents one component definition.
 * 
 * @author sangupta
 *
 */
public class ComponentDef {
	
	public final String name;

	public final String sourcePath;
	
	public final ComponentType componentType;
	
	public String description;
	
	public final List<PropDef> props = new ArrayList<>();
	
	public ComponentDef(String name, String path, ComponentType componentType) {
		this.name = name;
		this.sourcePath = path;
		this.componentType = componentType;
	}
	
	@Override
	public String toString() {
		return "[Component: " + this.name + "; Path: " + this.sourcePath + "; Props: " + this.props.size() + "]";
	}
	
}
