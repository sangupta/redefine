# redefine

`redefine` is a simple documentation tool for UI components. It currently
supports React components written in Typescript.

**Work in progress**

## Usage

```sh
$ redefine <action> <folder>
```

* `action`:  (optional) specify non-default actions other than generation
of `components.json` file. Available actions are described below.

* `folder`: Root folder where either `package.json` or `redefine.config.json` 
exists. `redefine` employs convention-over-configuration approach and thus, for
simple `module` projects, if you have a proper `package.json` file, there is 
no configuration needed. This allows `redefine` to be a part of your tool chain
without being intrusive.

However, if you would like to customize all or certain aspects of `redefine`,
you may create the `redefine.config.json` file. Details on all the parameters
are available below.

### Available actions

* `serve`: Starts a local server to serve the documentation files, and
 optionally the built component library.

* `build`: Exports all documentation files into a folder, so that they can
be deployed on a static file server, like Github pages or Netlify, to be
served for public consumption.

## Redefine Config

```json
{
	// the root folder under which scanning for components shall happen
	// the default values are `src`, `lib`, and `packages` in that order
	"srcFolder": "src",

	// an array of wildcard strings, that represent which files to include in scanning
	// the default values are listed below
	"includes": [ "*.js", "*.jsx", "*.ts", "*.tsx" ],

	// the folder relative to this file where markdown docs can be found
	// the default value is `docs` and shall be used if no value is specified
	"docsFolder": "docs",

	// the title to use in documentation.
	// if not present, the `name` value from package.json shall be used
	// if package.json cannot be found, the name of the root folder shall be used
	"title": "",

	// the path where the built library is available
	// this is the final binary path
	// if not value is specified, we use the `main` attribute from package.json
	"libraryPath": "",

	// the URL where the hosted version of the library can be found
	// this file is loaded dynamically in the Redefine UI 
	"libraryUrl": "",

	// specify the file which represents the contents to display on the home
	// page for the documentation. The default value is shown below, so you
	// skip specifying it. If the file is not present, a default home page is
	// built using the information from the `package.json` file
	"indexFile": "docs/index.md"
}
```

# Author

* [Sandeep Gupta](https://sangupta.com)


# License

MIT License. Copyright (c) 2022, Sandeep Gupta.
