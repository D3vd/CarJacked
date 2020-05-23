import React from "react";
import { Row, Col } from "antd";
import { navigate } from "gatsby";

import styles from "./dashboard.module.scss";

function Header() {
  const logout = () => {
    localStorage.removeItem("token");
    navigate("/");
  };

  return (
    <Row className={styles.header}>
      <Col lg={12} xs={24}>
        <div className={styles.logo}>
          <img src={require("../../images/logo.png")} alt="Carjacked Logo" />
          <h1>Car Jacked</h1>
        </div>
      </Col>
      <Col lg={12} xs={24}>
        <div className={styles.link} onClick={logout}>
          Logout
        </div>
      </Col>
    </Row>
  );
}

export default Header;
