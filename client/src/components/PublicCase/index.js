import React, { useEffect, useState } from "react";
import axios from "axios";

import Layout from "../Layout";
import Case from "./Case";
import Error from "./Error";

import styles from "./publicCase.module.scss";

function PublicCase({ id }) {
  let [publicCase, setPublicCase] = useState(null);
  let [error, setError] = useState(false);

  useEffect(() => {
    axios
      .get(`https://carjacked.herokuapp.com/case/id/${id}`)
      .then((res) => {
        setPublicCase(res.data);
      })
      .catch((err) => {
        console.log(err);
        setError(true);
      });
  }, [id]);

  return (
    <Layout>
      <div className={styles.container}>
        {error ? (
          <Error />
        ) : publicCase ? (
          <div>
            <Case publicCase={publicCase} />
          </div>
        ) : (
          ""
        )}
      </div>
    </Layout>
  );
}

export default PublicCase;
