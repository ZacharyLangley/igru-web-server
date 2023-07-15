import React, { ChangeEvent, useEffect, useState } from 'react';
import { Form, FormGroup, Input, Label } from 'reactstrap';

import Button from '../../../../../../../common/components/Button/Button';
import Card from '../../../../../../../common/components/Card/Card';

import './styles.scss';
import useGroup from 'src/store/useGroup/useGroup';
import useUser from 'src/store/useUser/useUser';

interface UserGroupListProps {}

const UserGroupList: React.FC<UserGroupListProps> = () => {
    const {user} = useUser();
    const {groupsBySelectedUser, getAllGroupsByUser, deleteGroup} = useGroup();

    useEffect(() => {
        if (!groupsBySelectedUser) getAllGroupsByUser(user?.id);
        else return;
    }, [user]);

    const onRefresh = async () => {
        getAllGroupsByUser(user?.id);
    }

    const onDelete = (groupId: string, name: string) => {
        if (name !== 'admin@admin') {
            deleteGroup(groupId)
        } else return;
    }

    return (
        <div className="group-list-container">
            <Card
                header='All Groups Current User is Assigned'
                footer={
                    <Button title={'Refresh'} onClick={onRefresh}/>
                }>
                    {!groupsBySelectedUser && <div className='empty-data'><span className='message'>{'No Groups Currently Exist'}</span></div>}
                    <div className='group-body-container'>
                        {groupsBySelectedUser && groupsBySelectedUser?.map((group) => {
                            return (
                                <div className='group-list-item' key={group.id}>
                                    <div className='item'><span className='label'>{'Name:'}</span><span className='value'>{group.name}</span></div>
                                    <div className='item'><span className='label'>{'Members:'}</span><span className='value'>{group.numMembers.toString()}</span></div>
                                    <div className='item'><span className='label'>{'Created:'}</span><span className='value'>{group.createdAt?.toJsonString()}</span></div>
                                    <div className='actions'><Button title={'Delete'} color={'danger'} onClick={() => onDelete(group.id, group.name)}/></div>
                                </div>
                            );
                        })}
                    </div>
 
            </Card>
        </div>
    )
}

export default UserGroupList;