import styled from "styled-components";
import { Link as LinkR } from "react-router-dom";


export const BookInfoContainer = styled.div`
    height: 1000px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background: #000;

    @media screen and (max-width: 1000px){
        height: 1400px;
    }
    @media screen and (max-width: 768px) {
        height: 2000px;
    }
`
export const BookInfoWrapper = styled.div`
    max-width: auto;
    margin: 0 auto;
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    align-items: center;
    grid-gap: 16px;

    
    @media screen and (max-width: 1000px){
        grid-template-columns: 1fr 1fr;
    }
    @media screen and (max-width: 768px){
        grid-template-columns: 1fr;

    }
`
export const BookInfoCard = styled.div`
    background: #fff;
    display: flex;
    flex-direction: flex-start;
    align-items: center;
    border-radius: 30px;
    max-height: 340px;
    padding: 30px;
    box-shadow: 0 1px 3px rgba(0,0,0,0.2);
    transition: all 0.2s ease-in-out;

    &:hover{
        transform: scale(1.02);
        transition: all 0.2s ease-in-out;
        cursor: pointer;
    }
`

export const ImgWrap = styled.div`
    max-width: 340px;
    height: 200px;
    
`;

export const Img1 = styled.img`
    src: img1;
    alt: "pic1";
    width: 150px;
    height: 150px;
    
`;
export const Img2 = styled.img`
    src: Img2;
    alt: "pic1";
    width: 150px;
    height: 150px;

`
export const Img3 = styled.img`
    src: Img3;
    alt: "pic1";
    width: 150px;
    height: 150px;

`
export const Img4 = styled.img`
    src: Img4;
    alt: "pic1";
    width: 150px;
    height: 150px;

`
export const Img5 = styled.img`
    src: Img5;
    alt: "pic1";
    width: 150px;
    height: 150px;

`
export const Img6 = styled.img`
    src: Img6;
    alt: "pic1";
    width: 150px;
    height: 150px;

`


export const BookInfoH1 = styled.h1`
    text-align: center;
    font-size: 80px;
    color: #fff;

    @media screen and (max-width: 780px){
        width: 500px;
        font-size: 50px;
    }
    @media screen and (max-width: 480px){
        width: 500px;
        font-size:30px;
    }
`
export const Titlep = styled.p`
    font-size: 1rem;
    max-width: 1200px;
    text-align: center;
    font-weight: 100;
    color: #fff;
    margin-bottom: 45px;
`
export const BookInfoH2 = styled.h2`
    font-size: 1rem;
    text-align: center;
    margin-bottom: 10px;
`
export const BookInfoP = styled.p`
    font-weight: 200;
    font-size: 1rem;
    text-align: center;
`
export const BookBtn = styled.nav`
    margin-top: 40px;
    display: flex;
    justify-content: flex-end;
    align-items: center;

    @media screen and (max-width: 768px) {
        display: flex;
    }
`;
export const BookBtnLink = styled(LinkR)`
    border-radius: 50px;
    background: #14abab;
    white-space: nowrap;
    padding: 10px 22px;
    color: #fff;
    font-size: 15px;
    outline: none;
    border: none;
    cursor: pointer;
    text-decoration: none;
    transition: all 0.2s ease-in-out;

    &:hover {
        transition: all 0.2s ease-in-out;
        background: #108989;
    }
`;