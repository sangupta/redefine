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
import styled from 'styled-components';
import ComponentDisplay from './ComponentDisplay';

interface ContentPaneProps {
    className?: string;
    component?: ComponentDef;
    example?: ComponentExample;
}

export default class ContentPane extends React.Component<ContentPaneProps> {

    renderDetails() {
        const { component, example } = this.props;
        if (!component) {
            return "Content Pane"
        }

        return <ComponentDisplay key={component.sourcePath} component={component} example={example} />
    }

    render() {
        return <Main>
            {this.renderDetails()}
        </Main>
    }

}

const Main = styled.main`
    padding: 20px;
    overflow-y: scroll;
    width: 100%;
    border-left: 1px solid var(--redefine-border-color)
`;
