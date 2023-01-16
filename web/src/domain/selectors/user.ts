import {createSelector} from 'reselect';
import {UserState} from '../reducers/user';

const getUserState = (state: {user: UserState}) => state?.user;

const userSelector = createSelector(
  getUserState,
  (userState) => userState?.user
);

export const userIDSelector = createSelector(userSelector, (user) => user?.id);

export const userFirstNameSelector = createSelector(
  userSelector,
  (user) => user?.firstName
);

export const userLastNameSelector = createSelector(
  userSelector,
  (user) => user?.lastName
);

export const userEmailSelector = createSelector(
  userSelector,
  (user) => user?.email
);
