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

    onSelect: (def: ComponentDef) => void;
}

export default class ComponentItem extends React.Component<ComponentItemProps> {

    handleClick = () => {
        this.props.onSelect(this.props.component);
    }

    render() {
        const { component } = this.props;
        return <a href='#' className='list-group-item list-group-item-action py-2 lh-tight pointer' onClick={this.handleClick}>
            {component.name}
        </a>
    }

}
