import {UserActionTypes} from '../types/user';
import {
  DispatchSignUpAction,
  SignUpSuccessAction,
  SignUpFailureAction,
  SignUpResetAction,
  DispatchDeleteUserAction,
  SetDeleteUserAction,
  DispatchUpdateUserAction,
  SetUpdateUserAction,
  DispatchGetUsersAction,
  SetGetUsersAction,
  DispatchResetPasswordAction,
  SetResetPasswordAction,
} from '../interfaces/user';
import { 
  CreateUserRequest, CreateUserResponse,
  DeleteUserRequest, DeleteUserResponse,
  UpdateUserRequest, UpdateUserResponse,
  GetUsersRequest, GetUsersResponse,
  ResetUserPasswordRequest, ResetUserPasswordResponse
 } from 'src/client/authentication/v1/user_pb';

export const dispatchSignUpAction = (payload: Partial<CreateUserRequest>): DispatchSignUpAction => ({
  type: UserActionTypes.DISPATCH_SIGN_UP_ACTION,
  payload
});

export const signUpSuccessAction = (payload: Partial<CreateUserResponse>): SignUpSuccessAction => ({
  type: UserActionTypes.SIGN_UP_SUCCESS_ACTION,
  payload
});

export const signUpFailedAction = (): SignUpFailureAction => ({type: UserActionTypes.SIGN_UP_FAILURE_ACTION});

export const signUpResetAction = (): SignUpResetAction => ({type: UserActionTypes.SIGN_UP_RESET_ACTION});

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