
import React from "react";
import { Container, FormWrap, FormContent, Icon, Form, FormH1, FormLabel, Input, FormButton } from "./BookPageElements";


const JINBookPage = () => {
    return(
        <Container className="main-form">
            <FormWrap>
                <Icon to="/">MONITORAPP</Icon>
                    <FormContent>
                        <Form id='form'>
                            <FormH1>search test</FormH1>
                            <FormLabel >코드</FormLabel>
                            <Input name="status" type='text' required/>
                            <FormButton type='submit'>검색</FormButton>
                        </Form>
                </FormContent>
            </FormWrap>
    </Container>
    );
}

export default JINBookPage