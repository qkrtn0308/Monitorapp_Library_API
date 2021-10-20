import React, { useEffect, useState } from "react";
import { FaBars } from "react-icons/fa";
import { animateScroll as scroll } from "react-scroll";
import {
    Nav,
    NavbarContainer,
    NavLogo,
    MobileIcon,
    NavMenu,
    NavItem,
    NavLinks,
    NavLinkr,
    NavBtn,
    NavBtnLink,
} from "./NavBarElements";
import { IconContext } from "react-icons/lib";

const Navbar = ({ toggle }) => {
    const [scrollNav] = useState(false);

    useEffect(() => {
        scroll.scrollToTop()
    })

    const toggleHome = () => {
        scroll.scrollToTop()
    }

    return (
        <>
        <IconContext.Provider value={{color: '#000'}}>
            <Nav scrollNav={scrollNav}>
                <NavbarContainer>
                    <NavLogo to='/' onClick={toggleHome}>MONITORAPP</NavLogo>
                    <MobileIcon onClick={toggle}>
                        <FaBars />
                    </MobileIcon>
                    <NavMenu>
                        <NavItem>
                            <NavLinks to='infosection'>ABOUT</NavLinks>
                        </NavItem>
                        <NavItem>
                            <NavLinkr to='/book'>BOOK</NavLinkr>
                        </NavItem>
                        <NavItem>
                            <NavLinkr to='/rent-return'>RENT/RETURN</NavLinkr>
                        </NavItem>
                        <NavItem>
                            <NavLinkr to='/user'>USER</NavLinkr>
                        </NavItem>
                    </NavMenu>
                    <NavBtn>
                        <NavBtnLink to='/signin'>Sign In</NavBtnLink>
                    </NavBtn>
                </NavbarContainer>
            </Nav>
            </IconContext.Provider>
        </>
    );
};

export default Navbar;
