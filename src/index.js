import React, { Component } from 'react';
import ReactDOM from 'react-dom'; 
import AppContent from '../frontend/AppContent';
import AppFooter from '../frontend/AppFooter';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = { posts: [] };
    this.handlePostChange = this.handlePostChange.bind(this);
  }

  handlePostChange(posts) {
    this.setState({ posts: posts });
  }

  render() {
    const myProps = {
      title: "My App!",
      subject: "My subject",
      year: "3024"
    };

    return (
      <div>
        <h1>Hello Man!</h1>
        <AppContent handlePostChange={this.handlePostChange} />
        <AppFooter {...myProps} />
      </div>
    );
  }
}

ReactDOM.render(<App />, document.getElementById('root'));
