import React, { Component } from 'react';
import axios from 'axios';

import {
  Container, Col, Form,
  FormGroup, Label, Input,
  Button, FormText, FormFeedback,
} from 'reactstrap';


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
      validate :{
        urlState: '',
      },
    }
    this.handleChange = this.handleChange.bind(this);
  }

  handleChange = event => {
    const imageUrlRex = /^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)jpg$/;
    const { validate } = this.state
    if (imageUrlRex.test(event.target.value)) {
      validate.urlState = 'has-success'
    } else {
      validate.urlState = 'has-danger'
    }

    this.setState({validate})
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
      <Container fluid >
        <h2>Image Classifier</h2>
      <Form className="form" onSubmit={this.handleSubmit}>
        <Col>
        <FormGroup>
          <Label>ImageUrl</Label>
            <Input
              name="url"
              type="url"
              id="exampleUrl"
              placeholder="https://image.jpg"
              value={url}
              valid={ this.state.validate.urlState === 'has-success' }
              invalid={ this.state.validate.urlState === 'has-danger' }
              onChange={this.handleChange}/>
            <FormText className="text-muted">Image URL should of type jpeg.</FormText>
            <FormFeedback valid>
              This looks like a valid URL
            </FormFeedback>
            <FormFeedback invalid>
              This is an invalid URL
            </FormFeedback>
        </FormGroup>
        </Col>
        <Button color="success" type="submit">Submit</Button>
      </Form>
      </Container>
  );
  }
}

export default PostPrediction