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
    NavBtn,
    NavBtnLink,
} from "./NavBarElements";
import { IconContext } from "react-icons/lib";

const Navbar = ({ toggle }) => {
    const [scrollNav, setScrollNav] = useState(false);

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
                            <NavLinks to='/about'>ABOUT</NavLinks>
                        </NavItem>
                        <NavItem>
                            <NavLinks to='/book'>BOOK</NavLinks>
                        </NavItem>
                        <NavItem>
                            <NavLinks to='/rent-return'>RENT/RETURN</NavLinks>
                        </NavItem>
                        <NavItem>
                            <NavLinks to='/user'>USER</NavLinks>
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
