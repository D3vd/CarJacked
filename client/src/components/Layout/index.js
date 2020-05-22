import React from "react";
import PropTypes from "prop-types";
import { Helmet } from "react-helmet";

import "antd/dist/antd.css";
import "./global.scss";

import Navbar from "../Navbar";

const Layout = ({ children }) => {
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
