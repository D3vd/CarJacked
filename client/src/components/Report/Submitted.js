import React from "react";
import { Alert } from "antd";

import styles from "./report.module.scss";

function Submitted({ caseID, data }) {
  return (
    <div className={styles.submitted}>
      <h1>Your report has been successfully Submitted!</h1>
      <h2>
        Keep track of this url to be up to date with the case. <br />
        <a
          href={`https://carjackked.now.sh/case/${caseID}`}
        >{`https://carjacked.now.sh/case/${caseID}`}</a>
      </h2>
      <h3>Your Case ID is {caseID}</h3>
      {data.assigned ? (
        <Alert
          message={`Officer ${data.officer.name} is on your case`}
          description={`${data.officer.name} has been assigned to your case. They will be in contact with you if
          there is any updates.`}
          type="success"
          style={{ marginTop: "2rem" }}
        ></Alert>
      ) : (
        <Alert
          message="No Free Officers at the moment"
          description="There are no free Officers at the moment to take your case. You will
          be notified when an officer is assigned to you."
          type="error"
          style={{ marginTop: "2rem" }}
        ></Alert>
      )}
    </div>
  );
}

export default Submitted;
