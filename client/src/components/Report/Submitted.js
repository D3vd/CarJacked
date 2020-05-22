import React from "react";

import styles from "./report.module.scss";

function Submitted({ caseID }) {
  return (
    <div className={styles.submitted}>
      <h1>Your report has been successfully Submitted!</h1>
      <h2>
        An email with Information about your Case has been sent your Email ID.
        Please check it for more information.
      </h2>
      <h3>Your Case ID is {caseID}</h3>
    </div>
  );
}

export default Submitted;
