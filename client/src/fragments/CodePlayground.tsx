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
import styled, { StyledComponent } from 'styled-components';

import { LiveProvider, LiveEditor, LiveError, LivePreview } from 'react-live';

interface Props {
    source: string;
    editable?: boolean;
}

interface CodePlaygroundState {
    showReact: boolean;
}

const PreviewRenderer = (props: any) => {
    return <div id={props.id} className='preview-render'>{props.children}</div>
}

export default class CodePlayground extends React.PureComponent<Props, CodePlaygroundState> {

    static defaultProps = {
        editable: true
    }

    uniqueID: string;

    constructor(props: Props) {
        super(props);

        // generate a unique ID
        this.uniqueID = 'unique-' + crypto.randomUUID();

        this.state = {
            showReact: true
        }
    }

    showCode = () => {
        const ele = document.getElementById(this.uniqueID);
        if (ele) {
            console.log(ele.innerHTML);
        }
    }

    selectReact = () => {
        this.setState({ showReact: true });
    }

    selectHtml = () => {
        this.setState({ showReact: false });
    }

    getGeneratedHtmlCode = () => {
        const ele = document.getElementById(this.uniqueID);
        if (ele) {
            return ele.innerHTML;
        }

        return null;
    }

    render() {
        const { source, editable } = this.props;
        const { showReact } = this.state;

        const ComponentLibrary = (window as any).__ComponentLibrary || [];

        if (!ComponentLibrary || ComponentLibrary.length === 0) {
            // this is read only mode
            return <code>{source}</code>
        }

        // used to fix vscode warning
        const PreviewComponent: any = LivePreview;

        return <LiveProvider code={source} scope={{ ...ComponentLibrary }} disabled={!editable}>
            <PreviewWrapper>
                <PreviewComponent id={this.uniqueID} Component={PreviewRenderer} />
            </PreviewWrapper>
            <LiveError />
            <CodeTabs>
                <CodeTab active={showReact} onClick={this.selectReact}>React</CodeTab>
                <CodeTab active={!showReact} onClick={this.selectHtml}>HTML</CodeTab>
            </CodeTabs>
            <CodeWrapper>
                {showReact ? <LiveEditor style={{ fontSize: '16px', lineHeight: '22px' }} /> : ''}
                {!showReact ? <pre>{this.getGeneratedHtmlCode()}</pre> : ''}
            </CodeWrapper>
        </LiveProvider>
    }

}

const PreviewWrapper = styled.div`
    --dot: 0 0 0;
    padding: 20px;
    border: 1px solid var(--redefine-border-color);
    overflow-x: scroll;
    margin: 8px 0;
    border-radius: 8px;
    background-image: radial-gradient(rgb(var(--dot)/.2) .75px,transparent .75px);
    background-size: 8px 8px;
`;

const CodeTabs = styled.div`
    display: flex;
    flex-direction: row;
    background-color: var(--redefine-bg);
    border: 1px solid var(--redefine-border-color);
    border-bottom: none;
    border-radius: 8px 8px 0 0;
    height: 36px;
    align-items: center;
`;

const CodeTab: StyledComponent<"span", any, any, never> = styled.span`
    padding: 0 8px;
    margin: 0 8px;
    cursor: pointer;

    border-bottom: ${props => props.active ? '2px solid blue' : 'none'};
`;

const CodeWrapper = styled.div`
    padding: 10px;
    border: 1px solid var(--redefine-border-color);
    // background-color: var(--redefine-bg);
    margin: 8px 0;
    border-radius: 0 0 8px 8px;
    margin-top: 0;
`;
