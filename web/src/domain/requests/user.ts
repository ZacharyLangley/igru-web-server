import {
    createPromiseClient,
    createConnectTransport,
} from '@bufbuild/connect-web';
import { CreateUserResponse } from '../../client/authentication/v1/user_pb';

import {
    UserService
} from '../../client/authentication/v1/user_connectweb';

const client = createPromiseClient(
    UserService,
    createConnectTransport({
        baseUrl: 'http://172.19.0.8:3000'
    })
);

export function* signUpRequest (email: string, password: string) {
    console.log('signUpRequest', {email, password})
    const response: CreateUserResponse = yield client.createUser({email, password});
    return response;
}

// export function* deleteUserRequest (data: any) {
//     console.log('deleteUserRequest', {data})
//     return client.deleteUser(data);
// }

// export function* updateUserRequest (data: any) {
//     console.log('updateUserRequest', {data})
//     return client.deleteUser(data);
// }

// export function* getUsersRequest (data: any) {
//     console.log('getUsersRequest', {data})
//     return client.deleteUser(data);
// }

// export function* resetPasswordRequest (data: any) {
//     console.log('resetPasswordRequest', {data})
//     return client.deleteUser(data)
// }
