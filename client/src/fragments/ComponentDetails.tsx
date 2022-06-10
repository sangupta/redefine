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
import ReactMarkdown from 'react-markdown';
import { ComponentDef, PropDef, propsSorter } from '../Utils';
import CodePlayground from './CodePlayground';
import ComponentPlayground from './ComponentPlayground';
import TabContainer from './Tabs';

interface ComponentDetailsProps {
    component: ComponentDef;
}

export default class ComponentDetails extends React.Component<ComponentDetailsProps> {

    renderType(prop: PropDef) {
        if (!prop.type) {
            return prop.type;
        }

        switch (prop.type) {
            case '$function':
                if (prop.params?.length || prop.returnType) {
                    let paramStr = '';
                    if (prop.params && prop.params.length > 0) {
                        for (let index = 0; index < prop.params.length; index++) {
                            const param = prop.params[index];
                            if (index > 0) {
                                paramStr += ', ';
                            }
                            paramStr += param.name;
                            if (param.type && param.type !== '$unknown') {
                                paramStr += ':' + param.type;
                            }
                        }
                    }
                    return '(' + paramStr + ') => ' + prop.returnType;
                }

                return prop.type;

            case '$enum':
                if (prop.enumOf && prop.enumOf.length > 0) {
                    let combined = '';
                    for (let index = 0; index < prop.enumOf.length; index++) {
                        const paramDef = prop.enumOf[index];

                        if (index > 0) {
                            combined += ' | ';
                        }

                        if (paramDef.type === 'string') {
                            combined += '"' + paramDef.name + '"'
                        } else {
                            combined += paramDef.name
                        }
                    }

                    return combined;
                }

                return '$enum';

            default:
                return prop.type;
        }
    }

    renderProps() {
        const { component } = this.props;
        if (!component.props || component.props.length === 0) {
            return null;
        }

        component.props = component.props.sort(propsSorter)

        const rows = [];
        for (let index = 0; index < component.props.length; index++) {
            const prop = component.props[index];

            rows.push(<tr key={prop.name}>
                <td><code>{prop.name}</code></td>
                <td><code>{this.renderType(prop)}</code></td>
                <td><pre>{'' + prop.required}</pre></td>
                <td><pre>{prop.defaultValue || ''}</pre></td>
                <td>{prop.description || ''}</td>
            </tr>);
        }

        return <>
            <h5 className='props-title'>Props</h5>
            <table className='table table-striped table-bordered props-table'>
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Type</th>
                        <th>Required</th>
                        <th>Default Value</th>
                        <th>Description</th>
                    </tr>
                </thead>
                <tbody>
                    {rows}
                </tbody>
            </table>
        </>
    }

    /**
     * Render the associated markdown file. This includes rendering
     * all the editable code examples here.
     */
    renderMarkdownFile = () => {
        const { component } = this.props;
        if (component.docs) {
            return <ReactMarkdown className='markdown-docs' components={{
                code({ node, inline, className, children, ...props }) {
                    const match = /language-(\w+)/.exec(className || '')
                    const sourceCode = '<>\n' + String(children).trim() + '\n</>'

                    return (!inline && match)
                        ? <CodePlayground source={sourceCode} />
                        : <code className={className} {...props}>{children}</code>
                }
            }}>{component.docs}</ReactMarkdown>
        }

        return null
    }

    render() {
        const { component } = this.props;

        const tabs = [];
        if (component.props && component.props.length > 0) {
            tabs.push({
                name: 'Props',
                component: this.renderProps()
            });
        }

        if (component.docs && component.docs.trim().length > 0) {
            tabs.push({
                name: 'Examples',
                component: this.renderMarkdownFile()
            });
        }

        if (component.props && component.props.length > 0) {
            tabs.push({
                name: 'Playground',
                component: <ComponentPlayground component={component} />
            });
        }

        return <div className='component-details'>
            <h1 className='component-name'>{component.name}</h1>
            <pre className='component-source-path'>{component.sourcePath}</pre>

            <ReactMarkdown className='component-description'>{component.description}</ReactMarkdown>

            <TabContainer tabs={tabs} />
        </div>
    }

}
