import React, { Component, Fragment } from 'react';

export default class AppFooter extends Component {
    render() {
        const currentYear = new Date().getFullYear(); // Add parentheses here to call the method
        return (
            <Fragment>
                <hr />
                {this.props.year}
                <p> Copyright &copy; {currentYear} LLC. </p>
            </Fragment>
        );
    }
}
