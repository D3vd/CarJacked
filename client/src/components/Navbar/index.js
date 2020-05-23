import React from "react";
import { Link } from "gatsby";

import styles from "./navbar.module.scss";
import { Row, Col } from "antd";

function Navbar() {
  return (
    <Row className={styles.container}>
      <Col lg={12} xs={24}>
        <Link to="/" className={styles.logo}>
          <img src={require("../../images/logo.png")} alt="Carjacked Logo" />
          <h1>Car Jacked</h1>
        </Link>
      </Col>
      <Col lg={12} xs={24}>
        <div className={styles.links}>
          <Link className={styles.link + " " + styles.login} to="/login/">
            Officer Login
          </Link>
          <Link className={styles.link + " " + styles.report} to="/report/">
            File a Report
          </Link>
        </div>
      </Col>
    </Row>
  );
}

export default Navbar;
