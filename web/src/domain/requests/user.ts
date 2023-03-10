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
        baseUrl: '',
    })
);

export function* signUpRequest (email: string, password: string) {
    const response: CreateUserResponse = yield client.createUser({email, password});
    return response;
}
