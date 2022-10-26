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
import { default as Jsx2Json } from 'simplified-jsx-to-json';

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
                    const match = /language-(\w+)/.exec(className || '');

                    // wrap up in fragment if there are multiple children
                    const code = String(children).trim();
                    const json = Jsx2Json(code);
                    const multiple = json && json.length > 1;
                    const sourceCode = multiple ? '<>\n' + code + '\n</>' : code;

                    return (!inline && match)
                        ? <CodePlayground source={sourceCode} />
                        : <code className={className} {...props}>{children}</code>
                }
            }}>{markdown}</ReactMarkdown>
        </ExampleDisplay>
    }

    return null
}
