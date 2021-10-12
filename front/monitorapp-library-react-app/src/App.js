import React from "react";
import "./App.css";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import Home from "./pages";
import SigninPage from "./pages/signin";
import BookPage from "./pages/book";
import RentReturnPage from "./pages/rent-return";
import UserPage from "./pages/user";
import AboutPage from "./pages/about";
import SignupPage from "./pages/signup";


function App() {
    return (
        <Router>
            <Switch>
                <Route path="/" component={Home} exact />
                <Route path="/about" component={AboutPage} exact />
                <Route path="/book" component={BookPage} exact />
                <Route path="/rent-return" component={RentReturnPage} exact />
                <Route path="/user" component={UserPage} exact />
                <Route path="/signin" component={SigninPage} exact />
                <Route path="/signup" component={SignupPage} exact />
            </Switch>
        </Router>
    );
}

export default App;
