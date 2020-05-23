import React from "react";

import styles from "./dashboard.module.scss";

function NoCase({ checkForNewCase }) {
  return (
    <div className={styles.noCase}>
      <h1>There is no Unassigned case at the moment</h1>
      <h3>Enjoy your break</h3>
      <img src={require("../../images/noCase.png")} alt="No Case" />

      <div className={styles.newCase} onClick={checkForNewCase}>
        Check for New Case
      </div>
    </div>
  );
}

export default NoCase;
