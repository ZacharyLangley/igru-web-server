import React from 'react';

import './styles.scss';
import CreateGroup from './components/groups/CreateGroup/CreateGroup';
import GroupList from './components/groups/GroupList/GroupList';

export interface AdminProps {}

const Admin: React.FC<AdminProps> = () => {
  return (
    <div className="admin-screen-container">
      <CreateGroup />
      <GroupList />
    </div>
  );
};

export default Admin;
