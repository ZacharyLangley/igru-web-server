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
} from 'client/authentication/v1/sessions_pb';

export const dispatchSignUpAction = (payload: CreateSessionRequest): DispatchSignInAction => ({
    type: SessionActionTypes.DISPATCH_SIGN_IN_ACTION,
    payload
});

export const setSignUpAction = (payload: CreateSessionResponse): SetSignInAction => ({
    type: SessionActionTypes.SET_SIGN_IN_ACTION,
    payload
});

export const dispatchSignOutAction = (payload: DeleteSessionRequest): DispatchSignOutAction => ({
    type: SessionActionTypes.DISPATCH_SIGN_OUT_ACTION,
    payload
});

export const setSignOutAction = (payload: DeleteSessionResponse): SetSignOutAction => ({ 
    type: SessionActionTypes.SET_SIGN_OUT_ACTION,
    payload
});

export const dispatchValidateSessionAction = (payload: GetSessionsRequest): DispatchValidateSessionAction => ({
    type: SessionActionTypes.DISPATCH_VALIDATE_SESSION_ACTION,
    payload
});

export const setValidatedSessionAction = (payload: GetSessionsResponse): SetValidatedSessionAction => ({
    type: SessionActionTypes.SET_VALIDATED_SESSION_ACTION,
    payload 
});

export const dispatchValidateSessionPermissionsAction = (payload: CheckSessionPermissionsRequest): DispatchValidateSessionPermissionsAction => ({ 
    type: SessionActionTypes.DISPATCH_VALIDATE_SESSION_PERMISSIONS_ACTION,
    payload
});

export const setValidatedSessionPermissionsAction = (payload: CheckSessionPermissionsResponse): SetValidatedSessionPermissionsAction => ({
    type: SessionActionTypes.SET_VALIDATED_SESSION_PERMISSIONS_ACTION,
    payload
});
