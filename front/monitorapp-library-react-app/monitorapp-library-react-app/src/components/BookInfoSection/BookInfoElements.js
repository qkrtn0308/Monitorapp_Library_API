import styled from "styled-components";

export const BookInfoContainer = styled.div`
    height: 800px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background: #fff;

    @media screen and(max-width: 768px){
        height: 1100px;
    }
    @media screen and(max-width: 180px){
        height: 1100px;
    }
`
export const BookInfoWrapper = styled.div`
    max-width: 1000px;
    margin: 0 auto;
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    align-items: center;
    grid-gap: 16px;
    padding: 0 50px;

    
    @media screen and (max-width: 1000px){
        grid-template-columns: 1fr 1fr;
    }
    @media screen and (max-width: 768px){
        grid-template-columns: 1fr;
        padding: 0 20px;
    }
`
export const BookInfoCard = styled.div`
    background: #fff;
    display: flex;
    flex-direction: flex-start;
    align-items: center;
    border-radius: 10px;
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
export const Img1 = styled.div`
    width: 150;
    height: 150;
    margin-bottom: 10px;
`;
export const Img2 = styled.div`
    width: 150;
    height: 150;
    margin-bottom: 10px;
`;
export const Img3 = styled.div`
    width: 150;
    height: 150;
    margin-bottom: 10px;
`;
export const Img4 = styled.div`
    width: 150;
    height: 150;
    margin-bottom: 10px;
`;
export const Img5 = styled.div`
    width: 150;
    height: 150;
    margin-bottom: 10px;
`;
export const Img6 = styled.div`
    width: 150;
    height: 150;
    margin-bottom: 10px;
`;


export const BookInfoH1 = styled.h1`
    font-size: 4rem;
    color: #000;
    margin-bottom: 64px;

    @media screen and(max-width: 480px){
        font-size: 2rem;
    }
`
export const BookInfoH2 = styled.h2`
    font-size: 1rem;
    margin-bottom: 10px;
`
export const BookInfoP = styled.p`
    font-size: 1rem;
    text-align: center;

`
export const BtnWrap = styled.div`
    margin-top: 40px;
    display: flex;
    justify-content: flex-end;
`;