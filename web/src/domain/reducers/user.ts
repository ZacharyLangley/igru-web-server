export interface UserState {
  user: any
}

const INITIAL_STATE: UserState = { user: undefined };

export default function users(
  state: UserState = INITIAL_STATE,
  action: any
) {
  switch (action.type) {
    default:
      return state;
  }
}
