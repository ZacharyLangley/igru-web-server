import {Action} from 'redux';
import {SessionActionTypes} from '../types/sessions';
import {
    CreateSessionRequest, 
    GetSessionsRequest,
    CheckSessionPermissionsRequest, CheckSessionPermissionsResponse,
} from 'client/authentication/v1/session_pb';

export interface DispatchSignInAction extends Action {
    type: SessionActionTypes.DISPATCH_SIGN_IN_ACTION;
    payload: Partial<CreateSessionRequest>;
  }

export interface SignInSuccessAction extends Action {
    type: SessionActionTypes.SIGN_IN_SUCCESS_ACTION;
}

export interface SignInFailureAction extends Action {
    type: SessionActionTypes.SIGN_IN_FAILURE_ACTION;
}

export interface DispatchValidateSessionAction extends Action {
    type: SessionActionTypes.DISPATCH_VALIDATE_SESSION_ACTION;
    payload: Partial<GetSessionsRequest>;
}

export interface ValidateSessionSuccessAction extends Action {
    type: SessionActionTypes.VALIDATE_SESSION_SUCCESS_ACTION
}

export interface ValidateSessionFailureAction extends Action {
    type: SessionActionTypes.VALIDATE_SESSION_FAILURE_ACTION
}

export interface DispatchSignOutAction extends Action {
    type: SessionActionTypes.DISPATCH_SIGN_OUT_ACTION;
}

export interface SignOutSuccessAction extends Action {
    type: SessionActionTypes.SIGN_OUT_SUCCESS_ACTION;
}

export interface DispatchValidateSessionPermissionsAction extends Action {
    type: SessionActionTypes.DISPATCH_VALIDATE_SESSION_PERMISSIONS_ACTION;
    payload: Partial<CheckSessionPermissionsRequest>;
}   

export interface SetValidatedSessionPermissionsAction extends Action {
    type: SessionActionTypes.SET_VALIDATED_SESSION_PERMISSIONS_ACTION;
    payload: Partial<CheckSessionPermissionsResponse>
}