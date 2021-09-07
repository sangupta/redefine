import React from 'react';
import { ComponentDef } from '../Utils';
import ComponentDetails from './ComponentDetails';

interface ContentPaneProps {
    component: ComponentDef;
}

export default class ContentPane extends React.Component<ContentPaneProps> {

    renderDetails() {
        const { component } = this.props;
        if (!component) {
            return "Content Pane"
        }

        return <ComponentDetails component={component} />
    }

    render() {
        return <main className='content-pane'>
            {this.renderDetails()}
        </main>
    }

}
