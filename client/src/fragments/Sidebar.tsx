import React from 'react';
import styled from 'styled-components';

import { ComponentDef, componentSorter } from './../Utils';
import ComponentItem from './ComponentItem';

interface SidebarProps {
    className?: string;
    components: Array<ComponentDef>;
    onComponentSelect: (def: ComponentDef) => void;
}

class Sidebar extends React.Component<SidebarProps> {

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
        if (!components || components.length === 0) {
            return null;
        }

        return <div className={'d-flex flex-column align-items-stretch flex-shrink-0 bg-white ' + this.props.className}>
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