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

interface ComponentItemProps {
    component: ComponentDef;

    onSelect: (def: ComponentDef, example?: ComponentExample) => void;
}

const ComponentLink = styled.a`
    padding: 8px;
    color: #212529;
    cursor: pointer;
    text-decoration: none;

    :hover {
        background-color: #f8f9fa;
        color: #495057;
        text-decoration: none;
    }
`;

export default class ComponentItem extends React.Component<ComponentItemProps> {

    handleClick = () => {
        this.props.onSelect(this.props.component);
    }

    handleExampleSelect = (example: ComponentExample) => {
        this.props.onSelect(this.props.component, example);
    }

    render() {
        const { component } = this.props;
        const examples: Array<React.ReactElement> = [];

        if (component.examples && component.examples.length > 0) {
            component.examples.forEach((example, index) => {
                examples.push(<ExampleItem key={example.name || 'example-' + index} example={example} onSelect={this.handleExampleSelect} />)
            });
        }

        return <>
            <ComponentLink href='#' onClick={this.handleClick}>
                {component.name}
            </ComponentLink>
            {examples}
        </>
    }

}

interface ExampleItemProps {
    example: ComponentExample;
    onSelect: (example: ComponentExample) => void;
}

const ExampleLink = styled.a`
    color: #212529;
    padding: 8px;
    padding-left: 16px;
    text-decoration: none;
    display: block;
    position: relative;
`;

class ExampleItem extends React.Component<ExampleItemProps> {

    handleExampleSelect = (e: React.MouseEvent) => {
        e.preventDefault();
        e.stopPropagation();

        const { example, onSelect } = this.props;
        if (onSelect) {
            onSelect(example);
        }
    }

    render() {
        const { example } = this.props;

        return <ExampleLink href='#' onClick={this.handleExampleSelect}>
            {example.name}
        </ExampleLink>
    }
}
