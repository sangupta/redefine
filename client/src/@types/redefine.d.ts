/**
 *
 * Redefine - UI component documentation
 * 
 * MIT License.
 * Copyright (c) 2022, Sandeep Gupta.
 * 
 * Use of this source code is governed by a MIT style license
 * that can be found in LICENSE file in the code repository.
 * 
 **/

interface NoProps {

}

/**
 * Defines the type of a function parameter.
 * 
 */
interface ParamDef {
    name: string;
    type?: string;
}

/**
 * Attributes in component JSON that define a component `prop`.
 */
interface PropDef {
    name: string;
    type?: string;
    required: boolean;
    enumOf?: Array<ParamDef>
    defaultValue?: string;
    description?: string;
    returnType?: string;
    params?: Array<ParamDef>;
}

interface ComponentExample {
    name: string;
    markdown: string;
}

/**
 * Attributes in component JSON that define a `component`.
 */
interface ComponentDef {
    name: string;
    sourcePath: string;
    componentType: string;
    description: string;
    props?: Array<PropDef>
    docs: string;
    url?: string;

    // following are the evaluated properties
    examples: Array<ComponentExample>; // holds the markdown for each section of example
    playground: string; // the code block that defines the playground
}

interface Author {
    name?: string;
    email?: string;
    url?: string;
}

interface RedefinePayload {
    title?: string;
    favicon?:string;
    description?: string;
    libDocs?: string;
    version?: string;
    homePage?: string;
    author?: Author;
    license?: string;
    components?: Array<ComponentDef>;
    customCSS?:string;
    library?:string;
}
