import React from 'react';
import './App.css'; 
import {Container} from 'semantic-ui-react';
import Todolist from './Todo-List'

function App() {
  return (
    <div>
      <Container>
        <Todolist />
      </Container>
    </div>
  )
}

export default App;
