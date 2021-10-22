import { React } from "react";
import styled from "styled-components"
import { Layout, Menu, Typography  } from 'antd';
const { Title } = Typography;

const { Header, Content, Footer, Sider } = Layout;

export default function Main () {
    return (
        <Layout>
        <Header className="header">
        <div className="logo" />
        <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['2']}>
            <Menu.Item key="logo">MONITORAPP</Menu.Item>
            <Menu.Item key="menu 1">User</Menu.Item>
            <Menu.Item key="menu 2">Book</Menu.Item>
            <Menu.Item key="menu 3">Log</Menu.Item>
        </Menu>
        </Header>

        
        <Content style={{  display: "flex",minHeight: 700}}>
            <Title style={{ 
                display: "flex",
                fontSize: 200, 
                textAlign: "center"}}>
                    Library<br/> 
                    Admin<br/>
                    Tab
                </Title>
        </Content>
        <Footer style={{ 
            textAlign: 'center', 
            backgroundColor:"#010101", 
            color: "#fff", 
            minHeight: 50 }}>
                Monitorapp Library @2021 MONITORAPP INTERNSHIP PROJECT
            </Footer>
        </Layout>
        
        
    );
}