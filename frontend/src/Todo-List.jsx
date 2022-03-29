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
            <div className="row">
                <Header className="header" as="h2" color="green">
                    TO DO LIST
                </Header>
            </div>
            
            <div className="row">
                <Form onSubmit={this.submit}>
                    <Input
                    type="text"
                    name="task"
                    onChange={this.change}
                    value={this.state.task}
                    fluid
                    placeholder="Add New Task"
                    >
                    </Input>
                    {/* <Button>Create Task</Button> */}
                </Form>
            </div>

            <div className="row">
                <Card.Group>{this.state.items}</Card.Group>
            </div>
        </div>
        ) 
    }  

    componentDidMount() {
        this.getAllTasks();
    }

    change = (e) => {
        this.setState({
            [e.target.name]: e.target.value,
        }) 
    }

    getAllTasks = () => {
        axios.get(endpoint+"/").then(res => {
            if (res.data) {
                this.setState({
                    items: res.data.map(item => {
                        let color = 'green';
                        let style = {
                            wordWrap: "break-word",
                        };

                        if (item.status) {
                            color = 'yellow'
                            style["textDecorationLine"] = 'line-through';
                        }

                        return (
                            <Card key={item._id} color={color} fluid className='rough'>
                                <Card.Content>
                                    <Card.Header textAlign='left'>
                                        <div style={style}>{item.task}</div>
                                    </Card.Header>

                                    <Card.Meta textAlign='right'>  
                                        <Icon 
                                        name='check circle'
                                        color='blue'
                                        onClick={() => this.updateTask(item._id)}
                                        />
                                        <span style={{paddingRight: 10}}>Done</span>

                                        <Icon 
                                        name='undo'
                                        color='yellow'
                                        onClick={() => this.undoTask(item._id)}
                                        />
                                        <span style={{paddingRight: 10}}>Undo</span>

                                        <Icon 
                                        name='delete'
                                        color='red'
                                        onClick={() => this.deleteTask(item._id)}
                                        />
                                        <span style={{paddingRight: 10}}>Delete</span>
                                    </Card.Meta>
                                </Card.Content>
                            </Card>
                        )
                    })
                })
            } else {
                this.setState({
                    items: [],
                });
            };
        })
    }

    updateTask = (id) => {
        axios.put(endpoint+"/task/done/"+id, {
            headers: {
                "Content-Type:": "application/x-www-form-urlencoded",
            }, 
        }).then(res => {
            console.log(res)
            this.getAllTasks();
        })
    }

    undoTask = (id) => {
        axios.put(endpoint+"/task/undo/"+id, {
            headers: {
                "Content-Type:": "application/x-www-form-urlencoded",
            }, 
        }).then(res => {
            console.log(res)
            this.getAllTasks();
        })
    }

    deleteTask = (id) => {
        axios.delete(endpoint+"/task/delete/"+id, {
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            }
        }).then(res => {
            console.log(res)
            this.getAllTasks();
        })
    }

    submit = () => {
        let {task} = this.state;

        if (task) {
            axios.post(endpoint+"/task",
                {task,},
                {
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded",
                    },
                }
            ).then(res => {
                 this.getAllTasks()

                this.setState({
                    task: "",
                })

                console.log(res);
            })
        }
    }
}
 
export default Todolist; 