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
import ComponentDetails from './ComponentDetails';

interface ContentPaneProps {
    className?: string;
    component?: ComponentDef;
    example?: ComponentExample;
}

class ContentPane extends React.Component<ContentPaneProps> {

    renderDetails() {
        const { component, example } = this.props;
        if (!component) {
            return "Content Pane"
        }

        return <ComponentDetails key={component.sourcePath} component={component} example={example} />
    }

    render() {
        return <main className={'content-pane w-100 ' + this.props.className}>
            {this.renderDetails()}
        </main>
    }

}

export default styled(ContentPane)`
    padding: 20px;
    overflow-y: scroll;
`;
