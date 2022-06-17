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

import styled from 'styled-components';
import ReactMarkdown from 'react-markdown';
import CodePlayground from './CodePlayground';

const ExampleDisplay = styled.div`
    margin: 20px 0;
`;

/**
 * Render the associated markdown file. This includes rendering
 * all the editable code examples here.
 */
export default function ComponentExamples({ component, example }: { component: ComponentDef, example?: ComponentExample }) {
    const markdown = example?.markdown || component.docs || '';

    if (markdown) {

        return <ExampleDisplay>
            <h3>{example?.name}</h3>
            <ReactMarkdown className='markdown-docs' components={{
                code({ node, inline, className, children, ...props }) {
                    const match = /language-(\w+)/.exec(className || '')
                    const sourceCode = '<>\n' + String(children).trim() + '\n</>'

                    return (!inline && match)
                        ? <CodePlayground source={sourceCode} />
                        : <code className={className} {...props}>{children}</code>
                }
            }}>{markdown}</ReactMarkdown>
        </ExampleDisplay>
    }

    return null
}
