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

/**
 * Processes component information as received from server.
 *  
 * @param components 
 * @returns 
 */
export function processComponentInfo(components: Array<ComponentDef>): Array<ComponentDef> {
    if (!components || components.length === 0) {
        return [];
    }

    for (let index = 0; index < components.length; index++) {
        components[index] = processComponent(components[index]);
    }

    // finally return the components itself
    return components;
}

/**
 * Two things need to happen:
 * 
 * 1. divide the content into sections
 * which is signified by line starting with a single hash character and a white space '#'
 * 2. Find the code block that has language set to `playground`
 * this code block shall be used as playground data in tabs
 * 
 * @param component 
 * @returns 
 */
function processComponent(component: ComponentDef): ComponentDef {
    if (!component) {
        return component;
    }

    if ((component.docs || '').trim().length === 0) {
        return component;
    }

    // split into lines
    const lines = component.docs.split('\n');

    // find code block with playground lang
    let start = -1, end = -1;
    for (let index = 0; index < lines.length; index++) {
        const line = lines[index].trim();

        // find start
        if (start < 0 && line === '```js:playground') {
            start = index;
            continue;
        }

        if (start >= 0 && line === '```') {
            // we are done
            end = index;
            break;
        }
    }

    // did we find the playground?
    if (start >= 0) {
        if (end < 0) {
            end = lines.length;
        }

        component.playground = '';
        for (let index = start + 1; index < end; index++) {
            component.playground += lines[index];
        }

        // delete these lines from the list of lines
        lines.splice(start, end - start + 1);
    }

    // create splits for sections
    let section = '', title = '';
    component.examples = [];
    for (let index = 0; index < lines.length; index++) {
        const line = lines[index];

        if (line.startsWith('# ') && line.trim().length >= 3) {
            // this is a new section start
            if (section.trim().length > 0) {
                const example = { name: (title || 'Default'), markdown: section };
                if (example.name.startsWith('# ')) {
                    example.name = example.name.substring(2);
                }

                component.examples.push(example);
            }

            // clear up
            section = '';
            title = line.substring(2).trim();
            continue;
        }

        // add this line to section
        section += (line + '\n');
    }

    // all lines finish, add last section
    component.examples.push({ markdown: section, name: (title || 'Default') });

    // all done
    return component;
}
