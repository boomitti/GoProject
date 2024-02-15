// src/components/UserForm.tsx
import React from 'react';
import { Form, Input, Button } from 'antd';
import axios from 'axios';

interface UserFormProps {
  onFinish: (values: any) => void;
  onUserCreated: () => void;
}

const UserForm: React.FC<UserFormProps> = ({ onFinish, onUserCreated }) => {
  const [form] = Form.useForm();

  const createUserHandler = (values: any) => {
    // Make a POST request to create a new user
    axios.post('http://localhost:8080/users', values)
      .then(response => {
        console.log('User created successfully:', response.data);
        // Call the parent onFinish callback to handle any additional logic
        onFinish(values);
        onUserCreated();

        // Reset the form fields
        form.resetFields();
      })
      .catch(error => console.error('Error creating user:', error));
  };

  return (
    <Form form={form} onFinish={createUserHandler} layout="vertical">
      <Form.Item label="Username" name="userName" rules={[{ required: true, message: 'Please enter username' }]}>
        <Input />
      </Form.Item>
      <Form.Item label="Email" name="email" rules={[{ required: true, message: 'Please enter email' }]}>
        <Input />
      </Form.Item>
      <Form.Item label="Password" name="password" rules={[{ required: true, message: 'Please enter password' }]}>
        <Input.Password />
      </Form.Item>
      <Form.Item>
        <Button type="primary" htmlType="submit">
          Create User
        </Button>
      </Form.Item>
    </Form>
  );
};

export default UserForm;
