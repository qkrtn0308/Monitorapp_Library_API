import React, { Component } from "react";
import { Container, FormWrap, FormContent, Icon, Form, FormH1, FormLabel, FormInput, FormButton, Text,SigninBg, VideoBg } from "./SigninElements";
import Video from "../../videos/beach.mp4";

class SignIn extends Component {
    state = {
        email: '',
        password: ''
    }
    handleChange  = (e) => {
        this.setState({
            [e.target.name]: e.target.value
        });
        
    }
    render() {

    return(
        <>
            <Container>
                <SigninBg>
                    <VideoBg autoPlay loop muted src={Video} type="vedio/mp4" />
                </SigninBg>
                <FormWrap>
                    <Icon to="/">MONITORAPP</Icon>
                    <FormContent>
                        <Form id="form">
                            <FormH1>Sign in to your account</FormH1>
                            <FormLabel htmlFor='for'>Email</FormLabel>
                            <FormInput name="email" placeholder="gildong.hong@monitorapp.com" value={this.state.email} onChange={this.handleChange} type='email' required/>
                            <FormLabel htmlFor='for'>Password</FormLabel>
                            <FormInput name="password" placeholder="********" value={this.state.password} onChange={this.handleChange} type='password' required/>
                            <FormButton 
                                type='submit'
                                >
                                    Continue</FormButton>
                            <Text to="/help" >Forgot password</Text>
                            <Text to="/signup" >NEW USER</Text>
                        </Form>
                    </FormContent>
                </FormWrap>
            </Container>
        </>
        )
    }
}
export default SignIn
