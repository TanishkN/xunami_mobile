import React, { Component } from 'react';

export default class AppContent extends Component {
    constructor(props) {
        super(props);
        this.state = { posts: [] };
        this.listRef = React.createRef();
        this.handlePostChange = this.handlePostChange.bind(this);
    }

    handlePostChange(posts) {
        this.props.handlePostChange(posts);
    }

    portfolioSelection = () => {
        fetch('https://jsonplaceholder.typicode.com/posts')
            .then((response) => response.json())
            .then(json => {
                this.setState({ posts: json });
                this.handlePostChange(json);
            });
    }

    mouseHover = () => {
        console.log("Text appears");
    }

    clickedItem = (x) => {
        console.log("Clicked", x);
    }

    render() {
        return (
            <div>
                This is the content.
                <br />
                <hr />
                <div onMouseEnter={this.mouseHover}>This is some text</div>

                <button onClick={this.portfolioSelection} className="btn btn-primary">
                    Simulate Portfolio
                </button>

                <p>Posts is {this.state.posts.length} items long</p>

                <hr />
                <ul ref={this.listRef}>
                    {this.state.posts.map((c) => (
                        <li key={c.id}>
                            <a href="#!" onClick={() => this.clickedItem(c.id)}>
                                {c.title}
                            </a>
                        </li>
                    ))}
                </ul>
            </div>
        );
    }
}
