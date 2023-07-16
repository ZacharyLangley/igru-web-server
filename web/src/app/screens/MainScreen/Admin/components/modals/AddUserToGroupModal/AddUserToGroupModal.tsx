import React, { useCallback, useEffect, useMemo, useState } from 'react';

import { Modal } from '../../../../../../../common/components/Modal/Modal';
import useApp from 'src/store/useApp/useApp';
import Button from '../../../../../../../common/components/Button/Button';

import './styles.scss';
import SelectDropdown from '../../../../../../..//common/components/Dropdown/Dropdown';
import useUser from 'src/store/useUser/useUser';
import useGroup from 'src/store/useGroup/useGroup';
import { addGroupMemberRequest, getAllGroupsByUserRequest, removeGroupMemberRequest } from 'src/store/useGroup/requests';
import { Group } from 'client/authentication/v1/schema_pb';

interface AddUserToGroupModalProps {}

export const AddUserToGroupModal = (props: AddUserToGroupModalProps) => {
    const {isAddUserToGroupModalOpen, toggleAddUserToGroupModal} = useApp();
    const {getUsers, users, user} = useUser();
    const {groupsBySelectedUser, addGroupMember, removeGroupMember} = useGroup();
    const [formGroupsBySelectedUser, setFormGroupsBySelectedUser] = useState<Group[]>();

    const [selectedUserId, setSelectedUserId] = useState<string | undefined>(users?.[0].id);
    const isUpdatingActiveUser = (selectedUserId === user?.id)
    const currentGroups = isUpdatingActiveUser ? groupsBySelectedUser : formGroupsBySelectedUser

    const getSelectedGroups = useCallback(async (userId?: string) => {
        if (userId) {
            const groupsByUser = await getAllGroupsByUserRequest(userId);
            if (groupsByUser) {
                setFormGroupsBySelectedUser(groupsByUser?.groups);
            }
        }
    }, [])

    useEffect(() => {
        if (users && !selectedUserId) {
            setSelectedUserId(users?.[0].id);
            getSelectedGroups(users?.[0].id);
        }
    }, [users, getSelectedGroups, selectedUserId]);

    useEffect(() => {
        if (isAddUserToGroupModalOpen?.selectedGroup && isAddUserToGroupModalOpen.value) {
            getUsers();
        }
    }, [selectedUserId, isAddUserToGroupModalOpen, getUsers]);

    const userOptions = useMemo(() => {
        return users?.map((user) => ({
            name: user.email,
            ...user,
        }))
    }, [users])

    const handleUserSelect = async (userId?: string) => {
        if (userId) {
            const response = await getAllGroupsByUserRequest(userId);
            if (response) {
                setFormGroupsBySelectedUser(response?.groups);
            }
            setSelectedUserId(userId);
        }
    }

    const handleSubmit = async () => {
        if (isAddUserToGroupModalOpen?.selectedGroup && selectedUserId) {
            if (isUpdatingActiveUser && user) {
                addGroupMember(isAddUserToGroupModalOpen?.selectedGroup?.id, user.id)
            } else {
                const response = await addGroupMemberRequest(isAddUserToGroupModalOpen?.selectedGroup?.id, selectedUserId)
                if (response) {
                    getSelectedGroups(selectedUserId);
                }
            }
        }
        setSelectedUserId(undefined);
        toggleAddUserToGroupModal();
    }

    const handleCancel = () => {
        setSelectedUserId(undefined);
        toggleAddUserToGroupModal();
    }

    const handleGroupRemoval = async (groupId?: string) => {
        if (isUpdatingActiveUser && isAddUserToGroupModalOpen?.selectedGroup && user && groupId) {
            removeGroupMember(groupId, user?.id)
        } else {
            if (groupId && selectedUserId) {
                await removeGroupMemberRequest(groupId, selectedUserId);
                getSelectedGroups(selectedUserId);
            }
        }
    }

    return (
        <Modal 
            open={isAddUserToGroupModalOpen?.value}
            toggle={toggleAddUserToGroupModal}
            header={`Add a User to ${isAddUserToGroupModalOpen?.selectedGroup?.name}`}
            className={'add-user-to-group-modal'}
            body={
                <div className={'body'}>
                    <SelectDropdown label={'Select a User'} options={userOptions} onSelect={handleUserSelect}/>
                    <div className={'user-group-list'}>
                        <div className={'user-group-title'}>{`Groups Associated With User`}</div>
                        {!currentGroups && <span>{'This User is not Assigned to Any Groups'}</span>}
                        {currentGroups?.map((group) => {
                            return (
                                <div key={group.id} style={{display: 'flex', flexDirection: 'row', justifyContent: 'space-between', marginBottom: 8}}>
                                    <span>
                                        {group.name}
                                    </span>
                                    <Button size={'sm'} title={'Remove'} color='danger' onClick={() => handleGroupRemoval(group?.id)} style={{ marginRight: 8, maxWidth: 80 }}/>
                                </div>
                            )
                        })}
                    </div>
                </div>
            }
            footer={
                <div className={'actions'}>
                    <Button title={'Confirm'} color='primary' onClick={handleSubmit} style={{ marginRight: 8 }}/>
                    <Button title={'Cancel'} color='danger' onClick={handleCancel}/>
                </div> 
            }
        />
    )
}

export default AddUserToGroupModal;