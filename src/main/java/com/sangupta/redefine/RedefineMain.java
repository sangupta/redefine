package com.sangupta.redefine;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Collection;
import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;
import java.util.Set;

import org.apache.commons.io.FileUtils;

import com.sangupta.jerry.util.AssertUtils;
import com.sangupta.jerry.util.StringUtils;
import com.sangupta.redefine.ast.SourceFile;
import com.sangupta.redefine.model.ComponentDef;

/**
 * Generate documentation for React components.
 * 
 * @author sangupta
 *
 */
public class RedefineMain {

	private static final Map<File, SourceFile> AST = new HashMap<>();

	public static void main(String[] args) throws Exception {
		final RedefineConfig config = new RedefineConfig();
		long start = System.currentTimeMillis();
		
		try {
			build(config);
		} catch(IOException e) {
			
		}
		
		long end = System.currentTimeMillis();
		System.out.println("Total time: " + (end - start) + " millis.");
	}
	
	private static void build(RedefineConfig config) throws IOException {
		final AstExtractor extractor = new AstExtractor();

		final File baseFolder = new File(config.baseFolder);

		// find all files inside the folder
		Collection<File> files = scanFolder(baseFolder, config);
		System.out.println("Found total files: " + files.size());

		// parse AST from each file
		long start = System.currentTimeMillis();
		parseAst(extractor, files, config);
		long end = System.currentTimeMillis();
		System.out.println("Total time in parsing: " + (end - start));

		// extract components from each parsed file
		final List<ComponentDef> components = new ArrayList<>();
		extractComponents(baseFolder, components, config);
		
		// print a list of all components
		for(ComponentDef def : components) {
			System.out.println(def);
		}
	}

	private static void extractComponents(File baseFolder, List<ComponentDef> components, RedefineConfig config) {
		Set<File> keys = AST.keySet();
		for(File file : keys) {
			final SourceFile ast = AST.get(file);
			
			String name = file.getName();
			String path = file.getAbsolutePath();
			path = path.substring(baseFolder.getAbsolutePath().length());
			if(AssertUtils.isEmpty(path)) {
				path = file.getName();
			}
			
			List<ComponentDef> defs = ComponentExtractor.extactComponents(name, path, ast);
			components.addAll(defs);
		}
	}

	@SuppressWarnings({ "deprecation" })
	private static void parseAst(AstExtractor extractor, Collection<File> files, RedefineConfig config) throws IOException {
		for (File file : files) {
			String name = file.getName();

			SourceFile ast = extractor.getAst(name, FileUtils.readFileToString(file));
			AST.put(file, ast);
		}
	}

	private static Collection<File> scanFolder(File baseFolder, RedefineConfig config) {
		if(baseFolder.isFile()) {
			Collection<File> files = new ArrayList<File>();
			files.add(baseFolder);
			return files;
		}
			
		Collection<File> files = FileUtils.listFiles(baseFolder, null, true);
		if (AssertUtils.isEmpty(files)) {
			return files;
		}

		// no includes, everything is included
		if (AssertUtils.isEmpty(config.include)) {
			return files;
		}

		// filter
		Iterator<File> iterator = files.iterator();
		while (iterator.hasNext()) {
			File file = iterator.next();
			String name = file.getName();

			if (!isIncludedFile(name, config)) {
				iterator.remove();
			}
		}

		return files;
	}

	private static boolean isIncludedFile(String name, RedefineConfig config) {
		for (String pattern : config.include) {
			if (StringUtils.wildcardMatch(name, pattern)) {
				return true;
			}
		}

		return false;
	}

}
