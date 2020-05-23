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

  // Image Upload Functions
  let [carImage, setCarImage] = useState("");
  let [uploadError, setUploadError] = useState(false);

  const handleUploadError = (error) => {
    setUploadError(true);
    console.error(error);
  };

  const handleUploadSuccess = (filename) => {
    setCarImage(filename);

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
          image: carImage,
        },
      })
      .then((res) => {
        if (res.data.code === 200) {
          setSubmitted(true);
          setError(false);
          setCaseID(res.data.id);
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
        <Submitted caseID={caseID} />
      ) : (
        <>
          <h1>Fill in the form to create a Report</h1>

          <UserForm
            onFinish={onFinish}
            onFinishFailed={onFinishFailed}
            error={error}
            handleUploadError={handleUploadError}
            handleUploadSuccess={handleUploadSuccess}
          />
        </>
      )}
    </div>
  );
};

export default Report;
