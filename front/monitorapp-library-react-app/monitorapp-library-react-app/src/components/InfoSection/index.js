import React from "react";
import imgA from '../../images/pic_1.png';
import { Button } from "../ButtonElements";
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
    BtnWrap,
    ImgWrap,
    Img
} from "./InfoElements";

const InfoSection = () => {
    return (
        <>
            <InfoContainer >
                <InfoWrapper>
                    <InfoRow >
                        <Column1>
                            <TextWrapper>
                                <TopLine>ABOUT</TopLine>
                                <Heading>MONITORAPP LIBRARY</Heading>
                                <Subtitle>MONITORAPP library is at the 8th floor. This API helps both admin and user perfectly. Don't waste your precious time for Finding or making up datas. Just Use. Than, you can feel difference.   </Subtitle>
                                <BtnWrap>
                                    <Button to="home"
                                    smooth={true}
                                    duration={500}
                                    spy={true}
                                    exact="true"
                                    offset={-80}>find More !</Button>
                                </BtnWrap>
                            </TextWrapper>
                        </Column1>
                        <Column2>
                            <ImgWrap>
                                <Img src={ imgA  } />
                            </ImgWrap>
                        </Column2>
                        
                    </InfoRow>
                </InfoWrapper>
            </InfoContainer>
        </>
    );
};
export default InfoSection;
