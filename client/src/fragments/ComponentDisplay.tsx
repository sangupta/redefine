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
import ReactMarkdown from 'react-markdown';
import CodePlayground from './CodePlayground';
import ComponentPlayground from './ComponentPlayground';
import TabContainer from './Tabs';
import ComponentDetails from './ComponentDetails';
import CopyIcon from './CopyIcon';
import ComponentProps from './ComponentProps';
import ComponentExamples from './ComponentExamples';

interface ComponentDisplayProps {
    component: ComponentDef;
    example?: ComponentExample;
}

const DetailsContainer = styled.div`
    height: 100%;
`;

const ComponentName = styled.h1`
    font-size: 32px;
    line-height: 44px;
    margin-top: 0px;
    margin-bottom: 4px;
`;

const ComponentSourceFile = styled.a`
    margin-bottom: 16px;
    font-size: 14px;
    direction: ltr;
    unicode-bidi: bidi-override;
    overflow: auto;
    color: rgb(118, 118, 118);
    font-style: italic;
`;

export default class ComponentDisplay extends React.Component<ComponentDisplayProps> {

    render() {
        const { component, example } = this.props;

        const tabs = [];
        let exampleTab = 0;

        if (component.description) {
            exampleTab++;
            tabs.push({
                name: 'Details',
                component: <ComponentDetails component={component} />
            })
        }

        if (component.props && component.props.length > 0) {
            exampleTab++;

            tabs.push({
                name: 'Props',
                component: <ComponentProps component={component} />
            });
        }

        if (component.docs && component.docs.trim().length > 0) {
            exampleTab++;

            tabs.push({
                name: 'Examples',
                component: <ComponentExamples component={component} example={example} />
            });
        }

        if (component.playground) {
            tabs.push({
                name: 'Playground',
                component: <ComponentPlayground component={component} />
            });
        }

        return <DetailsContainer>
            <ComponentName>{component.name}</ComponentName>
            <ComponentSourceFile href={component.url || '#'}>
                {component.sourcePath}
                <span style={{ paddingLeft: '12px' }}><CopyIcon /></span>
            </ComponentSourceFile>

            <TabContainer key={component.name + '-' + example?.name} tabs={tabs} selectedTab={example ? exampleTab - 1 : 0} />
        </DetailsContainer>
    }

}
