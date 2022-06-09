import React from 'react';

import { ComponentDef, NoProps } from './Utils';

import Header from './fragments/Header';
import Sidebar from './fragments/Sidebar';
import ContentPane from './fragments/ContentPane';

interface AppState {
    components: Array<ComponentDef>;
    selectedComponent?: ComponentDef;
    title: string;
}

export default class App extends React.Component<NoProps, AppState> {

    constructor(props: NoProps) {
        super(props);

        this.state = {
            components: [],
            title: ''
        }
    }

    componentDidMount = async () => {
        const response = await fetch('http://localhost:13090/components.json')
        const data = await response.json();
        this.setState({ title: data.title, components: data.components });
    }

    handleComponentSelect = (def: ComponentDef): void => {
        this.setState({ selectedComponent: def });
    }

    render() {
        return <>
            <Header title={this.state.title} />
            <div className='d-flex flex-row flex-1'>
                <Sidebar components={this.state.components} onComponentSelect={this.handleComponentSelect} />
                <ContentPane component={this.state.selectedComponent} />
            </div>
            <footer className="footer mt-auto bg-dark">
                <div className='container-fluid'>
                    <span className="text-muted">powered by Redefine</span>
                </div>
            </footer>
        </>
    }

}
