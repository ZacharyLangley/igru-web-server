import React, { useEffect } from 'react';

import Button from '../../../../../../../common/components/Button/Button';
import Card from '../../../../../../../common/components/Card/Card';

import './styles.scss';
import useGroup from 'src/store/useGroup/useGroup';
import useApp from 'src/store/useApp/useApp';
import { Group } from 'client/authentication/v1/schema_pb';

interface GroupListProps {}

const GroupList: React.FC<GroupListProps> = () => {
    const {groups, getGroups, deleteGroup} = useGroup();
    const {toggleAddUserToGroupModal} = useApp();
    useEffect(() => {
        if (!groups) getGroups();
        else return;
    }, [getGroups, groups]);

    const onRefresh = async () => {
        getGroups();
    }

    const onDelete = (groupId: string, name: string) => deleteGroup(groupId)

    const addUserToGroup = (group: Group) => {
        toggleAddUserToGroupModal?.(group);
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
                        {groups?.map((group) => {
                            return (
                                <div className='group-list-item' key={group.id}>
                                    <div className='item'><span className='label'>{'Name:'}</span><span className='value'>{group.name}</span></div>
                                    <div className='item'><span className='label'>{'Members:'}</span><span className='value'>{group.numMembers.toString()}</span></div>
                                    <div className='item'><span className='label'>{'Created:'}</span><span className='value'>{group.createdAt?.toJsonString()}</span></div>
                                    <div className='actions'>
                                        <Button title={'Add a User'} color={'primary'} onClick={() => addUserToGroup(group)} style={{marginRight: 8}}/>
                                        <Button disable={group.name === 'admin@admin'} title={'Delete'} color={'danger'} onClick={() => onDelete(group.id, group.name)}/>
                                    </div>
                                </div>
                            );
                        })}
                    </div>
 
            </Card>
        </div>
    )
}

export default GroupList;