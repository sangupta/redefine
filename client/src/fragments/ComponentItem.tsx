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

interface ComponentItemProps {
    component: ComponentDef;

    onSelect: (def: ComponentDef, example?: ComponentExample) => void;
}

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
            <a href='#' className='list-group-item list-group-item-action py-2 lh-tight pointer' onClick={this.handleClick}>
                {component.name}
            </a>
            {examples}
        </>
    }

}

interface ExampleItemProps {
    example: ComponentExample;
    onSelect: (example: ComponentExample) => void;
}

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

        return <a href='#' className='list-group-item list-group-item-action py-2 lh-tight pointer' onClick={this.handleExampleSelect}>
            {example.name}
        </a>
    }
}
