import React from 'react';

import { ComponentDef } from './Utils';

import Header from './fragments/Header';
import Sidebar from './fragments/Sidebar';
import ContentPane from './fragments/ContentPane';

const SERVER_URL = 'http://localhost:13090';

interface AppState {
    components: Array<ComponentDef>;
    selectedComponent: ComponentDef | undefined;
}

export default class App extends React.Component<{}, AppState> {

    constructor(props) {
        super(props);

        this.state = {
            components: [],
            selectedComponent: undefined
        }
    }

    componentDidMount = async () => {
        const response = await fetch(SERVER_URL + '/components.json')
        const data = await response.json();
        this.setState({ components: data });
    }

    handleComponentSelect = (def: ComponentDef): void => {
        this.setState({ selectedComponent: def });
    }

    render() {
        return <div className='d-flex flex-column'>
            <Header />
            <div className='d-flex flex-row'>
                <Sidebar components={this.state.components} onComponentSelect={this.handleComponentSelect} />
                <ContentPane component={this.state.selectedComponent} />
            </div>
        </div>
    }

}
