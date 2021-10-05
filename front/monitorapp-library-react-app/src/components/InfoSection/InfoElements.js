import styled from "styled-components";
import { Link as LinkR } from "react-router-dom";

export const InfoContainer = styled.div`
    color: #fff;
    background: ${({ lightBg }) => (lightBg ? "#000" : "#fff")};
    @media screen and (max-width: 768px) {
        padding: 100px 0;
    }
`;
export const InfoWrapper = styled.div`
    display: grid;
    z-index: 1;
    height: 1000px;
    width: 100%;
    max-width: 1300px;
    margin-top: 100px;
    margin-right: auto;
    margin-left: auto;
    padding: 0 24px;
    justify-content: center;
`;
export const InfoRow = styled.div`
    display: grid;
    grid-auto-columns: minmax(auto, 1fr);
    align-items: center;
    grid-template-areas: ${({ imgStart }) =>
        imgStart ? `'col2 col1'` : `'col1 col2'`};

    @media screen and (max-width: 768px) {
        grid-template-areas: ${({ imgStart }) =>
            imgStart ? `'col1' 'col2'` : `'col1 col1' 'col2 col2'`};
    }
`;
export const Column1 = styled.div`
    margin-bottom: 15px;
    padding: 0 15px;
    grid-area: col1;
`;
export const Column2 = styled.div`
    margin-bottom: 15px;
    padding: 0 15px;
    grid-area: col2;
`;
export const TextWrapper = styled.div`
    max-width: 540px;
    padding-top: 0;
    padding-bottom: 60px;
`;
export const TopLine = styled.p`
    color: #14abab;
    font-size: 30px;
    line-height: 16px;
    font-weight: 700;
    letter-spacing: 1.4px;
    text-transform: uppercase;
    margin-bottom: 16px;
`;
export const Heading = styled.h1`
    margin-bottom: 24px;
    font-size: 80px;
    line-height: 1.1;
    font-weight: 700;
    color: ${({ lightText }) => (lightText ? "#fff" : "#000")};

    @media screen and (max-width: 480px) {
        font-size: 32px;
    }
`;
export const Subtitle = styled.p`
    max-width: 550px;
    margin-bottom: 10px;
    font-size: 25px;
    line-height: 30px;
    //letter-spacing: 5px;
    font-weight: 200;
    color: ${({ lightText }) => (lightText ? "#fff" : "#000")}; ;
`;
export const InfoBtn = styled.nav`
    display: flex;
    align-items: center;

    @media screen and (max-width: 768px) {
        display: flex;
    }
`;
export const InfoBtnLink = styled(LinkR)`
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

export const ImgWrap = styled.div`
    max-width: 555px;
    height: 100%;
`;
export const Img = styled.img`
    src:  imgA;
    alt: "pic1";
    width: 100%;
    display: center;
    margin: 0 0 10px 0;
    padding-right: 0;
`;
