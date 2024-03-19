import {
    createPromiseClient,
    createConnectTransport,
} from '@bufbuild/connect-web';
import { GetGroupsRequest, UpdateGroupRequest } from 'client/authentication/v1/group_pb';
import { GroupRole } from 'client/authentication/v1/schema_pb';
import { PaginationRequest } from 'src/client/common/v1/pagination_pb';
import { GroupService } from 'src/client/authentication/v1/group_connectweb';
import { getCredentials } from '../utils/auth';

const client = createPromiseClient(GroupService, createConnectTransport({ baseUrl: '' }));

export const createGroupRequest = async (name: string) => {
    if (!name) return;

    const options = await getCredentials();
    if (!options) return;

    return await client.createGroup({name}, options);
}

export const updateGroupRequest = async (group: Partial<UpdateGroupRequest>) => {
    if (!group) return;

    const options = await getCredentials();
    if (!options) return;

    return await client.updateGroup(group, options);
}

export const deleteGroupRequest = async (id: string) => {
    if (!id) return;

    const options = await getCredentials();
    if (!options) return;

    return await client.deleteGroup({id}, options);
}

export const getGroupRequest = async (id: string) => {
    if (!id) return;

    const options = await getCredentials();
    if (!options) return;

    return await client.getGroup({id}, options);
}

export const getAllGroupsRequest = async () => {
    const options = await getCredentials();
    if (!options) return;
    // @ts-ignore
    const pagination: PaginationRequest = {cursor: 0, length: 10};
    const request: Partial<GetGroupsRequest> = {pagination, includeUserGroups: true}

    return await client.getGroups(request, options);
}

export const getGroupsRequest = async () => {
    const options = await getCredentials();
    if (!options) return;

    // @ts-ignore
    const pagination: PaginationRequest = {cursor: 0, length: 10};
    const request: Partial<GetGroupsRequest> = {pagination}

    return await client.getGroups(request, options);
}

export const addGroupMemberRequest = async (groupId: string, userId: string) => {
    if (!groupId || !userId) return;

    const options = await getCredentials();
    if (!options) return;

    return await client.addGroupMember({groupId, userId}, options);
}

export const updateGroupMemberRequest = async (groupId: string, userId: string, role: GroupRole) => {
    if (!groupId || !userId || !role) return;

    const options = await getCredentials();
    if (!options) return;

    return await client.updateGroupMember({groupId, userId, role}, options);
}

export const removeGroupMemberRequest = async (groupId: string, userId: string) => {
    if (!groupId || !userId) return;

    const options = await getCredentials();
    if (!options) return;

    return await client.removeGroupMember({groupId, userId}, options);
}

export const getAllGroupMembersRequest = async (groupId: string) => {
    if (!groupId) return;

    const options = await getCredentials();
    if (!options) return;

    return await client.getGroupMembers({groupId}, options);
}

export const getAllGroupsByUserRequest = async (userId: string) => {
    if (!userId) return;

    const options = await getCredentials();
    if (!options) return;

    return await client.getUserGroups({userId}, options)
}
