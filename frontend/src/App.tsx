// src/App.tsx
import React, {useState} from 'react';
import { Layout, Typography } from 'antd';
import UserList from './components/UserList';
import UserForm from './components/UserForm';

const { Header, Content } = Layout;
const { Title } = Typography;

const App: React.FC = () => {
  const [shouldRefresh, setShouldRefresh] = useState<boolean>(false);

  const handleUserCreated = () => {
    // Set shouldRefresh to trigger a re-render of the UserList component
    setShouldRefresh(true);
  };

  return (
    <Layout>
      <Header>
        <Title style={{ color: 'white' }}>User Management</Title>
      </Header>
      <Content style={{ padding: '20px' }}>
        <Title level={2}>User List</Title>
        <UserList shouldRefresh={shouldRefresh} />
        <Title level={2} style={{ marginTop: '20px' }}>Add User</Title>
        <UserForm onUserCreated={handleUserCreated} onFinish={(values) => console.log('Form values:', values)} />
      </Content>
    </Layout>
  );
}

export default App;
