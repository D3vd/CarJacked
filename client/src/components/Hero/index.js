import React from "react";
import { Link } from "gatsby";

import styles from "./hero.module.scss";
import { Row, Col } from "antd";

function Hero() {
  return (
    <Row className={styles.container}>
      <Col lg={12} md={24} className={styles.content}>
        <h1>We can help you recover your missing vehicle</h1>
        <h3>
          File a missing car report now and have direct access to the officer
          working on the case until your car has been recovered. Keep track and
          get updates about your case by email.
        </h3>
        <Link to="/report" className={styles.link}>
          Create a Report &rarr;
        </Link>
      </Col>
      <Col lg={12} md={24} className={styles.image}>
        <img src={require("../../images/hero.png")} alt="Hero.jpg" />
      </Col>
    </Row>
  );
}

export default Hero;
