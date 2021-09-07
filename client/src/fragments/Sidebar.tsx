import React from 'react';
import { ComponentDef, componentSorter } from './../Utils';
import ComponentItem from './ComponentItem';

interface SidebarProps {
    components: Array<ComponentDef>;

    onComponentSelect: (def: ComponentDef) => void;
}

export default class Sidebar extends React.Component<SidebarProps> {

    handleComponentSelect = (def: ComponentDef) => {
        this.props.onComponentSelect(def);
    }

    renderComponents() {
        const result = [];
        const { components } = this.props;

        components.sort(componentSorter);

        for (let index = 0; index < components.length; index++) {
            const component = components[index];
            result.push(<ComponentItem component={component} onSelect={this.handleComponentSelect} />)
        }
        return <ul className='component-list'>{result}</ul>;
    }

    render() {
        return <div className='sidebar'>
            {this.renderComponents()}
        </div>
    }

}
