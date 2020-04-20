import React, { Component } from 'react';
import axios from 'axios';

const api = axios.create({
  baseURL: `http://localhost:8080`,
  headers: {
    'Content-Type': 'application/json',
  }
})


class PostPrediction extends Component {

  constructor(props) {
    super(props);

    this.state = {
      url: '',
    }
  }

  handleChange = event => {
    this.setState({ [event.target.name]: event.target.value });
  }

  handleSubmit = event => {
    event.preventDefault(); //stop browser from reloading page

    const URL = {
      url: this.state.url
    }

    api
    .post('/classify', URL)
    .then(res => {
        console.log(res.data)
      }
    )
  }

  render() {
    const { url } = this.state;
    return (
      <div className="App">
        <form onSubmit={this.handleSubmit}>
          <div>
            <input type="url" name="url" value={url} onChange={this.handleChange} />
          </div>
          <button type="submit">Predict</button>
        </form>
      </div>
  );
  }
}

export default PostPrediction