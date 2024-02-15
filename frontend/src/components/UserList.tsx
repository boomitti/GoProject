// src/components/UserList.tsx
import React, { useState, useEffect } from 'react';
import { Table, Button, Space, Modal, Form, Input } from 'antd';
import axios from 'axios';

interface User {
  ID: number;
  userName: string;
  email: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
}

interface UserListProps {
  shouldRefresh: boolean;
}

const UserList: React.FC<UserListProps> = ({ shouldRefresh }) => {
  const [users, setUsers] = useState<User[]>([]);
  const [deleteUserId, setDeleteUserId] = useState<number | null>(null);
  const [deleteModalVisible, setDeleteModalVisible] = useState(false);
  const [editUserId, setEditUserId] = useState<number | null>(null);
  const [editModalVisible, setEditModalVisible] = useState(false);
  const [form] = Form.useForm();

  useEffect(() => {
    // Fetch users from the backend API using Axios
    axios.get('http://localhost:8080/users')
      .then(response => {
        setUsers(response.data);
      })
      .catch(error => console.error('Error fetching users:', error));
  }, [shouldRefresh]);

  const columns = [
    { title: 'User ID', dataIndex: 'ID', key: 'ID' },
    { title: 'Username', dataIndex: 'userName', key: 'userName' },
    { title: 'Email', dataIndex: 'email', key: 'email' },
    { title: 'Created At', dataIndex: 'CreatedAt', key: 'CreatedAt' },
    { title: 'Updated At', dataIndex: 'UpdatedAt', key: 'UpdatedAt' },
    { title: 'Deleted At', dataIndex: 'DeletedAt', key: 'DeletedAt' },
    {
      title: 'Action',
      key: 'action',
      render: (text: any, record: User) => (
        <Space size="middle">
          <Button type="primary" onClick={() => showEditModal(record.ID)}>Edit</Button>
          <Button danger onClick={() => showDeleteModal(record.ID)}>Delete</Button>
        </Space>
      ),
    },
  ];

  const handleEdit = () => {
    if (editUserId) {
      form
        .validateFields()
        .then(values => {
          // Make a PUT request to the backend API
          axios.put(`http://localhost:8080/users/${editUserId}`, values)
            .then(response => {
              console.log('User updated successfully:', response.data);
              // Fetch updated users after edit
              axios.get('http://localhost:8080/users')
                .then(response => {
                  setUsers(response.data);
                })
                .catch(error => console.error('Error fetching users:', error));
            })
            .catch(error => console.error('Error updating user:', error))
            .finally(() => {
              setEditUserId(null);
              setEditModalVisible(false);
              form.resetFields();
            });
        })
        .catch(error => console.error('Validation error:', error));
    }
  };

  const showEditModal = (userID: number) => {
    setEditUserId(userID);
    setEditModalVisible(true);

    // Fetch user details for pre-filling the form
    axios.get(`http://localhost:8080/users/${userID}`)
      .then(response => {
        form.setFieldsValue(response.data);
      })
      .catch(error => console.error('Error fetching user details:', error));
  };

  const handleCancelEdit = () => {
    setEditUserId(null);
    setEditModalVisible(false);
    form.resetFields();
  };

  const showDeleteModal = (userID: number) => {
    setDeleteUserId(userID);
    setDeleteModalVisible(true);
  };

  const handleDelete = () => {
    if (deleteUserId) {
      // Make a DELETE request to the backend API
      axios.delete(`http://localhost:8080/users/${deleteUserId}`)
        .then(response => {
          console.log('User deleted successfully:', response.data);
          // Fetch updated users after deletion
          axios.get('http://localhost:8080/users')
            .then(response => {
              setUsers(response.data);
            })
            .catch(error => console.error('Error fetching users:', error));
        })
        .catch(error => console.error('Error deleting user:', error))
        .finally(() => {
          setDeleteUserId(null);
          setDeleteModalVisible(false);
        });
    }
  };

  const handleCancelDelete = () => {
    setDeleteUserId(null);
    setDeleteModalVisible(false);
  };

  return (
    <div>
      <Table dataSource={users} columns={columns} />

      <Modal
        title="Delete User"
        open={deleteModalVisible}
        onOk={handleDelete}
        onCancel={handleCancelDelete}
        okText="Delete"
        cancelText="Cancel"
      >
        <p>Are you sure you want to delete this user?</p>
      </Modal>

      <Modal
        title="Edit User"
        open={editModalVisible}
        onOk={handleEdit}
        onCancel={handleCancelEdit}
        okText="Save"
        cancelText="Cancel"
      >
        <Form form={form} layout="vertical">
          <Form.Item label="Username" name="userName" rules={[{ required: true, message: 'Please enter username' }]}>
            <Input />
          </Form.Item>
          <Form.Item label="Email" name="email" rules={[{ required: true, message: 'Please enter email' }]}>
            <Input />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
};

export default UserList;
