import React, { useEffect } from 'react';
import { GroupRole } from '../../../../../../../client/authentication/v1/schema_pb'
import Button from '../../../../../../../common/components/Button/Button';
import Card from '../../../../../../../common/components/Card/Card';

import './styles.scss';
import useUser from 'src/store/useUser/useUser';

interface UserListProps {}

const UserList: React.FC<UserListProps> = () => {
    const {users, getUsers, deleteUser} = useUser();

    useEffect(() => {
        if (!users) getUsers();
        else return;
    });

    const onRefresh = async () => {
        getUsers();
    }

    return (
        <div className="user-list-container">
            <Card
                header='All Users'
                footer={
                    <Button title={'Refresh'} onClick={onRefresh} color={'secondary'}/>
                }>
                    {!users && <div className='empty-data'><span className='message'>{'No Users Currently Exist'}</span></div>}
                    <div className='user-body-container'>
                        {users && users?.map((user) => {
                            return (
                                <div className='user-list-item' key={user.id}>
                                <div className='item'><span className='label'>{'Name:'}</span><span className='value'>{user.fullName}</span></div>
                                    <div className='item'><span className='label'>{'Email:'}</span><span className='value'>{user.email}</span></div>
                                    <div className='item'><span className='label'>{'Is Admin:'}</span><span className='value'>{user.globalRole === GroupRole.ADMIN ? 'True' : 'False'}</span></div>
                                    <div className='item'><span className='label'>{'Created:'}</span><span className='value'>{user.createdAt?.toJsonString()}</span></div>
                                    <div className='actions'><Button disable={user?.email === 'admin@admin'} title={'Delete'} color={'danger'} onClick={() => deleteUser(user.id)}/></div>
                                </div>
                            );
                        })}
                    </div>
 
            </Card>
        </div>
    )
}

export default UserList;