import {createSelector} from 'reselect';
import {SessionState} from '../reducers/session';

const getSessionState = (state: {session: SessionState}) => state?.session;

export const sessionTokenSelector = createSelector(
    getSessionState,
    (sessionState: SessionState) => sessionState.sessionToken
);

export const isSessionValidatedSelector = createSelector(
    getSessionState,
    (sessionState: SessionState) => sessionState.sessionValidated
)

export const userSignInStatusSelector = createSelector(
    getSessionState,
    (sessionState: SessionState) => sessionState.userSignInStatus
)