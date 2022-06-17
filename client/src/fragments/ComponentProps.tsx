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

import { propsSorter } from '../Utils';
import styled from 'styled-components';
import StyledMarkdown from './StyledMarkdown';

const StyledCode = styled.code`
    color: #d63384;
    word-wrap: break-word;
    font-size: 14px;
`;

const StyledPre = styled.pre`
    word-wrap: break-word;
    font-size: 14px;
    vertical-align: top;
    margin: 2px 0;
`;

const SectionName = styled.h5`
    font-size: 22px;
    line-height: 33px;
    margin: 16px 0;
`;

function renderType(prop: PropDef) {
    if (!prop.type) {
        return prop.type;
    }

    switch (prop.type) {
        case '$function':
            if (prop.params?.length || prop.returnType) {
                let paramStr = '';
                if (prop.params && prop.params.length > 0) {
                    for (let index = 0; index < prop.params.length; index++) {
                        const param = prop.params[index];
                        if (index > 0) {
                            paramStr += ', ';
                        }
                        paramStr += param.name;
                        if (param.type && param.type !== '$unknown') {
                            paramStr += ':' + param.type;
                        }
                    }
                }
                return '(' + paramStr + ') => ' + prop.returnType;
            }

            return prop.type;

        case '$enum':
            if (prop.enumOf && prop.enumOf.length > 0) {
                let combined = '';
                for (let index = 0; index < prop.enumOf.length; index++) {
                    const paramDef = prop.enumOf[index];

                    if (index > 0) {
                        combined += ' | ';
                    }

                    if (paramDef.type === 'string') {
                        combined += '"' + paramDef.name + '"'
                    } else {
                        combined += paramDef.name
                    }
                }

                return combined;
            }

            return '$enum';

        default:
            return prop.type;
    }
}

export default function ComponentProps({ component }: { component: ComponentDef }) {
    if (!component.props || component.props.length === 0) {
        return null;
    }

    component.props = component.props.sort(propsSorter)

    const rows = [];
    for (let index = 0; index < component.props.length; index++) {
        const prop = component.props[index];

        rows.push(<tr key={prop.name}>
            <td><StyledCode>{prop.name}</StyledCode></td>
            <td><StyledCode>{renderType(prop)}</StyledCode></td>
            <td><StyledPre>{'' + prop.required}</StyledPre></td>
            <td><StyledPre>{prop.defaultValue || ''}</StyledPre></td>
            <td><StyledMarkdown>{prop.description || ''}</StyledMarkdown></td>
        </tr>);
    }

    return <>
        <SectionName>Props</SectionName>
        <Table>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Type</th>
                    <th>Required</th>
                    <th>Default Value</th>
                    <th>Description</th>
                </tr>
            </thead>
            <tbody>
                {rows}
            </tbody>
        </Table>
    </>
}

const Table = styled.table`
    width: 100%;
    color: #212529;
    border-color: #dee2e6;
    margin-bottom: 16px;
    caption-side: bottom;
    border-collapse: collapse;

    & > thead {
        vertical-align: bottom;
    }

    & tr {
        border-color: inherit;
        border-style: solid;
    }

    & td {
        border-color: inherit;
        border-style: solid;
        vertical-align: top;
    }

    & th {
        border-width: 0 1px;
        border-color: inherit;
        border-style: solid;
    }

    & > :not(caption) > * {
        border-width: 1px 0;
    }

    & > :not(caption) > * > * {
        border-width: 0 1px;
        padding: 8px;
        box-shadow: inset 0 0 0 9999px transparent;
    }

    & > :not(:first-child) {
        border-top: 2px solid;
    }

    & tbody {
        vertical-align: inherit;
    }

    & > tbody > tr:nth-of-type(2n+1) > * {
        color: rgb(33, 37, 41);
    }
`;
