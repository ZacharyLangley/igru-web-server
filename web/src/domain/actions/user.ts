import {UserActionTypes} from '../types/user';
import {
  DispatchSignInAction,
  DispatchSignOutAction,
  DispatchSignUpAction,
  DispatchResetPasswordAction,
  DispatchCreateSessionAction,
  DispatchValidateSessionAction,
  SetUserInformationAction,
} from '../interfaces/user';

export const dispatchSignInAction = (): DispatchSignInAction => ({
  type: UserActionTypes.DISPATCH_SIGN_IN_ACTION,
});

export const dispatchSignOutAction = (): DispatchSignOutAction => ({
  type: UserActionTypes.DISPATCH_SIGN_OUT_ACTION,
});

export const dispatchSignUpAction = (): DispatchSignUpAction => ({
  type: UserActionTypes.DISPATCH_SIGN_UP_ACTION,
});

export const dispatchResetPasswordAction = (): DispatchResetPasswordAction => ({
  type: UserActionTypes.DISPATCH_RESET_PASSWORD_ACTION,
});

export const dispatchCreateSessionAction = (): DispatchCreateSessionAction => ({
  type: UserActionTypes.DISPATCH_CREATE_SESSION_ACTION,
});

export const dispatchValidateSessionAction =
  (): DispatchValidateSessionAction => ({
    type: UserActionTypes.DISPATCH_VALIDATE_SESSION_ACTION,
  });

export const setUserInformationAction = (): SetUserInformationAction => ({
  type: UserActionTypes.SET_USER_INFORMATION_ACTION,
  payload: {
    id: 'mock_user_id',
    firstName: 'mock_first_name',
    lastName: 'mock_last_name',
  },
});
