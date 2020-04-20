import React, { Component } from 'react';
import './App.css';

class App extends Component {
  
  async handleClick() {

    try {
      let response = await fetch('http://localhost:8080/classify', {
        method: 'post',
        mode: 'no-cors',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          url: "https://geertbaeke.files.wordpress.com/2019/01/cassette.jpg"
        })
      });
      
      console.log('Result: '+ response);

    } catch (error) {
      console.log(error)
    }
  }
  
  render() {
    return (
      <div className="App">
          <button onClick={this.handleClick}>
            Predict
            </button>
      </div>
  );
  }
}

export default App;
