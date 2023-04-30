import {create} from 'zustand';

import {signUpRequest} from './requests';

export enum Status {
    IDLE = 'IDLE',
    SUCCESS = 'SUCCESS',
    PENDING = 'PENDING',
    FAILURE = 'FAILURE'
}

interface UserState {
    user?: any;
    signUpStatus: Status
    error?: any
}

interface UserActions {
    signUp: (email?: string, password?: string) => void;
    setUser: (user: any) => void;
    resetSignUpStatus: () => void;
    deleteUser: () => void;
    updateUser: () => void;
    getUsers: () => void;
    resetPassword: () => void;
}

const useUser = create<UserState & UserActions>((set) => ({
    user: undefined,
    signUpStatus: Status.IDLE,
    error: undefined,

    signUp: async (email?: string, password?: string) => {
        try {
            if (email && password) {
                set({signUpStatus: Status.PENDING, error: undefined});
                const response = await signUpRequest(email, password);
                console.log(response)
                if (response) set({signUpStatus: Status.SUCCESS});
                else set({signUpStatus: Status.FAILURE})
            }
        } catch (error) {
            set({signUpStatus: Status.FAILURE, error});
        }
    },
    setUser: (user: any) => set({user}),
    resetSignUpStatus: () => {
        set({signUpStatus: Status.IDLE, error: undefined})
    },
    deleteUser: () => {},
    updateUser: () => {},
    getUsers: () => {},
    resetPassword: () => {}
}));

export default useUser;