import React, { useState } from "react";
import axios from "axios";

import LoginForm from "./LoginForm";

import styles from "./login.module.scss";

function Login() {
  let [signupPage, setSignupPage] = useState(false);

  const loginFormSubmit = (values) => {
    var userFormData = new FormData();
    userFormData.set("username", values.username);
    userFormData.set("password", values.password);

    axios({
      method: "post",
      url: "http://localhost:8080/auth/login",
      data: userFormData,
      headers: { "Content-Type": "multipart/form-data" },
    })
      .then((res) => {
        console.log(res);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  return (
    <div className={styles.container}>
      <LoginForm onSubmit={loginFormSubmit} setSignupPage={setSignupPage} />
    </div>
  );
}

export default Login;
