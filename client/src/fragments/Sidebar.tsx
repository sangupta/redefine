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

import { ComponentDef, componentSorter } from './../Utils';
import ComponentItem from './ComponentItem';

interface SidebarProps {
    className?: string;
    components: Array<ComponentDef>;
    onComponentSelect: (def: ComponentDef) => void;
}

interface SidebarState {
    filter?: string;
}

class Sidebar extends React.Component<SidebarProps, SidebarState> {

    constructor(props: SidebarProps) {
        super(props);

        this.state = {
            filter: ''
        }
    }

    handleComponentSelect = (def: ComponentDef) => {
        this.props.onComponentSelect(def);
    }

    renderComponents() {
        const result = [];
        const { components } = this.props;
        let { filter = '' } = this.state;
        filter = filter.toLowerCase();

        let filtered = [...components];
        if (filter) {
            filtered = filtered.filter(item => {
                return (item.name || '').toLowerCase().includes(filter);
            });
        }

        filtered.sort(componentSorter);

        for (let index = 0; index < filtered.length; index++) {
            const component = filtered[index];
            result.push(<ComponentItem key={component.name} component={component} onSelect={this.handleComponentSelect} />)
        }
        return <div className='list-group list-group-flush border-bottom scrollarea'>
            {result}
        </div>
    }

    handleFindChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({ filter: e.target.value });
    }

    render() {
        const { components } = this.props;

        if (!components || components.length === 0) {
            return null;
        }

        return <div className={'d-flex flex-column align-items-stretch flex-shrink-0 bg-white ' + this.props.className}>
            <div className='list-group-item list-group-item-action py-2 lh-tight'>
                <input type='text' placeholder='Find...' onChange={this.handleFindChange} />
            </div>

            {this.renderComponents()}
        </div>
    }

}

export default styled(Sidebar)`
    width: 200px;
    min-width: 200px;
    max-width: 200px;
    border-right: 1px solid;
`;
