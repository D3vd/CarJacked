import React, { useState } from "react";
import axios from "axios";
import { navigate } from "gatsby";

import LoginForm from "./LoginForm";
import SignUpForm from "./SignUpForm";

import styles from "./login.module.scss";

function Login() {
  let [signupPage, setSignupPage] = useState(false);
  let [usernameExists, setUsernameExists] = useState(false);

  let [error, setError] = useState(false);

  const loginFormSubmit = (values) => {
    axios
      .post("http://localhost:8080/auth/login", {
        username: values.username,
        password: values.password,
      })
      .then((res) => {
        if (res.data.code === 200) {
          localStorage.setItem("token", res.data.token);
          navigate("/dashboard");
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

  const signUpSubmit = (values) => {
    axios
      .post("http://localhost:8080/auth/signUp", {
        username: values.username,
        password: values.password,
        name: values.name,
        email: values.email,
        secret: values.secret,
      })
      .then((res) => {
        if (res.data.code === 200) {
          window.location = "/login";
        } else if (res.data.code === 400) {
          setUsernameExists(true);
        } else {
          console.log(res);
          setError(true);
        }
      })
      .catch((err) => {
        console.log(err);
        setUsernameExists(true);
      });
  };

  return (
    <div className={styles.container}>
      {signupPage ? (
        <SignUpForm
          onSubmit={signUpSubmit}
          setSignupPage={setSignupPage}
          usernameExists={usernameExists}
        />
      ) : (
        <LoginForm
          onSubmit={loginFormSubmit}
          setSignupPage={setSignupPage}
          error={error}
        />
      )}
    </div>
  );
}

export default Login;
