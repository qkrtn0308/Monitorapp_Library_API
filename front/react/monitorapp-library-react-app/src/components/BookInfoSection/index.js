import React from "react";
import { Button } from "../ButtonElements";

import {BookInfoContainer,BtnWrap, BookInfoH1, BookInfoWrapper, BookInfoCard, Img1,Img2,Img3,Img4,Img5,Img6, BookInfoH2, BookInfoP} from './BookInfoElements'

const BookInfoSection = () => {
    return(
        <BookInfoContainer id='bookinfo'>
            <BookInfoH1>Daily Recommendation</BookInfoH1>
            <BookInfoWrapper>
                <BookInfoCard>
                    <Img1 alt="1" src='../../images/Just_do_it.png'>
                    <BookInfoH2>Just Do It</BookInfoH2>
                    <BookInfoP>자기개발도서</BookInfoP>
                    </Img1>
                </BookInfoCard>
                <BookInfoCard>
                    <Img2 alt="2" src='../../images/Programing_with_best_teacher.png'>
                    <BookInfoH2>Programing with best teacher</BookInfoH2>
                    <BookInfoP>전공도서</BookInfoP>
                    </Img2>
                </BookInfoCard>
                <BookInfoCard>
                    <Img3 alt="3" src='../../images/little_prince.png'>
                    <BookInfoH2>little prince</BookInfoH2>
                    <BookInfoP>문학도서</BookInfoP>
                    </Img3>
                </BookInfoCard>
                <BookInfoCard>
                     <Img4 alt="4" src='../../images/Doge_go_mars.png'>
                    <BookInfoH2>Coin&money</BookInfoH2>
                    <BookInfoP>투자 </BookInfoP>
                    </Img4>
                </BookInfoCard>
                <BookInfoCard>
                     <Img5 alt="5" src='../../images/The_Universe.png'>    
                    <BookInfoH2>The Universe</BookInfoH2>
                    <BookInfoP>시사/상식</BookInfoP>
                    </Img5>
                </BookInfoCard>
                <BookInfoCard>
                    <Img6 alt="6" src='../../images/Sad_Monster.png'>        
                    <BookInfoH2>Sad Monster</BookInfoH2>
                    <BookInfoP>소설</BookInfoP>
                    </Img6>
                </BookInfoCard>
            </BookInfoWrapper>
            <BtnWrap>
                <Button to="home"
                    smooth={true}
                    duration={500}
                    spy={true}
                    exact="true"
                    offset={-80}>GET More !</Button>
            </BtnWrap>
        </BookInfoContainer>
    )

}

export default BookInfoSection