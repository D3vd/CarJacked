import React from "react";
import { Alert } from "antd";

function Status({ status }) {
  return (
    <div>
      {status ? (
        <Alert
          message="Your case is being worked on"
          description="Officers are working on the case. we will update you once the case is solved"
          type="error"
          style={{ marginBottom: "2rem" }}
        ></Alert>
      ) : (
        <Alert
          message="Your case was solved"
          description="Your Case was successfully solved. Check your email for more information"
          type="success"
          style={{ marginBottom: "2rem" }}
        ></Alert>
      )}
    </div>
  );
}

export default Status;
