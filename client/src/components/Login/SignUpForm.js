import React from "react";
import { Form, Input, Button, Alert } from "antd";

import styles from "./login.module.scss";

function SinUpForm({ onSubmit, setSignupPage, usernameExists }) {
  return (
    <div className={styles.signup}>
      <h1>Create Account</h1>

      {usernameExists ? (
        <Alert
          message="Username Taken"
          description="The username you chose is already taken. Please chose another one"
          type="error"
          style={{ marginBottom: "2rem" }}
        />
      ) : (
        ""
      )}

      <Form name="basic" initialValues={{ remember: true }} onFinish={onSubmit}>
        <Form.Item
          label="Name"
          name="name"
          rules={[{ required: true, message: "Please input your Name!" }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Username"
          name="username"
          rules={[{ required: true, message: "Please input your username!" }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Email ID"
          name="email"
          rules={[{ required: true, message: "Please input your Email ID!" }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Password"
          name="password"
          rules={[{ required: true, message: "Please input your password!" }]}
        >
          <Input.Password />
        </Form.Item>

        <Form.Item
          label="Admin Secret"
          name="secret"
          rules={[{ required: true, message: "Please input your Secret!" }]}
        >
          <Input.Password />
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>

      <Button
        type="primary"
        htmlType="submit"
        onClick={() => setSignupPage(false)}
        danger
      >
        Login
      </Button>
    </div>
  );
}

export default SinUpForm;
