import {UserActionTypes} from '../types/user';
import {User} from '../interfaces/user';
import {SetUserInformationAction} from '../interfaces/user';
export interface UserState {
  user?: User;
}

const INITIAL_STATE: UserState = {
  user: undefined,
};

export default function users(
  state: UserState = INITIAL_STATE,
  action: SetUserInformationAction
) {
  console.log(action);
  switch (action.type) {
    case UserActionTypes.SET_USER_INFORMATION_ACTION:
      return {
        ...state,
        user: action.payload,
      };
    default:
      return state;
  }
}
