import {create} from 'zustand';

import {
    signUpRequest,
    getUsersRequest,
    deleteUsersRequest,
} from './requests';
import { User } from 'client/authentication/v1/schema_pb';

export enum Status {
    IDLE = 'IDLE',
    SUCCESS = 'SUCCESS',
    PENDING = 'PENDING',
    FAILURE = 'FAILURE'
}

interface UserState {
    users?: User[];
    user?: User;
    signUpStatus: Status
    error?: any
}

interface UserActions {
    signUp: (email?: string, password?: string) => Promise<Status>;
    setUser: (user: any) => void;
    resetSignUpStatus: () => void;
    deleteUser: (userId: any) => void;
    updateUser: () => void;
    getUsers: () => void;
    resetPassword: () => void;
}

const useUser = create<UserState & UserActions>((set) => ({
    users: undefined,
    user: undefined,
    signUpStatus: Status.IDLE,
    error: undefined,

    signUp: async (email?: string, password?: string) => {
        try {
            if (email && password) {
                set({signUpStatus: Status.PENDING, error: undefined});
                const response = await signUpRequest(email, password);
                if (response) {
                    set({signUpStatus: Status.SUCCESS});
                    return Status.SUCCESS
                }
                else {
                    set({signUpStatus: Status.FAILURE})
                    return Status.FAILURE
                }
            }
            return Status.IDLE;
        } catch (error) {
            set({signUpStatus: Status.FAILURE, error});
            return Status.FAILURE
        }
    },
    setUser: (user: any) => set({user}),
    resetSignUpStatus: () => {
        set({signUpStatus: Status.IDLE, error: undefined})
    },
    deleteUser: async (userId: string) => {
        try {
            await deleteUsersRequest(userId);
        } catch (error) {set({error})}
    },
    updateUser: () => {},
    getUsers: async () => {
        try {
            const response = await getUsersRequest();
            if (response) set({users: response.users})
        } catch (error) {set({error})}
    },
    resetPassword: () => {}
}));

export default useUser;