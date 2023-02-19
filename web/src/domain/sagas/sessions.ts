import {takeLatest} from 'redux-saga/effects';
import {SessionActionTypes} from '../types/sessions';

import {
    DispatchSignInAction,
    DispatchSignOutAction,
    DispatchValidateSessionAction,
    DispatchValidateSessionPermissionsAction,
} from '../interfaces/sessions';
import { signInRequest } from '../../domain/requests/sessions';

export function* signIn (action: DispatchSignInAction) {
    try {
        const {email, password} = action.payload;
        if (email && password) {
            let resp: string = yield signInRequest(email, password);
            yield console.log('signIn: ', resp);
        }
    } catch (e) {
        yield console.log('signIn: ', e);
    }
}

export function* signOut (action: DispatchSignOutAction) {
    try {
        yield console.log('signOut: ', action);
    } catch (e) {
        yield console.log('signOut: ', e);
    }
}

export function* validateSession (action: DispatchValidateSessionAction) {
    try {
        yield console.log('validateSession: ', action);
    } catch (e) {
        yield console.log('validateSession: ', e);
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