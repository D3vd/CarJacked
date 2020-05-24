import React from "react";
import { Row, Col } from "antd";

import Status from "./Status";

import styles from "./publicCase.module.scss";

function Case({ publicCase }) {
  let { user, car } = publicCase;
  return (
    <div className={styles.case}>
      <Status status={publicCase.active} />
      <h1>{user.name} &apos;s Case </h1>
      <div className={styles.contact}>
        <h5>
          <b>Email ID:</b> {user.email}
        </h5>
        <h5>
          <b>Phone:</b> {user.phone}
        </h5>
      </div>
      <Row className={styles.car}>
        <Col className={styles.details} lg={12} xs={24}>
          <h2>Car Details</h2>
          <h3>
            <b>Model:</b> {car.model}
          </h3>
          <h3>
            <b>Registration No:</b> {car.regNo}
          </h3>
          <h3>
            <b>Color:</b> {car.color}
          </h3>
          <h3>
            <b>Last Seen Date:</b> {car.lastSeen}
          </h3>
          <h3>
            <b>Last Seen Location:</b> {car.location}
          </h3>
          <h3>
            <b>Description:</b> {car.description}
          </h3>
        </Col>
      </Row>
    </div>
  );
}

export default Case;
