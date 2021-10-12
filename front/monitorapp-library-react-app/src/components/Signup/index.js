import React from "react";
import { Container, FormWrap, FormContent, Icon, Form, FormH1, FormLabel, FormInput, FormButton, Text,SigninBg, VideoBg } from "./SignupElements";
import Video from "../../videos/beach.mp4";

const SignUp = () => {
    return(
        <>
            <Container>
                <SigninBg>
                    <VideoBg autoPlay loop muted src={Video} type="vedio/mp4" />
                </SigninBg>
                <FormWrap>
                    <Icon to="/">MONITORAPP</Icon>
                    <FormContent>
                        <Form action="#">
                            <FormH1>Sign up new account</FormH1>
                            <FormLabel htmlFor='for'>FirstName</FormLabel>
                            <FormInput type='text' required/>
                            <FormLabel htmlFor='for'>LastName</FormLabel>
                            <FormInput type='text' required/>
                            <FormLabel htmlFor='for'>phone number</FormLabel>
                            <FormInput type='tel' required/>
                            <FormButton type='submit'>Next</FormButton>
                            <Text to="/signin" >I have account</Text>
                        </Form>
                    </FormContent>
                </FormWrap>
            </Container>
        </>
    )
}

export default SignUp