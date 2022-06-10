/**
 *
 * Redefine - UI component documentation
 * 
 * MIT License.
 * Copyright (c) 2022, Sandeep Gupta.
 * 
 * Use of this source code is governed by a MIT style license
 * that can be found in LICENSE file in the code repository.
 * 
 **/

import React from 'react';
import ReactDOM from 'react-dom';
import styled from 'styled-components';

import { ComponentDef, NoProps } from './Utils';

import Header from './fragments/Header';
import Sidebar from './fragments/Sidebar';
import ContentPane from './fragments/ContentPane';

/**
 * State attributes for the app component.
 */
interface AppState {
    components: Array<ComponentDef>;
    selectedComponent?: ComponentDef;
    title: string;
}

/**
 * Styled `footer` HTML element.
 */
const Footer = styled.footer`
    font-size: 11px;
    color: white;
    line-height: 36px;
`;

/**
 * The main Redefine client application.
 * 
 * @author sangupta
 */
class App extends React.Component<NoProps, AppState> {

    /**
     * Constructor.
     * 
     * @param props 
     */
    constructor(props: NoProps) {
        super(props);

        this.state = {
            components: [],
            title: ''
        }
    }

    /**
     * Load components.json file once the app is mounted.
     * 
     */
    componentDidMount = async () => {
        const response = await fetch('http://localhost:1309/components.json')
        const data = await response.json();
        this.setState({ title: data.title, components: data.components });
    }

    /**
     * Handle selection of a particular component.
     * 
     * @param def 
     */
    handleComponentSelect = (def: ComponentDef): void => {
        this.setState({ selectedComponent: def });
    }

    /**
     * Render the component.
     * 
     * @returns 
     */
    render(): React.ReactNode {
        return <>
            <Header title={this.state.title} />
            <div className='d-flex flex-row flex-1'>
                <Sidebar components={this.state.components} onComponentSelect={this.handleComponentSelect} />
                <ContentPane component={this.state.selectedComponent} />
            </div>
            <Footer className="footer mt-auto bg-dark">
                <div className='container-fluid'>
                    <span className="text-muted">powered by Redefine</span>
                </div>
            </Footer>
        </>
    }

}

/**
 * We are all set, mount the application component.
 */
ReactDOM.render(<App />, document.body);
