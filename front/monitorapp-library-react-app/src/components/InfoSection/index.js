import React from "react";
import imgA from '../../images/pic_1.png';

import {
    InfoContainer,
    InfoWrapper,
    InfoRow,
    Column1,
    Column2,
    TextWrapper,
    TopLine,
    Heading,
    Subtitle,
    InfoBtn,
    InfoBtnLink,
    ImgWrap,
    Img
} from "./InfoElements";

const InfoSection = () => {
    return (
            <InfoContainer id="infosection" >
                <InfoWrapper>
                    <InfoRow >
                        <Column1>
                            <TextWrapper>
                                <TopLine>ABOUT</TopLine>
                                <Heading>MONITORAPP LIBRARY</Heading>
                                <Subtitle>MONITORAPP library is at the 8th floor. This API helps both admin and user perfectly. Don't waste your precious time for Finding or making up datas. Just Use. Than, you can feel difference.   </Subtitle>
                                <InfoBtn>
                                    <InfoBtnLink to="/signin">About more</InfoBtnLink>
                                </InfoBtn>
                            </TextWrapper>
                        </Column1>
                        <Column2>
                            <ImgWrap>
                                <Img src={ imgA } />
                            </ImgWrap>
                        </Column2>
                        
                    </InfoRow>
                </InfoWrapper>
            </InfoContainer>
    );
};
export default InfoSection;
