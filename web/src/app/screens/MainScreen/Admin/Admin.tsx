import React from 'react';

import './styles.scss';
import CreateGroup from './components/groups/CreateGroup/CreateGroup';
import GroupList from './components/groups/GroupList/GroupList';
import UserList from './components/users/UserList/UserList';

export interface AdminProps {}

const Admin: React.FC<AdminProps> = () => {
  return (
    <div className="admin-screen-container">
      <CreateGroup />
      <GroupList />
      <UserList />
    </div>
  );
};

export default Admin;
