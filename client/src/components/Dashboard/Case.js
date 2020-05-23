import React from "react";
import { Row, Col, Popconfirm, Button } from "antd";

import styles from "./dashboard.module.scss";

function Case({ activeCase, solveCase }) {
  let { user, car } = activeCase;
  return (
    <div className={styles.case}>
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

        <Col className={styles.image} lg={12} xs={24}>
          <img src={car.image} alt="" />
        </Col>
      </Row>
      <Popconfirm
        title="The Case reporter will be notified that the case has been solved. Do you want to proceed?"
        onConfirm={solveCase}
        okText="Yes"
        cancelText="No"
      >
        <Button type="primary">Case Solved!</Button>
      </Popconfirm>
    </div>
  );
}

export default Case;
