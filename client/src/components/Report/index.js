import React, { useState } from "react";
import axios from "axios";
import firebase from "firebase/app";
import "firebase/storage";

import UserForm from "./UserForm";
import Submitted from "./Submitted";

import styles from "./report.module.scss";

const Report = () => {
  let [submitted, setSubmitted] = useState(false);
  let [error, setError] = useState(false);
  let [caseID, setCaseID] = useState("");
  let [submittedData, setSubmittedData] = useState({});

  // Image Upload Functions
  let [carImage, setCarImage] = useState("");
  let [uploadError, setUploadError] = useState(false);

  const handleUploadError = (error) => {
    setUploadError(true);
    console.error(error);
  };

  const handleUploadSuccess = (filename) => {
    setCarImage(filename);
    setUploadError(false);

    firebase
      .storage()
      .ref("carImage")
      .child(filename)
      .getDownloadURL()
      .then((url) => {
        setCarImage(url);
        console.log(url);
      });
  };

  const onFinish = (values) => {
    axios
      .post("http://localhost:8080/case", {
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
          image: carImage,
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
            handleUploadError={handleUploadError}
            handleUploadSuccess={handleUploadSuccess}
            uploadError={uploadError}
          />
        </>
      )}
    </div>
  );
};

export default Report;
