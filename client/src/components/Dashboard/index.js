import React, { useEffect, useState } from "react";
import { navigate } from "gatsby";
import axios from "axios";

import Header from "./Header";
import Case from "./Case";
import NoCase from "./NoCase";
import Error from "./Error";

import styles from "./dashboard.module.scss";

function Dashboard() {
  let [token, setToken] = useState("");
  let [activeCase, setActiveCase] = useState(null);
  let [error, setError] = useState(false);

  // Validate if the toke is set
  useEffect(() => {
    let token = localStorage.getItem("token");
    if (token == null) navigate("/");
    setToken(token);

    axios
      .get("http://localhost:8080/admin/getCase", {
        headers: {
          Authorization: "Bearer " + token,
        },
      })
      .then((res) => {
        setActiveCase(res.data.case);
      })
      .catch((err) => {
        console.log(err);
        setError(true);
      });
  }, []);

  const handleSolveCase = () => {
    axios
      .get("http://localhost:8080/admin/solveCase", {
        headers: {
          Authorization: "Bearer " + token,
        },
      })
      .then((res) => {
        console.log(res);
        setActiveCase(null);
      })
      .catch((err) => {
        console.log(err);
        setError(true);
      });
  };

  return (
    <div className={styles.container}>
      <Header />
      {!error ? (
        activeCase ? (
          <Case activeCase={activeCase} solveCase={handleSolveCase} />
        ) : (
          <NoCase />
        )
      ) : (
        <Error />
      )}
    </div>
  );
}

export default Dashboard;
