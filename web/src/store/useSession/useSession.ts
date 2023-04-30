import {create} from 'zustand';

import {signInRequest, validateSessionRequest} from './requests';
import {getUserCookie, removeUserCookie, setUserCookie} from '../utils/cookies';

export enum Status {
    IDLE = 'IDLE',
    SUCCESS = 'SUCCESS',
    PENDING = 'PENDING',
    FAILURE = 'FAILURE'
}

interface SessionState {
    sessionValidated: boolean;
    signInStatus: Status;
    error?: any;
}

interface SessionActions {
    signIn: (email?: string, password?: string) => void;
    signOut: () => void;
    validateSession: () => void;
    validateSessionPermissions: () => void;
}

const useSession = create<SessionState & SessionActions>((set) => ({
    sessionValidated: false,
    signInStatus: Status.IDLE,
    error: undefined,

    signIn: async (email?: string, password?: string) => {
        try {
            if (email && password) {
                set({signInStatus: Status.PENDING, error: undefined});
                const response = await signInRequest(email, password);
                if (response) {
                    setUserCookie(response.token);
                    set({sessionValidated: true, signInStatus: Status.SUCCESS});
                    return response.user;
                } else set({signInStatus: Status.FAILURE})
            }
        } catch (error) {
            set({sessionValidated: false, signInStatus: Status.FAILURE, error});
        }
    },
    signOut: () => {
        removeUserCookie()
        set({sessionValidated: false, signInStatus: Status.IDLE, error: undefined});
    },
    validateSession: async () => {
        try {
            const token = await getUserCookie();
            if (token) {
                set({signInStatus: Status.PENDING});
                const response = await validateSessionRequest(token)
                if (response) {
                    set({sessionValidated: true, signInStatus: Status.SUCCESS})
                    return response.user;
                }
                else set({sessionValidated: false, signInStatus: Status.FAILURE})
            }

        } catch (error) {
            set({sessionValidated: false, signInStatus: Status.FAILURE, error});
        }
    },
    validateSessionPermissions: () => {}
}));

export default useSession;