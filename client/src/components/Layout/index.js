import React, { useEffect } from "react";
import PropTypes from "prop-types";
import { Helmet } from "react-helmet";
import axios from "axios";
import { navigate } from "gatsby";
import firebase from "firebase/app";

import firebaseConfig from "../../firebase/config";

import "antd/dist/antd.css";
import "./global.scss";

import Navbar from "../Navbar";

const Layout = ({ children }) => {
  useEffect(() => {
    axios.get("http://localhost:8080/ping").then((res) => {
      console.log(res.data);
    });
    let token = localStorage.getItem("token");
    if (token !== null) navigate("/dashboard");

    // Initialize Firebase App
    if (!firebase.apps.length) {
      firebase.initializeApp(firebaseConfig);
    }
  }, []);

  return (
    <>
      <Helmet>
        <html lang="en" />
        <title>Car Jacked || Stolen Cars Reporting Site</title>
        <meta
          name="description"
          content="Report your stolen car and contact a police officer immediately"
        />
        <meta name="keywords" content="Cars, Stolen, Theft, Police Officer" />
      </Helmet>
      <Navbar />
      <main>{children}</main>
    </>
  );
};

Layout.propTypes = {
  children: PropTypes.node.isRequired,
};

export default Layout;
