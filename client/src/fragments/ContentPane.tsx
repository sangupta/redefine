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
import StyledMarkdown from './StyledMarkdown';

interface ContentPaneProps {
    className?: string;
    component?: ComponentDef;
    example?: ComponentExample;
    meta: RedefinePayload;
}

export default class ContentPane extends React.Component<ContentPaneProps> {

    renderDetails() {
        const { component, example, meta } = this.props;
        if (!component) {
            // check what to display
            if (meta.libDocs) {
                return <StyledMarkdown>{meta.libDocs}</StyledMarkdown>
            }

            const keys = Object.keys(meta);
            if(keys.length === 0) {
                return null;
            }

            let kids: string = '';
            kids += '# ' + (meta.title + '');
            kids += '\n';
            if(meta.version) {
                kids += 'Version: ' + meta.version;
                kids += '\n\n';
            }
            
            if(meta.license) {
                kids += 'License: ' + meta.license;
                kids += '\n\n';
            }

            if(meta.homePage) {
                kids += 'Home: ' + `[${meta.homePage}](${meta.homePage})`
                kids += '\n\n';
            }

            kids += (meta.description || '');
            kids += '\n\n';

            if (meta.author && meta.author.name) {
                let author = meta.author.name;
                if(meta.author.url) {
                    author = `[${meta.author.name}](${meta.author.url})`
                }

                kids += 'Author: ' + author + (meta.author.email ? ` at [${meta.author.email}](mailto:${meta.author.email})` : '');
            }

            return <StyledMarkdown>{kids}</StyledMarkdown>
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
