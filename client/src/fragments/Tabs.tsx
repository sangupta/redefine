import React from 'react';

export interface TabValue {
    name: string,
    component: React.ReactNode
}

interface TabContainerProps {
    tabs: Array<TabValue>
}

interface TabContainerState {
    selected: number;
}

export default class TabContainer extends React.Component<TabContainerProps, TabContainerState> {

    constructor(props: TabContainerProps) {
        super(props);

        this.state = {
            selected: 0
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
            <ul className="nav nav-tabs">
                {tabs.map((item, index) => {
                    return <Tab index={index} selected={selected} title={item.name} onSelect={this.handleSelect} />
                })}
            </ul>
            {tabs[selected].component}
        </>
    }

}

interface TabProps {
    index: number;
    selected: number;
    title: string;
    onSelect: (index: number) => void;
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
        const { index, selected, title } = this.props;

        return <li className="nav-item">
            <a className={'nav-link ' + (index === selected ? 'active' : '')} aria-current="page" href="#" onClick={this.handleSelect}>{title}</a>
        </li>
    }
}
