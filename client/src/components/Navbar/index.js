import React from "react";
import { Link } from "gatsby";

import styles from "./navbar.module.scss";

function Navbar() {
  return (
    <div className={styles.container}>
      <Link to="/" className={styles.logo}>
        <img src={require("../../images/logo.png")} alt="Carjacked Logo" />
        <h1>Car Jacked</h1>
      </Link>
      <div className={styles.links}>
        <Link className={styles.link + " " + styles.report} to="/report">
          File a Report
        </Link>
        <Link className={styles.link + " " + styles.login} to="/login">
          Officer Login
        </Link>
      </div>
    </div>
  );
}

export default Navbar;
