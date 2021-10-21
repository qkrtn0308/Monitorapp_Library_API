import React from "react";
import {FooterContainer, FooterWrap, FooterLinksContainer, FooterLinksWrapper, FooterLinksItems, FooterLinksTitle, FooterLink} from './FooterElements'

function Footer() {
    return (
        <FooterContainer>
            <FooterWrap>
                <FooterLinksContainer>
                    <FooterLinksWrapper>
                        <FooterLinksItems>
                            <FooterLinksTitle>About Us</FooterLinksTitle>
                            <FooterLink to="https://www.monitorapp.com/">Company homepage</FooterLink>
                            <FooterLink to="https://github.com/qkrtn0308">who made this</FooterLink>
                            <FooterLink to="/Admin">Admin</FooterLink>
                        </FooterLinksItems>
                        <FooterLinksItems>
                            <FooterLinksTitle>About Me</FooterLinksTitle>
                            <FooterLink to="/signin">profile</FooterLink>
                            <FooterLink to="https://github.com/qkrtn0308">Github</FooterLink>
                            <FooterLink to="/signin">instagram</FooterLink>
                        </FooterLinksItems>
                    </FooterLinksWrapper>
                    <FooterLinksWrapper>
                        <FooterLinksItems>
                            <FooterLinksTitle>About Us</FooterLinksTitle>
                            <FooterLink to="/signin">Company homepage</FooterLink>
                            <FooterLink to="/signin">who made this</FooterLink>
                        </FooterLinksItems>
                        <FooterLinksItems>
                            <FooterLinksTitle>About Me</FooterLinksTitle>
                            <FooterLink to="/signin">profile</FooterLink>
                            <FooterLink to="/signin">Github</FooterLink>
                            <FooterLink to="/signin">instagram</FooterLink>
                        </FooterLinksItems>
                    </FooterLinksWrapper>
                </FooterLinksContainer>
            </FooterWrap>
        </FooterContainer>
    );
}

export default Footer;