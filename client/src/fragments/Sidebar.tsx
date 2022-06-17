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

import { componentSorter } from './../Utils';
import ComponentItem from './ComponentItem';

interface SidebarProps {
    className?: string;
    components: Array<ComponentDef>;
    onComponentSelect: (def: ComponentDef, example?: ComponentExample) => void;
}

interface SidebarState {
    filter?: string;
}

export default class Sidebar extends React.Component<SidebarProps, SidebarState> {

    constructor(props: SidebarProps) {
        super(props);

        this.state = {
            filter: ''
        }
    }

    handleComponentSelect = (def: ComponentDef, example?: ComponentExample) => {
        this.props.onComponentSelect(def, example);
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
        return <ComponentContainer>
            {result}
        </ComponentContainer>
    }

    handleFilterChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({ filter: e.target.value });
    }

    render() {
        const { components } = this.props;

        if (!components || components.length === 0) {
            return null;
        }

        return <SidebarContainer>
            <SearchContainer>
                <input type='text' placeholder='Find...' onChange={this.handleFilterChange} />
            </SearchContainer>

            {this.renderComponents()}
        </SidebarContainer>
    }

}

const SidebarContainer = styled.div`
    display: flex;
    flex-direction: column;
    align-items: stretch !important;
    flex-shrink: 0 !important;
    background-color: var(--redefine-bg);

    width: var(--redefine-sidebar-width);
    min-width: var(--redefine-sidebar-width);
    max-width: var(--redefine-sidebar-width);
`;

const SearchContainer = styled.div`
    color: #212529;
    padding: .5rem;
    text-decoration: none;
`;

const ComponentContainer = styled.div`
    flex-direction: column;
    margin-bottom: 0;
    padding-left: 0;
    display: flex;
    overflow-y: auto;
    line-height: 22px;
`;
