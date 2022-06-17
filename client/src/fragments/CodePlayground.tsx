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

export default class CodePlayground extends React.PureComponent<Props> {

    render() {
        const { source } = this.props;
        const ComponentLibrary = (window as any).__ComponentLibrary || [];

        if (!ComponentLibrary || ComponentLibrary.length === 0) {
            // this is read only mode
            return <code>{source}</code>
        }

        return <LiveProvider code={source} scope={{ ...ComponentLibrary }} >
            <PreviewWrapper>
                <LivePreview Component={PreviewRenderer} />
            </PreviewWrapper>
            <LiveError />
            <CodeWrapper>
                <LiveEditor style={{ fontSize: '16px', lineHeight: '22px' }} />
            </CodeWrapper>
        </LiveProvider>
    }

}

const PreviewWrapper = styled.div`
    padding: 10px;
    border: 1px solid var(--redefine-border-color);
    overflow-x: scroll;
    margin: 8px 0;
    border-radius: 2px;
`;

const CodeWrapper = styled.div`
    padding: 10px;
    border: 1px solid var(--redefine-border-color);
    background-color: var(--redefine-bg);
    margin: 8px 0;
    border-radius: 2px;
`;
