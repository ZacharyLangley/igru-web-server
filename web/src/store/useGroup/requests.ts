import {
    createPromiseClient,
    createConnectTransport,
} from '@bufbuild/connect-web';
import { GetGroupsRequest, UpdateGroupRequest } from 'client/authentication/v1/group_pb';
import { GroupRole } from 'client/authentication/v1/schema_pb';
import { PaginationRequest } from 'src/client/common/v1/pagination_pb';
import { GroupService } from 'src/client/authentication/v1/group_connectweb';

const client = createPromiseClient(GroupService, createConnectTransport({ baseUrl: '' }));

export const createGroupRequest = async (name: string) => {
    if (!name) return;
    // @ts-ignore
    return await client.createGroup({name});
}

export const updateGroupRequest = async (group: Partial<UpdateGroupRequest>) => {
    if (!group) return;
    return await client.updateGroup(group);
}

export const deleteGroupRequest = async (id: string) => {
    if (!id) return;
    return await client.deleteGroup({id});
}

export const getGroupRequest = async (id: string) => {
    if (!id) return;
    return await client.getGroup({id});
}

export const getAllGroupsRequest = async () => {
    // @ts-ignore
    const pagination: PaginationRequest = {cursor: 0, length: 10};
    const request: Partial<GetGroupsRequest> = {pagination, includeUserGroups: true}

    return await client.getGroups(request);
}

export const getGroupsRequest = async () => {
    // @ts-ignore
    const pagination: PaginationRequest = {cursor: 0, length: 10};
    const request: Partial<GetGroupsRequest> = {pagination}

    return await client.getGroups(request);
}

export const addGroupMemberRequest = async (groupId: string, userId: string) => {
    if (!groupId || !userId) return;
    return await client.addGroupMember({groupId, userId});
}

export const updateGroupMemberRequest = async (groupId: string, userId: string, role: GroupRole) => {
    if (!groupId || !userId || !role) return;
    return await client.updateGroupMember({groupId, userId, role});
}

export const removeGroupMemberRequest = async (groupId: string, userId: string) => {
    if (!groupId || !userId) return;
    return await client.removeGroupMember({groupId, userId});
}

export const getAllGroupMembersRequest = async (groupId: string) => {
    if (!groupId) return;
    return await client.getGroupMembers({groupId});
}

export const getAllGroupsByUserRequest = async (userId: string) => {
    if (!userId) return;
    return await client.getUserGroups({userId})
}
