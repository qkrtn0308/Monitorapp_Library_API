import React from "react";
import {BookInfoContainer,Titlep, BookBtn, BookBtnLink, BookInfoH1, BookInfoWrapper, BookInfoCard,ImgWrap, Img1,Img2,Img3,Img4,Img5,Img6, BookInfoH2, BookInfoP} from './BookInfoElements'
import img1 from '../../images/Just_do_it.png';
import img2 from '../../images/Programing_with_best_teacher.png';
import img3 from '../../images/little_prince.png';
import img4 from '../../images/Doge_go_mars.png';
import img5 from '../../images/The_universe.png';
import img6 from '../../images/Sad_Monster.png';

const BookInfoSection = () => {
    return(
        <BookInfoContainer id='bookinfo'>
            <BookInfoH1>Monthly Recommendation</BookInfoH1>
            <Titlep>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam ullamcorper, turpis vel blandit aliquam, orci leo gravida arcu, id tempor neque urna in diam. </Titlep>
            <BookInfoWrapper>
                <BookInfoCard>
                    <ImgWrap>
                        <Img1 src={img1} />
                        <BookInfoH2>Just Do It</BookInfoH2>
                        <BookInfoP>Self-development</BookInfoP>
                    </ImgWrap>
                </BookInfoCard>
                <BookInfoCard >
                    <ImgWrap>
                        <Img2 src={img2} />
                        <BookInfoH2>Programing with best teacher</BookInfoH2>
                        <BookInfoP>Programing</BookInfoP>
                    </ImgWrap>
                </BookInfoCard>
                <BookInfoCard>
                    <ImgWrap>
                        <Img3 src={img3} />
                        <BookInfoH2>little prince</BookInfoH2>
                        <BookInfoP>Literature</BookInfoP>
                    </ImgWrap>
                </BookInfoCard>
                <BookInfoCard>
                    <ImgWrap>
                        <Img4 src={img4}/>
                        <BookInfoH2>Coin and money</BookInfoH2>
                        <BookInfoP>Investment</BookInfoP>
                    </ImgWrap>
                </BookInfoCard>
                <BookInfoCard>
                    <ImgWrap>
                        <Img5 src={img5}/>    
                        <BookInfoH2>The Universe</BookInfoH2>
                        <BookInfoP>Science</BookInfoP>
                    </ImgWrap>
                </BookInfoCard>
                <BookInfoCard>
                    <ImgWrap>
                        <Img6 src={img6}/>        
                        <BookInfoH2>Sad Monster</BookInfoH2>
                        <BookInfoP>Fiction</BookInfoP>
                    </ImgWrap>
                </BookInfoCard>
            </BookInfoWrapper>
            <BookBtn>
                <BookBtnLink to="/book">Get more !</BookBtnLink>
            </BookBtn>
        </BookInfoContainer>
    )

}

export default BookInfoSection