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

export interface NoProps {

}

/**
 * Defines the type of a function parameter.
 * 
 */
export interface ParamDef {
    name: string;
    type?: string;
}

/**
 * Attributes in component JSON that define a component `prop`.
 */
export interface PropDef {
    name: string;
    type?: string;
    required: boolean;
    enumOf?: Array<ParamDef>
    defaultValue?: string;
    description?: string;
    returnType?: string;
    params?: Array<ParamDef>;
}

/**
 * Attributes in component JSON that define a `component`.
 */
export interface ComponentDef {
    name: string;
    sourcePath: string;
    componentType: string;
    description: string;
    props?: Array<PropDef>
    docs: string;
}

/**
 * Method to sort components on name.
 * 
 * @param a 
 * @param b 
 * @returns 
 */
export function componentSorter(a: ComponentDef, b: ComponentDef) {
    if (!a || !b) {
        console.warn('sorting received undefined');
        return 0;
    }

    return a.name.localeCompare(b.name);
}

/**
 * Method to sort component props on name.
 * 
 * @param a 
 * @param b 
 * @returns 
 */
export function propsSorter(a: PropDef, b: PropDef) {
    if (!a || !b) {
        console.warn('sorting received undefined');
        return 0;
    }

    return a.name.localeCompare(b.name);
}