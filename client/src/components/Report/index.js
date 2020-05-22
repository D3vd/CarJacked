import React from "react";

import UserForm from "./UserForm";

import styles from "./report.module.scss";

const Report = () => {
  const onFinish = (values) => {
    console.log("Success:", values);
  };

  const onFinishFailed = (errorInfo) => {
    console.log("Failed:", errorInfo);
  };

  return (
    <div className={styles.container}>
      <h1>Fill in the form to create a Report</h1>
      <UserForm onFinish={onFinish} onFinishFailed={onFinishFailed} />
    </div>
  );
};

export default Report;
