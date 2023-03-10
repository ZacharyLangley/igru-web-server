import {createSelector} from 'reselect';
import {UserState} from '../reducers/user';

const getUserState = (state: {user: UserState}) => state?.user;

const userSelector = createSelector(
  getUserState,
  (userState) => userState?.user
);

export const userIDSelector = createSelector(userSelector, (user) => user?.id);

export const userEmailSelector = createSelector(
  userSelector,
  (user) => user?.email
);

export const userSignUpStatusSelector = createSelector(
  getUserState,
  (userState) => userState.userSignUpStatus 
)
