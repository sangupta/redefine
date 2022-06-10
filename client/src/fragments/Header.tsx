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

interface HeaderProps {
    title: string;
}

export default class Header extends React.Component<HeaderProps> {

    render() {
        return <nav className="navbar navbar-dark bg-dark">
            <div className="container-fluid">
                <a className="navbar-brand" href="#">{this.props.title || 'Redefine'}</a>
            </div>
        </nav>
    }

}
