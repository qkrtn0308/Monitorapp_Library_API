import React, {useState} from 'react';
import axios from 'axios';
import {BookContainer, Form, FormInput, FormButton} from './BookPageElements'

const BookPageIndex = ()=>{
    var [data, setData] = useState(null);
    var [text, setText] = useState('');

    const onChange = (e) => {
        setText(e.target.value);
    };
    const CategoryClick = (e) => {
        // 버튼 누르면 토글 리스트
        // 하나 고르면 버튼 value 바뀜 
    };

    const onClick = async()=>{
        try{
            setText(text)
            var url = 'http://localhost:4000/book/code?status=' + text
            //axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';
            console.log(url)
            const response = await axios.get(url)
            console.log(response.data);
            setData(response.data);  
        }catch(e){
            console.log(e);
        }
    }
    return (
        <div>
            <BookContainer>
                <Form >
                    <FormButton type="button" onClick={CategoryClick}> 통합검색</FormButton>
                    <FormInput onChange={onChange} value={text} name="code" placeholder="Type here" type='text' required/>
                    <FormButton type="submit" onClick={onClick}> Search</FormButton>
                </Form>
                {data && <textarea rows={50} value={JSON.stringify(data, null, 2)} readOnly={true} />}
            </BookContainer>
            
        </div>
    );
}

export default BookPageIndex
