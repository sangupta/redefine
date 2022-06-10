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

import { LiveProvider, LiveEditor, LiveError, LivePreview } from 'react-live';

interface Props {
    source: string;
}

const PreviewRenderer = (props: any) => {
    return <div className='preview-render'>{props.children}</div>
}

export default class CodePlayground extends React.Component<Props> {

    render() {
        const {source} = this.props;
        const ComponentLibrary = (window as any).__ComponentLibrary || [];

        if(!ComponentLibrary || ComponentLibrary.length === 0)  {
            // this is read only mode
            return <code>{source}</code>
        }

        return <LiveProvider code={source} scope={{ ...ComponentLibrary }} >
            <PreviewWrapper>
                <LivePreview Component={PreviewRenderer} />
            </PreviewWrapper>
            <LiveError />
            <CodeWrapper>
                <LiveEditor />
            </CodeWrapper>
        </LiveProvider>
    }

}

const PreviewWrapper = styled.div`
    width: 100%;
    max-width: 100%;
    min-width: 100%;
    padding: 10px;
    border: 1px solid #e8e8e8;
    overflow-x: scroll;
    margin: 8px 0;
`;

const CodeWrapper = styled.div`
    width: 100%;
    max-width: 100%;
    min-width: 100%;
    padding: 10px;
    border: 1px solid #e8e8e8;
    background-color: #f5f5f5;
    margin: 8px 0;
`;
