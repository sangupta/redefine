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

export interface TabValue {
    name: string,
    component: React.ReactNode
}

interface TabContainerProps {
    tabs: Array<TabValue>;
    selectedTab: number;
}

interface TabContainerState {
    selected: number;
}

export default class TabContainer extends React.Component<TabContainerProps, TabContainerState> {

    constructor(props: TabContainerProps) {
        super(props);

        this.state = {
            selected: props.selectedTab || 0
        }
    }

    handleSelect = (index: number) => {
        this.setState({ selected: index });
    }

    render() {
        const { tabs } = this.props;
        if (!tabs || tabs.length === 0) {
            return null;
        }

        const { selected } = this.state;

        // render tab list
        return <>
            <TabNav>
                {tabs.map((item, index) => {
                    return <StyledTab key={item.name || 'tab-index-' + index} index={index} selected={selected} title={item.name} onSelect={this.handleSelect} />
                })}
            </TabNav>
            {(tabs[selected] || {}).component}
        </>
    }

}

const TabNav = styled.ul`
    display: flex;
    flex-wrap: nowrap;
    border-bottom: 1px solid #dee2e6;
`;

interface TabProps {
    index: number;
    selected: number;
    title: string;
    onSelect: (index: number) => void;
    className?: string;
}

class Tab extends React.Component<TabProps> {

    handleSelect = (e: React.MouseEvent) => {
        e.preventDefault();
        e.stopPropagation();
        if (this.props.onSelect) {
            this.props.onSelect(this.props.index);
        }
    }

    render() {
        const { className, title } = this.props;

        return <li className={className}>
            <a aria-current="page" href="#" onClick={this.handleSelect}>{title}</a>
        </li>
    }
}

const StyledTab = styled(Tab)`
    list-style: none;

    & a {
        border-top-left-radius: 4px;
        border-top-right-radius: 4px;
        margin-bottom: -1px;
        display: block;
        border: 1px solid #0000;

        color: ${props => props.index !== props.selected ? '#0d6efd' : '#495057'};
        text-decoration: none;
        transition: color .15s ease-in-out, background-color .15s ease-in-out, border-color .15s ease-in-out;
        padding: 8px 16px;
        border-color: ${props => props.index !== props.selected ? '#0000' : '#dee2e6 #dee2e6 #fff'};
    }
`;
