import React, { Component } from 'react';
import './App.css';
import axios from 'axios';

const api = axios.create({
  baseURL: `http://localhost:8080`,
  headers: {
    'Content-Type': 'application/json',
  }
})


class App extends Component {

  state = {
    url: '',
  };

  handleChange = event => {
    this.setState({ name: event.target.value });
  }

  handleSubmit = event => {
    event.preventDefault(); //stop browser from reloading page
    console.log('The link was clicked.');
  
    const url = {
      url: this.state.url,
    }

    api
    .post('/classify', { url })
    .then(res => {
        console.log(res);
        console.log(res.data)
      }
    )
  }

  render() {
    return (
      <div className="App">
        <form>
          <label>Image URL:
            <input type="url" name="Image URL" onChange={this.handleChange} />
            </label>
        </form>
          <button type="submit" onSubmit={this.handleSubmit}>Predict</button>
      </div>
  );
  }
}

export default App;
