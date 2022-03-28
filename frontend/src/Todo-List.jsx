import React, { Component } from 'react'
import axios from 'axios';
import {Card, Header, Form, Icon, Input} from 'semantic-ui-react';

let endpoint = 'http://localhost:8080';

class Todolist extends Component {
    constructor(props) {
      super(props)
    
      this.state = {
         task: "",
         items: [],
      }
    }

    render() {
        return (
        <div>
            <div className='row'>
                <Header className='header' as='h2' color='green'>
                    TO DO LIST
                </Header>
            </div>
            
            <div className='row'>
                <Form>
                    <Input
                    type='text'
                    name='task'
                    onChange={this.Change}
                    value={this.state.task}
                    fluid
                    placeholder='Add New Task'
                    >

                    </Input>
                </Form>
            </div>
        </div>
        )
    }

    componentDidMount() {
        this.getAllTasks();
    }
}

export default Todolist;