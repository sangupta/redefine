package com.sangupta.redefine.ast;

public class AstObject extends AstNode {
	
	public String escapedText;
	
	public String comment;
	
	public String text;
	
	public boolean hasExtendedUnicodeEscape;
	
	@Override
	public String toString() {
		String txt = this.text != null ? this.text : this.escapedText;
		if(txt == null) {
			if(this.comment == null) {
				return "[Kind: " + this.kind + "]";
			}
			
			return "[Kind: " + this.kind + "; Comment: " + this.comment + "]";
		}
		
		return "[Kind: " + this.kind + "; Text: " + txt  + "]";
	}

}
