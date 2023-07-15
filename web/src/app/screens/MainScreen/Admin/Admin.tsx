import React from 'react';

import './styles.scss';
import CreateGroup from './components/groups/CreateGroup/CreateGroup';
import GroupList from './components/groups/GroupList/GroupList';
import UserList from './components/users/UserList/UserList';
import UserGroupList from './components/users/UserGroupList/UserGroupList';

export interface AdminProps {}

const Admin: React.FC<AdminProps> = () => {
  return (
    <div className="admin-screen-container">
      <div className={'panel'}>
        <div className={'panel-item'}>
          <CreateGroup />
        </div>
        <div className={'panel-item'}>
          <GroupList />
        </div>
      </div>
      <div className={'panel'}>
        <div className={'panel-item'}>
          <UserList />
        </div>
        <div className={'panel-item'}>
          <UserGroupList />
        </div>
      </div>
    </div>
  );
};

export default Admin;
