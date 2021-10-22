import styled from "styled-components";

export const BookContainer = styled.div`
    background-color: #cdf3d9;
    display: flex;
    height: 1000px;
    display: flex;
`
export const Form = styled.div`
    position: relative;
    background: #fff;
    max-width: 600px;
    height: 300px;
    width: 100%;
    z-index: 1;
    display: grid;
    margin: 100px auto;
    padding: 80px 32px;
    border-radius: 4px;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.9);

    @media screen and (max-width: 400px){
        padding: 32px 32px;
    }
`
export const FormInput = styled.input`
    padding: 16px 16px;
    margin-bottom: 32px;
    border: none;
    border-radius: 4px;
    background: #dedede;
`
export const FormButton= styled.button`
    background: #108989;
    padding: 16px 0;
    border: none;
    border-radius: 4px;
    color: #fff;
    font-size: 20px;
    cursor: pointer;
`