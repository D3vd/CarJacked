import React from "react";
import { Link } from "gatsby";

import styles from "./publicCase.module.scss";

function Error() {
  return (
    <div className={styles.error}>
      <h1>Error</h1>
      <img src={require("../../images/error.png")} alt="Error" />
      <h2>Invalid Case ID. Please check and try Again</h2>
      <Link to="/">Go Home</Link>
    </div>
  );
}

export default Error;
