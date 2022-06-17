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

interface HeaderProps {
    title: string;
}

const Nav = styled.nav`
    display: flex;
    width: 100%;
    justify-content: space-between;
    align-items: center;
    background-color: var(--redefine-bg);
    height: 40px;
    border-bottom: 1px solid var(--redefine-border-color);
`;

const Container = styled.div`
    width: 100%;
    padding-left: 12px;
    padding-right: 12px;
`;

const BrandLink = styled.a`
    color: var(--redefine-alt-color);
    white-space: nowrap;
    text-decoration: none;
    font-size: 20px;
    padding-top: 4px;
    padding-bottom: 4px;
    margin-right: 16px;    
`;

export default class Header extends React.Component<HeaderProps> {

    render() {
        return <Nav>
            <Container>
                <BrandLink href="#">{this.props.title || ''}</BrandLink>
                <Redefine href='https://redefine.sangupta.com' target='_blank'>redefined</Redefine>
            </Container>
        </Nav>
    }

}

const Redefine = styled.a`
    color: #aaa;
    font-style: italic;
    border-left: 1px solid #aaa;
    padding-left: 20px;
    text-decoration: none;

    &:hover {
        text-decoration: none;
    }
`;
