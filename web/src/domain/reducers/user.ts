import { UserActionTypes, UserSignUpStatus } from "../../domain/types/user";

export interface UserState {
  user: any
  userSignUpStatus: UserSignUpStatus;
}

const INITIAL_STATE: UserState = { user: undefined, userSignUpStatus: UserSignUpStatus.IDLE };

export default function users(
  state: UserState = INITIAL_STATE,
  action: any
) {
  switch (action.type) {
    case UserActionTypes.SIGN_UP_SUCCESS_ACTION:
      return {
        ...state,
        userSignUpStatus: UserSignUpStatus.SUCCESS
      }
    case UserActionTypes.SIGN_UP_FAILURE_ACTION:
      return {
        ...state,
        userSignUpStatus: UserSignUpStatus.FAILURE
      }
    case UserActionTypes.SIGN_UP_RESET_ACTION:
      return {
        ...state,
        userSignUpStatus: UserSignUpStatus.IDLE
      }
    default:
      return state;
  }
}
