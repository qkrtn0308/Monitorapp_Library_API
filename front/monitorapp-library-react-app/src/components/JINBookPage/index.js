import React, {useState} from "react";
import { Container, FormWrap, FormContent, Icon, Form, FormH1, Input, FormButton, Text } from "./BookPageElements";
import axios from 'axios'

const JINBookPage = async() => {
    axios.get('http://localhost:4000/', {
        params: {
            keyword: this.state
        }
    });

    return(
        <Container className="main-form">
            <FormWrap>
                <Icon to="/">MONITORAPP</Icon>
                    <FormContent>
                        <Form id='form'>
                            <FormH1>search test</FormH1>
                            <Text to="/book" >키워드로 찾기</Text>
                            <Text to="/book/code">상태 코드</Text>
                            <Input name="status" type='text' required/>
                            <FormButton type='button'>검색</FormButton>
                        </Form>
                </FormContent>
            </FormWrap>
    </Container>
    );
}

export default JINBookPage