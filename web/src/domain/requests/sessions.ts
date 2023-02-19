import {
    createPromiseClient,
    createConnectTransport,
} from '@bufbuild/connect-web';
import { SessionService } from 'src/client/authentication/v1/session_connectweb';

const client = createPromiseClient(
    SessionService,
    createConnectTransport({
        baseUrl: '',
    })
);

export function* signInRequest (email: string, password: string) {
    console.log('signInRequest', {email, password})
    var sessionToken = ""
    yield client.createSession({email, password}, {
        onHeader: (header: Headers) => {
            let token = header.get("session")
            if (token) {
                sessionToken = token
            }
        }});
    return sessionToken;
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