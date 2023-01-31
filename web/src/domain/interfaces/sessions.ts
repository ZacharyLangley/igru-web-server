import {Action} from 'redux';
import {SessionActionTypes} from '../types/sessions';
import {
    CreateSessionRequest, CreateSessionResponse, 
    DeleteSessionRequest, DeleteSessionResponse,
    GetSessionsRequest, GetSessionsResponse,
    CheckSessionPermissionsRequest, CheckSessionPermissionsResponse,
} from 'client/authentication/v1/sessions_pb';

export interface DispatchSignInAction extends Action {
    type: SessionActionTypes.DISPATCH_SIGN_IN_ACTION;
    payload: CreateSessionRequest;
  }

export interface SetSignInAction extends Action {
    type: SessionActionTypes.SET_SIGN_IN_ACTION;
    payload: CreateSessionResponse;
}

export interface DispatchSignOutAction extends Action {
    type: SessionActionTypes.DISPATCH_SIGN_OUT_ACTION;
    payload: DeleteSessionRequest;
}

export interface SetSignOutAction extends Action {
    type: SessionActionTypes.SET_SIGN_OUT_ACTION;
    payload: DeleteSessionResponse;
}

export interface DispatchValidateSessionAction extends Action {
    type: SessionActionTypes.DISPATCH_VALIDATE_SESSION_ACTION;
    payload: GetSessionsRequest;
}

export interface SetValidatedSessionAction extends Action {
    type: SessionActionTypes.SET_VALIDATED_SESSION_ACTION;
    payload: GetSessionsResponse
}

export interface DispatchValidateSessionPermissionsAction extends Action {
    type: SessionActionTypes.DISPATCH_VALIDATE_SESSION_PERMISSIONS_ACTION;
    payload: CheckSessionPermissionsRequest;
}   

export interface SetValidatedSessionPermissionsAction extends Action {
    type: SessionActionTypes.SET_VALIDATED_SESSION_PERMISSIONS_ACTION;
    payload: CheckSessionPermissionsResponse
}