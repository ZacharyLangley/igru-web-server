import { Group } from 'client/authentication/v1/schema_pb';
import {create} from 'zustand';

interface AppState {
    isAddUserToGroupModalOpen?: {
        value?: boolean;
        selectedGroup?: Group;
    };

}

interface AppActions {
    toggleAddUserToGroupModal: (group?: Group) => void;
}

const useApp = create<AppState & AppActions>((set, get) => ({
    isAddUserToGroupModalOpen: {
        toggleValue: false,
        selectedGroup: undefined
    },
    toggleAddUserToGroupModal: (group?: Group) => {
        const value = get().isAddUserToGroupModalOpen?.value;
        set({isAddUserToGroupModalOpen: {
            value: !value,
            selectedGroup: group
        }})
    }
}))

export default useApp;