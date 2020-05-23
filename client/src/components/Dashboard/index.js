import React, { useEffect, useState } from "react";
import { Layout, Menu } from "antd";
import { navigate } from "gatsby";

const { Header, Content } = Layout;

import styles from "./dashboard.module.scss";

function Dashboard() {
  let [token, setToken] = useState("");

  // Validate if the toke is set
  useEffect(() => {
    let token = localStorage.getItem("token");
    if (token == null) navigate("/");
    setToken(token);
  }, []);

  return (
    <div className={styles.container}>
      <Layout className={styles.layout}>
        <Header style={{ position: "fixed", zIndex: 1, width: "100%" }}>
          {/* <div className="logo" /> */}
          <Menu theme="light" mode="horizontal" defaultSelectedKeys={["2"]}>
            <Menu.Item key="1">nav 1</Menu.Item>
            <Menu.Item key="2">nav 2</Menu.Item>
            <Menu.Item key="3">nav 3</Menu.Item>
          </Menu>
        </Header>
        <Content
          className="site-layout"
          style={{ padding: "0 50px", marginTop: 64, minHeight: "100vh" }}
        >
          <div
            className="site-layout-background"
            style={{ padding: 24, minHeight: 380 }}
          >
            Content
          </div>
        </Content>
      </Layout>
    </div>
  );
}

export default Dashboard;
