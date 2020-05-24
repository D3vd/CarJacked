import React, { useState } from "react";
import axios from "axios";

import UserForm from "./UserForm";
import Submitted from "./Submitted";

import styles from "./report.module.scss";

const Report = () => {
  let [submitted, setSubmitted] = useState(false);
  let [error, setError] = useState(false);
  let [caseID, setCaseID] = useState("");
  let [submittedData, setSubmittedData] = useState({});

  const onFinish = (values) => {
    axios
      .post("https://carjacked.herokuapp.com/case", {
        user: {
          name: values.name,
          email: values.email,
          phone: values.phone,
        },
        car: {
          color: values.color,
          regNo: values.regNo,
          model: values.model,
          lastSeen: values.lastSeenDate.format("YYYY-MM-DD"),
          location: values.location,
          description: values.description,
        },
      })
      .then((res) => {
        if (res.data.code === 200) {
          setSubmitted(true);
          setError(false);
          setCaseID(res.data.id);
          setSubmittedData(res.data);
        } else {
          console.log(res);
          setError(true);
        }
      })
      .catch((err) => {
        console.log(err);
        setError(true);
      });
  };

  const onFinishFailed = (errorInfo) => {
    console.log("Failed:", errorInfo);
  };

  return (
    <div className={styles.container}>
      {submitted ? (
        <Submitted caseID={caseID} data={submittedData} />
      ) : (
        <>
          <h1>Fill in the form to create a Report</h1>

          <UserForm
            onFinish={onFinish}
            onFinishFailed={onFinishFailed}
            error={error}
          />
        </>
      )}
    </div>
  );
};

export default Report;
