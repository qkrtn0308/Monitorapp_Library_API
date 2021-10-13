import React, { useState } from "react";
import { Container, FormWrap, FormContent, Icon, Form, FormH1, FormLabel, FormInput, FormButton, Text,SigninBg, VideoBg } from "./SigninElements";
import Video from "../../videos/beach.mp4";
const SignIn = () => {
        
    return(
        <>
            <Container>
                <SigninBg>
                    <VideoBg autoPlay loop muted src={Video} type="vedio/mp4" />
                </SigninBg>
                <FormWrap>
                    <Icon to="/">MONITORAPP</Icon>
                    <FormContent>
                        <Form action="#" method="get">
                            <FormH1>Sign in to your account</FormH1>
                            <FormLabel htmlFor='for'>Email</FormLabel>
                            <FormInput type='email' required/>
                            <FormLabel htmlFor='for'>Password</FormLabel>
                            <FormInput type='password' required/>
                            <FormButton type='submit'>Continue</FormButton>
                            <Text to="/help" >Forgot password</Text>
                            <Text to="/signup" >NEW USER</Text>
                        </Form>
                    </FormContent>
                </FormWrap>
            </Container>
        </>
    )
}

export default SignIn