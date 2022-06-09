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
