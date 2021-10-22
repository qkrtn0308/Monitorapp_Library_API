import React, {useState} from 'react';
import axios from 'axios';
import { Container, FormWrap, FormContent, Icon, Form, FormH1, FormLabel, FormInput, FormButton, Text, SigninBg, VideoBg} from "./SigninElements";
import Video from "../../videos/beach.mp4";


function SignIn() {
    var [data, setData] = useState(null);
    var [email, setemail] = useState('');
    var [password, setpw] = useState('');

    const onChangeemail = (e) => {
        setemail(e.target.value);
    };
    const onChangepw = (e) => {
        setpw(e.target.value);
    };

    const onClick = async()=>{
        try{
            setemail(email)
            setpw(password)
            const url1 = 'http://localhost:4000/login'
            console.log(url1)
            const response = await axios.post(url1, {
                email: email,
                password: password
            })
            console.log(response.data);
            setData(response.data);  
        }catch(e){
            console.log(e);
        }
    }


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
                            <FormLabel >Email</FormLabel>
                            <FormInput onChange={onChangeemail} value={email} name="email" placeholder="gildong.hong@monitorapp.com" type='email' required/>
                            <FormLabel >Password</FormLabel>
                            <FormInput onChange={onChangepw} value={password} name="password" placeholder="********" type='password' required/>
                            <FormButton onClick={onClick} type="button">Continue</FormButton>
                            <Text to="/help" >Forgot password</Text>
                            <Text to="/signup" >NEW USER</Text>
                            {data && <textarea rows={1} value={JSON.stringify(data, null, 2)} readOnly={true} />}
                        </Form>
                    </FormContent>
                </FormWrap>
            </Container>
        </>
        )
    }

export default SignIn
