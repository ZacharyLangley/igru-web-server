import {UserActionTypes} from '../types/user';
import {
  DispatchSignUpAction,
  SetSignUpAction,
  DispatchDeleteUserAction,
  SetDeleteUserAction,
  DispatchUpdateUserAction,
  SetUpdateUserAction,
  DispatchGetUsersAction,
  SetGetUsersAction,
  DispatchResetPasswordAction,
  SetResetPasswordAction
} from '../interfaces/user';
import { 
  CreateUserRequest, CreateUserResponse,
  DeleteUserRequest, DeleteUserResponse,
  UpdateUserRequest, UpdateUserResponse,
  GetUsersRequest, GetUsersResponse,
  ResetUserPasswordRequest, ResetUserPasswordResponse
 } from 'client/authentication/v1/user_pb';

export const dispatchSignUpAction = (payload: Partial<CreateUserRequest>): DispatchSignUpAction => ({
  type: UserActionTypes.DISPATCH_SIGN_UP_ACTION,
  payload
});

export const setSignUpAction = (payload: Partial<CreateUserResponse>): SetSignUpAction => ({
  type: UserActionTypes.SET_SIGN_UP_ACTION,
  payload
});

export const dispatchDeleteUserAction = (payload: Partial<DeleteUserRequest>): DispatchDeleteUserAction => ({
  type: UserActionTypes.DISPATCH_DELETE_USER_ACTION,
  payload
});

export const setDeleteUserAction = (payload: Partial<DeleteUserResponse>): SetDeleteUserAction => ({
  type: UserActionTypes.SET_DELETE_USER_ACTION,
  payload
});

export const dispatchUpdateUserAction = (payload: Partial<UpdateUserRequest>): DispatchUpdateUserAction => ({
  type: UserActionTypes.DISPATCH_UPDATE_USER_ACTION,
  payload
});

export const setUpdateUserAction = (payload: Partial<UpdateUserResponse>): SetUpdateUserAction => ({
  type: UserActionTypes.SET_UPDATE_USER_ACTION,
  payload
});

export const dispatchGetUsersAction = (payload: Partial<GetUsersRequest>): DispatchGetUsersAction => ({
  type: UserActionTypes.DISPATCH_GET_USERS_ACTION,
  payload
});

export const setGetUsersAction = (payload: Partial<GetUsersResponse>): SetGetUsersAction => ({
  type: UserActionTypes.SET_GET_USERS_ACTION,
  payload
});

export const dispatchResetPasswordAction = (payload: Partial<ResetUserPasswordRequest>): DispatchResetPasswordAction => ({
  type: UserActionTypes.DISPATCH_RESET_PASSWORD_ACTION,
  payload
});

export const setResetPasswordAction = (payload: Partial<ResetUserPasswordResponse>): SetResetPasswordAction => ({
  type: UserActionTypes.SET_RESET_PASSWORD_ACTION,
  payload
});