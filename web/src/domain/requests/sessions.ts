import {
    createPromiseClient,
    createConnectTransport,
    CallOptions,
} from '@bufbuild/connect-web';
import { put } from 'redux-saga/effects';

import { GetSessionsRequest, GetSessionsResponse } from 'src/client/authentication/v1/session_pb';
import { SessionService } from 'src/client/authentication/v1/session_connectweb';
import { PaginationRequest } from 'src/client/common/v1/pagination_pb';
import { signOutSuccessAction } from 'src/domain/actions/sessions';

const client = createPromiseClient(
    SessionService,
    createConnectTransport({
        baseUrl: '',
    })
);

export function* signInRequest (email: string, password: string) {
    var sessionToken = ""
    yield client.createSession({email, password}, {
        onHeader: (header: Headers) => {
            const token = header.get("session")
            if (token) sessionToken = token
        }});
    return sessionToken;
}

export function* signOutRequest () {
    yield put(signOutSuccessAction());
    return;
}

export function* validateSessionRequest (token: string) {
    // @ts-ignore
    const pagination: PaginationRequest = {cursor: 1, length: 10};
    const request: Partial<GetSessionsRequest> = {pagination}
    const options: CallOptions = {headers: {session: token}}
    const response: GetSessionsResponse = yield client.getSessions(request, options);

    return response;
}