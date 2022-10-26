# redefine

`redefine` is a simple documentation tool for UI components. It currently
supports React components written in Typescript. It differs from [StoryBook]
(https://storybook.js.org/) and others in the following ways:

* No compiler/bundler specific changes required
* Allows realtime editing of demo code
* Generates playground (like knobs) with a single line of code
* Generates a JSON file as docs, allowing you to fully customize your own docs UI
* Super-fast

**Work in progress**

## TODO

* Better functional component detection
* Enhance UI playground with various knobs
* UI sleekness

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
 optionally the built component library. The server can be accessed at
 http://localhost:1309.

* `build`: Exports all documentation files into a folder, so that they can
be deployed on a static file server, like Github pages or Netlify, to be
served for public consumption.

## Redefine Config

The following `redefine` section can be added to your `package.json` file
(using `redefine` as the key) or written directly to `redefine.config.json` file.

```json
{
	"src": {
		"root": "src",
		"includes": [
			"*.ts",
			"*.tsx"
		]
	},
	"docs": {
		"root": "docs",
		"includes": [
			"*.md"
		],
		"hasFrontMatter": true,
		"index": "index.md"
	},
	"build": {
		"dist": "dist",
		"publishFolder": "publish",
		"css": [
			"https://cdn.jsdelivr.net/npm/bootstrap@5.2.2/dist/css/bootstrap.min.css"
		],
		"fonts": [
			"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.2.0/css/all.min.css"
		],
		"js": [
			"demo/dist/bedrock-demo.js"
		]
	},
	"template": {
		"title": "My Component Library",
		"favicon": "myfavicon.png"
	}
}
```

# Author

* [Sandeep Gupta](https://sangupta.com)

# License

MIT License. Copyright (c) 2022, Sandeep Gupta.
