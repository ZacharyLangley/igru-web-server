import {
    createPromiseClient,
    createConnectTransport,
} from '@bufbuild/connect-web';
import { CreateSessionRequest } from 'client/authentication/v1/sessions_pb';

import { SessionService } from '../../client/authentication/v1/session_connectweb';

const client = createPromiseClient(
    SessionService,
    createConnectTransport({
        baseUrl: '',
    })
);

export function* signInRequest (email: string, password: string) {
    console.log('signUpRequest', {email, password})
    const response: CreateSessionRequest = yield client.createSession({email, password})
    return;
}

// export const signOutRequest = async () => {
//     return;
// }

// export const validateSessionRequest = async () => {
//     return;
// }

// export const validateSessionPermissions = async () => {
//     return;
// }