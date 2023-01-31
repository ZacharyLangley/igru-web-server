import {Action} from 'redux';
import {UserActionTypes} from '../types/user';
import {
  CreateUserRequest, CreateUserResponse,
  DeleteUserRequest, DeleteUserResponse,
  UpdateUserRequest, UpdateUserResponse,
  GetUsersRequest, GetUsersResponse,
  ResetUserPasswordRequest, ResetUserPasswordResponse
} from 'client/authentication/v1/user_pb';

export interface DispatchSignUpAction extends Action {
  type: UserActionTypes.DISPATCH_SIGN_UP_ACTION;
  payload: CreateUserRequest;
}

export interface SetSignUpAction extends Action {
  type: UserActionTypes.SET_SIGN_UP_ACTION;
  payload: CreateUserResponse;
}

export interface DispatchDeleteUserAction extends Action {
  type: UserActionTypes.DISPATCH_DELETE_USER_ACTION;
  payload: DeleteUserRequest;
}

export interface SetDeleteUserAction extends Action {
  type: UserActionTypes.SET_DELETE_USER_ACTION;
  payload: DeleteUserResponse;
}

export interface DispatchUpdateUserAction extends Action {
  type: UserActionTypes.DISPATCH_UPDATE_USER_ACTION;
  payload: UpdateUserRequest;
}

export interface SetUpdateUserAction extends Action {
  type: UserActionTypes.SET_UPDATE_USER_ACTION;
  payload: UpdateUserResponse
}

export interface DispatchGetUsersAction extends Action {
  type: UserActionTypes.DISPATCH_GET_USERS_ACTION;
  payload: GetUsersRequest
}

export interface SetGetUsersAction extends Action {
  type: UserActionTypes.SET_GET_USERS_ACTION
  payload: GetUsersResponse
}

export interface DispatchResetPasswordAction extends Action {
  type: UserActionTypes.DISPATCH_RESET_PASSWORD_ACTION;
  payload: ResetUserPasswordRequest
}

export interface SetResetPasswordAction extends Action {
  type: UserActionTypes.SET_RESET_PASSWORD_ACTION;
  payload: ResetUserPasswordResponse
}