import React from "react";
import { Form, Input, Button, Checkbox, Popconfirm } from "antd";
import { UserOutlined, LockOutlined } from "@ant-design/icons";

import styles from "./login.module.scss";

const LoginForm = ({ onSubmit, setSignupPage }) => {
  return (
    <div className={styles.login}>
      <h1>Officer Login</h1>

      <Form
        name="normal_login"
        className="login-form"
        initialValues={{ remember: true }}
        onFinish={onSubmit}
      >
        <Form.Item
          name="username"
          rules={[{ required: true, message: "Please input your Username!" }]}
        >
          <Input
            prefix={<UserOutlined className="site-form-item-icon" />}
            placeholder="Username"
          />
        </Form.Item>
        <Form.Item
          name="password"
          rules={[{ required: true, message: "Please input your Password!" }]}
        >
          <Input
            prefix={<LockOutlined className="site-form-item-icon" />}
            type="password"
            placeholder="Password"
          />
        </Form.Item>
        <Form.Item>
          <Form.Item name="remember" valuePropName="checked" noStyle>
            <Checkbox>Remember me</Checkbox>
          </Form.Item>
        </Form.Item>

        <Form.Item>
          <Button
            type="primary"
            htmlType="submit"
            className="login-form-button"
          >
            Log in
          </Button>
        </Form.Item>
      </Form>

      <Popconfirm
        title="Only an valid officer can perform this action. Do you want to proceed?"
        onConfirm={() => setSignupPage(true)}
        okText="Yes"
        cancelText="No"
      >
        <Button type="primary" danger>
          Create Account
        </Button>
      </Popconfirm>
    </div>
  );
};

export default LoginForm;
