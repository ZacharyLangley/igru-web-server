import {
    createPromiseClient,
    createConnectTransport,
} from '@bufbuild/connect-web';

import {
    UserService
} from '../../client/authentication/v1/user_connectweb';

const client = createPromiseClient(
    UserService,
    createConnectTransport({
        baseUrl: '',
    })
);

export const signUpRequest = async (email: string, password: string) => {
    return await client.createUser({email, password});
}
