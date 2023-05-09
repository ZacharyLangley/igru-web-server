import {
    createPromiseClient,
    createConnectTransport,
    CallOptions,
} from '@bufbuild/connect-web';

import { GetSessionsRequest, GetSessionUserResponse } from 'src/client/authentication/v1/session_pb';
import { SessionService } from 'src/client/authentication/v1/session_connectweb';
import { PaginationRequest } from 'src/client/common/v1/pagination_pb';

const client = createPromiseClient(
    SessionService,
    createConnectTransport({
        baseUrl: '',
    })
);

export const signInRequest = async (email: string, password: string) => {
    var sessionToken = ""
    const response = await client.createSession({email, password}, {
        onHeader: (header: Headers) => {
            const token = header.get("session")
            if (token) sessionToken = token
        }});
    return {token: sessionToken, user: response.user};
}

export const validateSessionRequest = async (token: string) => {
    // @ts-ignore
    const pagination: PaginationRequest = {cursor: 1, length: 10};
    const request: Partial<GetSessionsRequest> = {pagination}
    const options: CallOptions = {headers: {session: token}}
    const response: GetSessionUserResponse = await client.getSessionUser(request, options);

    return response;
}