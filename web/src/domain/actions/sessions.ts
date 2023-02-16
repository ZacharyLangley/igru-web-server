import {SessionActionTypes} from '../types/sessions';
import {
    DispatchSignInAction,
    SetSignInAction,
    DispatchSignOutAction,
    SetSignOutAction,
    DispatchValidateSessionAction,
    SetValidatedSessionAction,
    DispatchValidateSessionPermissionsAction,
    SetValidatedSessionPermissionsAction
} from '../interfaces/sessions';
import {
    CreateSessionRequest, CreateSessionResponse, 
    DeleteSessionRequest, DeleteSessionResponse,
    GetSessionsRequest, GetSessionsResponse,
    CheckSessionPermissionsRequest, CheckSessionPermissionsResponse,
} from 'client/authentication/v1/session_pb';

export const dispatchSignInAction = (payload: Partial<CreateSessionRequest>): DispatchSignInAction => ({
    type: SessionActionTypes.DISPATCH_SIGN_IN_ACTION,
    payload
});

export const setSignInAction = (payload: Partial<CreateSessionResponse>): SetSignInAction => ({
    type: SessionActionTypes.SET_SIGN_IN_ACTION,
    payload
});

export const dispatchSignOutAction = (payload: Partial<DeleteSessionRequest>): DispatchSignOutAction => ({
    type: SessionActionTypes.DISPATCH_SIGN_OUT_ACTION,
    payload
});

export const setSignOutAction = (payload: Partial<DeleteSessionResponse>): SetSignOutAction => ({ 
    type: SessionActionTypes.SET_SIGN_OUT_ACTION,
    payload
});

export const dispatchValidateSessionAction = (payload: Partial<GetSessionsRequest>): DispatchValidateSessionAction => ({
    type: SessionActionTypes.DISPATCH_VALIDATE_SESSION_ACTION,
    payload
});

export const setValidatedSessionAction = (payload: Partial<GetSessionsResponse>): SetValidatedSessionAction => ({
    type: SessionActionTypes.SET_VALIDATED_SESSION_ACTION,
    payload 
});

export const dispatchValidateSessionPermissionsAction = (payload: Partial<CheckSessionPermissionsRequest>): DispatchValidateSessionPermissionsAction => ({ 
    type: SessionActionTypes.DISPATCH_VALIDATE_SESSION_PERMISSIONS_ACTION,
    payload
});

export const setValidatedSessionPermissionsAction = (payload: Partial<CheckSessionPermissionsResponse>): SetValidatedSessionPermissionsAction => ({
    type: SessionActionTypes.SET_VALIDATED_SESSION_PERMISSIONS_ACTION,
    payload
});
