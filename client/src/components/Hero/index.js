import React from "react";
import { Link } from "gatsby";

import styles from "./hero.module.scss";

function Hero() {
  return (
    <div className={styles.container}>
      <div className={styles.content}>
        <h1>We can help you recover your missing vehicle</h1>
        <h3>
          File a missing car report now and have direct access to the officer
          working on the case until your car has been recovered. Keep track and
          get updates about your case by email.
        </h3>
        <Link to="/report" className={styles.link}>
          Create a Report &rarr;
        </Link>
      </div>
      <div className={styles.image}>
        <img src={require("../../images/hero.png")} alt="Hero.jpg" />
      </div>
    </div>
  );
}

export default Hero;
