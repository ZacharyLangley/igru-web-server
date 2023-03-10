import {call, put, takeLatest} from 'redux-saga/effects';
import {SessionActionTypes} from '../types/sessions';

import {
    DispatchSignInAction,
    DispatchValidateSessionPermissionsAction,
} from '../interfaces/sessions';
import { signInRequest, signOutRequest, validateSessionRequest } from '../../domain/requests/sessions';
import { signInFailureAction, signInSuccessAction, validateSessionFailureAction, validateSessionSuccessAction } from '../../domain/actions/sessions';
import { getUserCookie, removeUserCookie, setUserCookie } from 'src/domain/utils/cookies';
import { GetSessionsResponse } from 'src/client/authentication/v1/session_pb';

export function* signIn (action: DispatchSignInAction) {
    try {
        const {email, password} = action.payload;
        if (email && password) {
            const response: string = yield call(signInRequest, email, password);
            if (!response) yield put(signInFailureAction())
            else {
                yield call(setUserCookie, response);
                yield put(signInSuccessAction())
            }
        }
    } catch (e) {
        yield put(signInFailureAction())
    }
}

// TODO: Needs delete user session logic
export function* signOut () {
    try {
        yield call(removeUserCookie);
        yield call(signOutRequest);
    } catch (e) {
        yield console.log('signOut: ', e);
    }
}

// Needs Expiration Logic 
export function* validateSession () {
    try {
        const token: string = yield getUserCookie();
        if (!token) yield put(validateSessionFailureAction());
        else {
            const response: GetSessionsResponse = yield call(validateSessionRequest, token)
            if (!response) yield put(validateSessionFailureAction());
            else yield put(validateSessionSuccessAction())
        }
    } catch (e) {
        yield put(validateSessionFailureAction());
    }
}

export function* validateSessionPermissions (action: DispatchValidateSessionPermissionsAction) {
    try {
        yield console.log('validateSessionPermissions: ', action);
    } catch (e) {
        yield console.log('validateSessionPermissions: ', e);
    }
}

const sessionsSagas = [
    takeLatest(SessionActionTypes.DISPATCH_SIGN_IN_ACTION, signIn),
    takeLatest(SessionActionTypes.DISPATCH_SIGN_OUT_ACTION, signOut),
    takeLatest(SessionActionTypes.DISPATCH_VALIDATE_SESSION_ACTION, validateSession),
    takeLatest(SessionActionTypes.DISPATCH_VALIDATE_SESSION_PERMISSIONS_ACTION, validateSessionPermissions)
]

export default sessionsSagas;