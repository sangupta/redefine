import React from 'react';
import { ComponentDef } from '../Utils';

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
