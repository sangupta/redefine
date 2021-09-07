import React from 'react';
import { ComponentDef } from '../Utils';

interface ComponentDetailsProps {
    component: ComponentDef;
}

export default class ComponentDetails extends React.Component<ComponentDetailsProps> {

    renderProps() {
        const { component } = this.props;
        if (!component.props || component.props.length === 0) {
            return null;
        }

        const rows = [];
        for (let index = 0; index < component.props.length; index++) {
            const prop = component.props[index];

            rows.push(<tr>
                <td>{prop.name}</td>
                <td>{prop.type}</td>
                <td>{'' + prop.required}</td>
                <td>{prop.defaultValue || ''}</td>
                <td>{prop.description || ''}</td>
            </tr>);
        }

        return <table className='table table-striped table-bordered'>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Type</th>
                    <th>Required</th>
                    <th>Default Value</th>
                    <th>Description</th>
                </tr>
            </thead>
            <tbody>
                {rows}
            </tbody>
        </table>
    }

    render() {
        const { component } = this.props;

        return <div className='component-details'>
            <h1>{component.name}</h1>
            <h6>{component.sourcePath}</h6>

            <p>{component.description}</p>

            {this.renderProps()}
        </div>
    }

}
