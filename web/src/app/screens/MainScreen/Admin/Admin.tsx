import React from 'react';

import './styles.scss';
import CreateGroup from './components/groups/CreateGroup/CreateGroup';

export interface AdminProps {}

const Admin: React.FC<AdminProps> = () => {
  return (
    <div className="admin-screen-container">
      <CreateGroup />
    </div>
  );
};

export default Admin;
