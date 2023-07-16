import {create} from 'zustand';

import { UpdateGroupRequest } from 'client/authentication/v1/group_pb';
import { Group, GroupMember, GroupRole } from 'client/authentication/v1/schema_pb';
import {
    createGroupRequest,
    updateGroupRequest,
    deleteGroupRequest,
    getGroupRequest,
    getGroupsRequest,
    getAllGroupsRequest,
    addGroupMemberRequest,
    updateGroupMemberRequest,
    removeGroupMemberRequest,
    getAllGroupMembersRequest,
    getAllGroupsByUserRequest,
} from './requests';

export enum Status {
    IDLE = 'IDLE',
    SUCCESS = 'SUCCESS',
    PENDING = 'PENDING',
    FAILURE = 'FAILURE'
}

interface GroupState {
    groups?: Group[];
    activeUserGroup?: Group;
    selectedGroup?: Group;
    selectedGroupMembers?: GroupMember[];
    groupsBySelectedUser?: Group[];
    error?: any;
}

interface GroupActions {
    setActiveGroup: (group: Group) => void;
    createGroup: (name: string) => void;
    updateGroup: (group: Partial<UpdateGroupRequest>) => void;
    deleteGroup: (id: string) => void;
    getGroup: (id: string) => void;
    getAllGroups: () => void;
    getGroups: () => void;
    addGroupMember: (groupId: string, userId: string) => void;
    updateGroupMember: (groupId: string, userId: string, role: GroupRole) => void;
    removeGroupMember: (groupId: string, userId: string) => void;
    getAllGroupMembers: (groupId: string) => void;
    getAllGroupsByUser: (userId?: string) => void;
}

const useGroup = create<GroupState & GroupActions>((set, get) => ({
    groups: undefined,
    // @ts-ignore
    activeUserGroup: undefined,
    selectedGroup: undefined,
    selectedGroupMembers: undefined,
    groupsBySelectedUser: undefined,
    error: undefined,
    setActiveGroup: async (group: Group) => {
        if (!group) return;
        set({activeUserGroup: group})
    },
    createGroup: async (name: string) => {
        try {
            if (!name) return;
            const response = await createGroupRequest(name);
            if (response) await get().getAllGroups();
        } catch (error) {set({error})}
    },
    updateGroup: async (group: Partial<UpdateGroupRequest>) => {
        try {
            if (!group) return;
            const response = await updateGroupRequest(group);
            if (response) await get().getAllGroups();
        } catch (error) {set({error})}
    },
    deleteGroup: async (id: string) => {
        try {
            if (!id) return;
            const response = await deleteGroupRequest(id);
            if (response) await get().getAllGroups();
        } catch (error) {set({error})}
    },
    getGroup: async (id: string) => {
        try {
            if (!id) return;
            const response = await getGroupRequest(id);
            if (response) set({selectedGroup: response.group})
        } catch (error) {set({error})}
    },
    getGroups: async () => {
        try {
            const response = await getGroupsRequest();
            if (response) set({groups: response.groups})
            if (get().activeUserGroup === undefined) set({ activeUserGroup: response.groups[0] })
        } catch (error) {set({error})}
    },
    getAllGroups: async () => {
        try {
            const response = await getAllGroupsRequest();
            if (response) set({groups: response.groups})
        } catch (error) {set({error})}
    },
    addGroupMember: async (groupId: string, userId: string) => {
        try {
            if (!groupId || !userId) return;
            const response = await addGroupMemberRequest(groupId, userId)
            if (response) await get().getAllGroupsByUser(userId);
        } catch (error) {set({error})}
    },
    updateGroupMember: async (groupId: string, userId: string, role: GroupRole) => {
        try {
            if (!groupId || !userId || !role) return;
            const response = await updateGroupMemberRequest(groupId, userId, role);
            if (response) await get().getAllGroupMembers(groupId);
        } catch (error) {set({error})}
    },
    removeGroupMember: async (groupId: string, userId: string) => {
        try {
            if (!groupId || !userId) return;
            const response = await removeGroupMemberRequest(groupId, userId);
            if (response) await get().getAllGroupsByUser(userId);
        } catch (error) {set({error})}
    },
    getAllGroupMembers: async (groupId: string) => {
        try {
            if (!groupId) return;
            const response = await getAllGroupMembersRequest(groupId);
            if (response) set({selectedGroupMembers: response.members})
        } catch (error) {set({error})}
    },
    getAllGroupsByUser: async (userId?: string) => {
        try {
            if (!userId) return;
            const response = await getAllGroupsByUserRequest(userId);
            if (response) set({groupsBySelectedUser: response?.groups});
        } catch (error) {set({error})}
    },
}));

export default useGroup;