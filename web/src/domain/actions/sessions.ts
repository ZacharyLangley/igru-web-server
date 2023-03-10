import {SessionActionTypes} from '../types/sessions';
import {
    DispatchSignInAction,
    SignInSuccessAction,
    SignInFailureAction,
    DispatchValidateSessionAction,
    ValidateSessionSuccessAction,
    ValidateSessionFailureAction,
    DispatchSignOutAction,
    SignOutSuccessAction,
    DispatchValidateSessionPermissionsAction,
    SetValidatedSessionPermissionsAction,
} from '../interfaces/sessions';
import {
    CreateSessionRequest,
    GetSessionsRequest,
    CheckSessionPermissionsRequest,
    CheckSessionPermissionsResponse,
} from 'src/client/authentication/v1/session_pb';

export const dispatchSignInAction = (payload: Partial<CreateSessionRequest>): DispatchSignInAction => ({
    type: SessionActionTypes.DISPATCH_SIGN_IN_ACTION,
    payload
});

export const signInSuccessAction = (): SignInSuccessAction => ({type: SessionActionTypes.SIGN_IN_SUCCESS_ACTION});

export const signInFailureAction = (): SignInFailureAction => ({type: SessionActionTypes.SIGN_IN_FAILURE_ACTION});

export const dispatchValidateSessionAction = (payload: Partial<GetSessionsRequest>): DispatchValidateSessionAction => ({
    type: SessionActionTypes.DISPATCH_VALIDATE_SESSION_ACTION,
    payload
});

export const validateSessionSuccessAction = (): ValidateSessionSuccessAction => ({type: SessionActionTypes.VALIDATE_SESSION_SUCCESS_ACTION})

export const validateSessionFailureAction = (): ValidateSessionFailureAction => ({type: SessionActionTypes.VALIDATE_SESSION_FAILURE_ACTION})

export const dispatchSignOutAction = (): DispatchSignOutAction => ({
    type: SessionActionTypes.DISPATCH_SIGN_OUT_ACTION,
});

export const signOutSuccessAction = (): SignOutSuccessAction => ({type: SessionActionTypes.SIGN_OUT_SUCCESS_ACTION});

export const dispatchValidateSessionPermissionsAction = (payload: Partial<CheckSessionPermissionsRequest>): DispatchValidateSessionPermissionsAction => ({ 
    type: SessionActionTypes.DISPATCH_VALIDATE_SESSION_PERMISSIONS_ACTION,
    payload
});

export const setValidatedSessionPermissionsAction = (payload: Partial<CheckSessionPermissionsResponse>): SetValidatedSessionPermissionsAction => ({
    type: SessionActionTypes.SET_VALIDATED_SESSION_PERMISSIONS_ACTION,
    payload
});
