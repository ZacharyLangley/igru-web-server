import { SessionActionTypes, UserSignInStatus } from "../../domain/types/sessions";

export interface SessionState {
  sessionToken?: string;
  sessionValidated: boolean;
  userSignInStatus: UserSignInStatus;
}

const INITIAL_STATE: SessionState = { sessionValidated: false, userSignInStatus: UserSignInStatus.IDLE };

export default function users(
  state: SessionState = INITIAL_STATE,
  action: any
) {
  switch (action.type) {
    case SessionActionTypes.SIGN_IN_SUCCESS_ACTION:
      return {
        ...state,
        sessionValidated: true,
        userSignInStatus: UserSignInStatus.SUCCESS
      }
    case SessionActionTypes.SIGN_IN_FAILURE_ACTION:
      return {
        ...state,
        sessionValidated: false,
        userSignInStatus: UserSignInStatus.FAILURE
      }
    case SessionActionTypes.VALIDATE_SESSION_SUCCESS_ACTION:
      return {
        ...state,
        sessionValidated: true,
      }
    case SessionActionTypes.SIGN_OUT_SUCCESS_ACTION:
    case SessionActionTypes.VALIDATE_SESSION_FAILURE_ACTION:
      return {
        ...state,
        sessionValidated: false,
        userSignInStatus: UserSignInStatus.IDLE,
      }
    default:
      return state;
  }
}
