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

import { processComponentInfo } from './Utils';

import Header from './fragments/Header';
import Sidebar from './fragments/Sidebar';
import ContentPane from './fragments/ContentPane';

/**
 * State attributes for the app component.
 */
interface AppState {
    components: Array<ComponentDef>;
    selectedComponent?: ComponentDef;
    selectedExample?: ComponentExample;
    title: string;
    error: boolean;
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
            title: '',
            error: false
        }
    }

    /**
     * Load components.json file once the app is mounted.
     * 
     */
    componentDidMount = async () => {
        try {
            const response = await fetch('http://localhost:1309/components.json')
            const data = await response.json();
            this.setState({ title: data.title, components: processComponentInfo(data.components) });
        } catch (e) {
            this.setState({ error: true });
        }
    }

    /**
     * Handle selection of a particular component.
     * 
     * @param def 
     */
    handleComponentSelect = (def: ComponentDef, example?: ComponentExample): void => {
        this.setState({ selectedComponent: def, selectedExample: example });
    }

    /**
     * Render the component.
     * 
     * @returns 
     */
    render(): React.ReactNode {
        const { error } = this.state;
        if (error) {
            return <>
                <Header title={this.state.title} />
                <BodyContainer>
                    <Alert>Unable to load component definition file.</Alert>
                </BodyContainer>
            </>
        }
        return <>
            <Header title={this.state.title} />
            <BodyContainer>
                <Sidebar components={this.state.components} onComponentSelect={this.handleComponentSelect} />
                <ContentPane component={this.state.selectedComponent} example={this.state.selectedExample} />
            </BodyContainer>
        </>
    }

}

const BodyContainer = styled.div`
    display: flex;
    flex-direction: row;
    flex: 1;
    overflow: hidden;
`;

/**
 * We are all set, mount the application component.
 */
ReactDOM.render(<App />, document.body);

const Alert = styled.div`
    color: #842029;
    background-color: #f8d7da;
    border-color: #f5c2c7;
    border: 1px solid;
    border-radius: 6px;
    height: fit-content;
    padding: 12px;
    margin: 0 auto;
    margin-top: 120px;
`;
