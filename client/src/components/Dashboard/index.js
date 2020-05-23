import React, { useEffect, useState } from "react";
import { navigate } from "gatsby";

import Header from "./Header";

import styles from "./dashboard.module.scss";

function Dashboard() {
  let [token, setToken] = useState("");

  // Validate if the toke is set
  useEffect(() => {
    let token = localStorage.getItem("token");
    if (token == null) navigate("/");
    setToken(token);
  }, []);

  return (
    <div className={styles.container}>
      <Header />
    </div>
  );
}

export default Dashboard;
