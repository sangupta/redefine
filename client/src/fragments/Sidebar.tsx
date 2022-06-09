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
            result.push(<ComponentItem key={component.name} component={component} onSelect={this.handleComponentSelect} />)
        }
        return <div className='list-group list-group-flush border-bottom scrollarea'>
            {result}
        </div>
    }

    render() {
        const { components } = this.props;
        if(!components || components.length === 0) {
            return null;
        }

        return <div className='d-flex flex-column align-items-stretch flex-shrink-0 bg-white sidebar'>
            {this.renderComponents()}
        </div>
    }

}
