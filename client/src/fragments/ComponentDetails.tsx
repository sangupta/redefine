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
import styled from 'styled-components';
import ReactMarkdown from 'react-markdown';
import { propsSorter } from '../Utils';
import CodePlayground from './CodePlayground';
import ComponentPlayground from './ComponentPlayground';
import TabContainer from './Tabs';

interface ComponentDetailsProps {
    component: ComponentDef;
    example?: ComponentExample;
}

const ExampleDisplay = styled.div`
    margin: 20px 0;
`;

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
            <h5>Props</h5>
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
    renderExamples = () => {
        const { component, example } = this.props;
        const markdown = example?.markdown || component.docs || '';

        if (markdown) {

            return <ExampleDisplay>
                <h3>{example?.name}</h3>
                <ReactMarkdown className='markdown-docs' components={{
                    code({ node, inline, className, children, ...props }) {
                        const match = /language-(\w+)/.exec(className || '')
                        const sourceCode = '<>\n' + String(children).trim() + '\n</>'

                        return (!inline && match)
                            ? <CodePlayground source={sourceCode} />
                            : <code className={className} {...props}>{children}</code>
                    }
                }}>{markdown}</ReactMarkdown>
            </ExampleDisplay>
        }

        return null
    }

    render() {
        const { component, example } = this.props;

        const tabs = [];
        let exampleTab = 0;
        if (component.props && component.props.length > 0) {
            exampleTab++;

            tabs.push({
                name: 'Props',
                component: this.renderProps()
            });
        }

        if (component.docs && component.docs.trim().length > 0) {
            exampleTab++;

            tabs.push({
                name: 'Examples',
                component: this.renderExamples()
            });
        }

        if (component.playground) {
            tabs.push({
                name: 'Playground',
                component: <ComponentPlayground component={component} />
            });
        }

        return <DetailsContainer>
            <ComponentName>{component.name}</ComponentName>
            <ComponentSourceFile>{component.sourcePath}</ComponentSourceFile>

            <ReactMarkdown>{component.description}</ReactMarkdown>

            <TabContainer key={component.name + '-' + example?.name} tabs={tabs} selectedTab={example ? exampleTab - 1 : 0} />
        </DetailsContainer>
    }

}

const DetailsContainer = styled.div`
    height: 100%;
`;

const ComponentName = styled.h1`
    font-size: 32px;
    line-height: 44px;
`;

const ComponentSourceFile = styled.pre`
    margin-bottom: 16px;
    font-size: 14px;
    direction: ltr;
    unicode-bidi: bidi-override;
    overflow: auto;
`;
