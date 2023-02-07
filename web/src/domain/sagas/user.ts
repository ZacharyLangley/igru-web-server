import {takeLatest} from 'redux-saga/effects';

import {UserActionTypes} from '../types/user';
import {
    DispatchSignUpAction,
    DispatchDeleteUserAction,
    DispatchUpdateUserAction,
    DispatchGetUsersAction,
    DispatchResetPasswordAction
  } from '../interfaces/user';
import { signUpRequest } from '../../domain/requests/user';
import { CreateUserResponse } from '../../client/authentication/v1/user_pb';

export function* signUpUser (action: DispatchSignUpAction) {
    try {
        const {email, password} = action.payload;
        if (email && password) {
            const response: CreateUserResponse = yield signUpRequest(email, password);
            yield console.log('signUpUser: ', action, response);
        }
        yield console.log('signUpUser: ', action);
    } catch (e) {
        yield console.log('signUpUser: ', e);
    }
}

export function* deleteUser (action: DispatchDeleteUserAction) {
    try {
        yield console.log('deleteUser: ', action);
    } catch (e) {
        yield console.log('deleteUser: ', e);
    }
}

export function* updateUser (action: DispatchUpdateUserAction) {
    try {
        yield console.log('updateUser: ', action);
    } catch (e) {
        yield console.log('updateUser: ', e);
    }
}

export function* getUsers (action: DispatchGetUsersAction) {
    try {
        yield console.log('getUsers: ', action);
    } catch (e) {
        yield console.log('getUsers: ', e);
    }
}

export function* resetPassword (action: DispatchResetPasswordAction) {
    try {
        yield console.log('resetPassword: ', action);
    } catch (e) {
        yield console.log('resetPassword: ', e);
    }
}

const userSagas = [
    takeLatest(UserActionTypes.DISPATCH_SIGN_UP_ACTION, signUpUser),
    takeLatest(UserActionTypes.DISPATCH_DELETE_USER_ACTION, deleteUser),
    takeLatest(UserActionTypes.DISPATCH_UPDATE_USER_ACTION, updateUser),
    takeLatest(UserActionTypes.DISPATCH_GET_USERS_ACTION, getUsers),
    takeLatest(UserActionTypes.DISPATCH_RESET_PASSWORD_ACTION, resetPassword)
];

export default userSagas