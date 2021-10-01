import React from "react";
import {
    SidebarContainer,
    Icon,
    CloseIcon,
    SidebarWrapper,
    SidebarMenu,
    SidebarLink,
    SidebtnWrap,
    SidebarRoute,
} from "./SideBarElements";

const Sidebar = ({ isOpen, toggle }) => {
    return (
        <SidebarContainer isOpen={isOpen} onClick={toggle}>
            <Icon onClick={toggle}>
                <CloseIcon />
            </Icon>
            <SidebarWrapper>
                <SidebarMenu>
                    <SidebarLink to="about" onClick={toggle}>
                        ABOUT
                    </SidebarLink>
                    <SidebarLink to="book" onClick={toggle}>
                        BOOK
                    </SidebarLink>
                    <SidebarLink to="rent-return" onClick={toggle}>
                        RENT/RETURN
                    </SidebarLink>
                    <SidebarLink to="user" onClick={toggle}>
                        USER
                    </SidebarLink>
                </SidebarMenu>
                <SidebtnWrap>
                    <SidebarRoute to="/signin" onClick={toggle}>
                        Sign In
                    </SidebarRoute>
                </SidebtnWrap>
            </SidebarWrapper>
        </SidebarContainer>
    );
};

export default Sidebar;
