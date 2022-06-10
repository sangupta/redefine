import React from 'react';
import { ComponentDef, propsSorter } from '../Utils';

interface ComponentPlaygroundProps {
    component: ComponentDef;
}

export default class ComponentPlayground extends React.Component<ComponentPlaygroundProps> {

    render() {
        const { component } = this.props;
        if (!component.props || component.props.length === 0) {
            return null;
        }

        // sort
        component.props = component.props.sort(propsSorter)

        // render the element for each type of form element
        for (let index = 0; index < component.props.length; index++) {
            const prop = component.props[index];
        }
        
        return null;
    }

}
