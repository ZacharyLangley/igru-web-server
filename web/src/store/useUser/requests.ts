import {
    createPromiseClient,
    createConnectTransport,
} from '@bufbuild/connect-web';
import {
    GetUsersRequest,
    DeleteUserRequest
} from 'client/authentication/v1/user_pb';

import {
    UserService
} from '../../client/authentication/v1/user_connectweb';

const client = createPromiseClient(UserService, createConnectTransport({ baseUrl: '' }));

export const signUpRequest = async (email: string, password: string) => {
    return await client.createUser({email, password});
}

export const getUsersRequest = async () => {
    // @ts-ignore
    const pagination: PaginationRequest = {cursor: 0, length: 10};
    const request: Partial<GetUsersRequest> = {pagination}

    return await client.getUsers(request);
}

export const deleteUsersRequest = async (userId: string) => {
    // @ts-ignore
    const request: Partial<DeleteUserRequest> = {userId: userId}

    return await client.deleteUser(request);
}
