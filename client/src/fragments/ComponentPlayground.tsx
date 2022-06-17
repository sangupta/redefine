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

import React from 'react';
import { propsSorter } from '../Utils';
import { default as Jsx2Json } from 'simplified-jsx-to-json';
import CodePlayground from './CodePlayground';

interface ComponentPlaygroundProps {
    component: ComponentDef;
}

interface ComponentPlaygroundState {
    code: string;
    propValues: { [key: string]: any };
}

export default class ComponentPlayground extends React.Component<ComponentPlaygroundProps, ComponentPlaygroundState> {

    propFields: Array<React.ReactChild> | undefined = undefined;

    propMap: { [key: string]: PropDef } = {};

    constructor(props: ComponentPlaygroundProps) {
        super(props);

        this.state = {
            code: props.component.playground,
            propValues: {}
        }
    }

    renderPropsFields = () => {
        if (this.propFields) {
            return this.propFields;
        }

        const { component } = this.props;
        if (!component || !component.props) {
            return null;
        }
        // sort
        component.props = component.props.sort(propsSorter)

        // render the element for each type of form element
        const propFields: Array<React.ReactChild> = [];
        for (let index = 0; index < component.props.length; index++) {
            const prop = component.props[index];

            // save in propMap
            this.propMap[prop.name] = prop;

            // create propFields
            if (prop.type === 'boolean') {
                propFields.push(<React.Fragment key={prop.name}>
                    <input
                        type='checkbox'
                        name={prop.name}
                        onChange={(e) => {
                            const values: any = { ...this.state.propValues };
                            values[prop.name] = e.target.checked;

                            this.setState({ propValues: values });
                        }} />&nbsp;{prop.name}
                    <br />
                </React.Fragment>)

                continue;
            }

            if (prop.type === 'string') {
                propFields.push(<React.Fragment key={prop.name}>
                    <input
                        type='text'
                        name={prop.name}
                        placeholder={prop.name}
                        onChange={(e) => {
                            const values: any = { ...this.state.propValues };
                            values[prop.name] = e.target.value;

                            this.setState({ propValues: values });
                        }} />
                    <br />
                </React.Fragment>)

                continue;
            }
        }

        this.propFields = propFields;
        return this.propFields;
    }

    render() {
        const { component } = this.props;
        if (!component.props || component.props.length === 0) {
            return null;
        }

        const { code, propValues } = this.state;
        // console.log('re-render: ', propValues);

        const updatedCode = this.updateCode(component, code, propValues);
        // render the preview
        return <>
            <CodePlayground key={updatedCode.length} source={updatedCode} editable={false} />
            {this.renderPropsFields()}
        </>
    }

    updateCode(component: ComponentDef, code: string, propValues: any): string {
        const keys = Object.keys(propValues);
        if (keys.length === 0) {
            return code;
        }

        // we need to modify the attributes of each prop in there
        const json = Jsx2Json(code);

        const first = json[0];

        const element = first[0];
        const props: any = first[1];
        const children = first[2];

        // update props
        if (element === component.name) {
            const updated = { ...props, ...propValues };
            code = json2jsx(element, updated, children, this.propMap);
            // console.log(code);
        }

        return code;
    }
}

function json2jsx(element: string, props: any, children: any, propMap: { [key: string]: PropDef }): string {
    let s = '';

    s = '<' + element + '';
    Object.keys(props).forEach(key => {
        const value = props[key];

        if (value !== undefined) {
            const propDef = propMap[key];
            const type = propDef.type || '';

            if (type === 'boolean' || type === 'number') {
                s += ' ' + key + '={' + value + '}';
            } else {
                s += ' ' + key + '="' + value + '"';
            }
        }
    });
    s += '>';

    s += (children || '');
    s += '</' + element + '>';

    return s;
}
