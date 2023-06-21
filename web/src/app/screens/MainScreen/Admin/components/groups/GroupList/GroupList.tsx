import React, { ChangeEvent, useEffect, useState } from 'react';
import { Form, FormGroup, Input, Label } from 'reactstrap';

import Button from '../../../../../../../common/components/Button/Button';
import Card from '../../../../../../../common/components/Card/Card';

import './styles.scss';
import useGroup from 'src/store/useGroup/useGroup';

interface GroupListProps {}

const GroupList: React.FC<GroupListProps> = () => {
    const {groups, getAllGroups, deleteGroup} = useGroup();

    useEffect(() => {
        if (!groups) getAllGroups();
        else return;
    }, []);

    const onRefresh = async () => {
        getAllGroups();
    }

    const onDelete = (groupId: string, name: string) => {
        if (name !== 'admin@admin') {
            deleteGroup(groupId)
        } else return;
    }

    return (
        <div className="group-list-container">
            <Card
                header='All Groups'
                footer={
                    <Button title={'Refresh'} onClick={onRefresh}/>
                }>
                    {!groups && <div className='empty-data'><span className='message'>{'No Groups Currently Exist'}</span></div>}
                    <div className='group-body-container'>
                        {groups && groups?.map((group) => {
                            return (
                                <div className='group-list-item'>
                                    <div className='item'><span className='label'>{'ID:'}</span><span className='value'>{group.id}</span></div>
                                    <div className='item'><span className='label'>{'Name:'}</span><span className='value'>{group.name}</span></div>
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

export default GroupList;