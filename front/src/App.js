import React from 'react';
import {
  BrowserRouter,
  Switch,
  Route
} from "react-router-dom";
import Container from 'react-bootstrap/container';

import Header from "./components/header/header";
import Home from "./components/home/home";
import About from "./components/about/about";
import Footer from "./components/footer/footer";
import Login from "./components/login/login";
import Logout from "./components/logout/logout";
import Client from "./components/client/client";
import Activity from "./components/activity/activity";
import Document from "./components/document/document";
import Message from "./components/message/message";

import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';

function App() {
  return (
    <Container>
    <div className="App">
      <BrowserRouter>
        <Header></Header>
        <Switch>
          <Route path="/" exact>
            <Home></Home>
          </Route>
          <Route path="/home">
            <Home></Home>
          </Route>
          <Route path="/client">
            <Client></Client>
          </Route>
          <Route path="/activity">
            <Activity></Activity>
          </Route>
          <Route path="/document">
            <Document></Document>
          </Route>
          <Route path="/message">
            <Message></Message>
          </Route>
          <Route path="/about">
            <About></About>
          </Route>
          <Route path="/login">
            <Login />
          </Route>
          <Route path="/logout">
            <Logout />
          </Route>
        </Switch>
        <Footer></Footer>
      </BrowserRouter>
    </div>
    </Container>
  );
}

export default App;
