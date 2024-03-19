import {create} from 'zustand';

import { createGardenRequest, deleteGardenRequest, getAllGardensRequest, getGardenRequest, updateGardenRequest } from './requests';
import { CreateGardenRequest, UpdateGardenRequest } from '../../client/garden/v1/garden_pb';

export enum Status {
    IDLE = 'IDLE',
    SUCCESS = 'SUCCESS',
    PENDING = 'PENDING',
    FAILURE = 'FAILURE'
}

interface GardenState {
    gardens?: any;
    selectedGarden?: any;
    status: Status;
    error?: any;
}

interface GardenActions {
    getGarden: (id?: string) => void;
    getAllGardens: (groupId?: string) => void;
    createGarden: (garden?: Partial<CreateGardenRequest>) => void;
    updateGarden: (garden?: Partial<UpdateGardenRequest>) => void;
    deleteGarden: (id?: string, groupId?: string) => void;

}

const useGarden = create<GardenState & GardenActions>((set) => ({
    gardens: undefined,
    selectedGarden: undefined,
    status: Status.IDLE,
    error: undefined,
    getGarden: async (id?: string) => {
        try {
            if (!id) return;
            set({status: Status.PENDING, error: undefined});
            const response = await getGardenRequest(id);
            if (response) set({status: Status.SUCCESS, selectedGarden: response});
            else set({status: Status.FAILURE});
        } catch (error) {
            set({status: Status.FAILURE, error});
        }
    },
    getAllGardens: async (groupId?: string) => {
        try {
            if (!groupId) return;
            set({status: Status.PENDING, error: undefined});
            const response = await getAllGardensRequest(groupId);
            if (response) set({status: Status.SUCCESS, gardens: response?.gardens});
            else set({status: Status.FAILURE});
        } catch (error) {
            set({status: Status.FAILURE, error});
        }
    },
    createGarden: async (garden?: Partial<CreateGardenRequest>) => {
        try {
            if (!garden) return;
            set({status: Status.PENDING, error: undefined});
            const response = await createGardenRequest(garden);
            if (response) set({status: Status.SUCCESS});
            else set({status: Status.FAILURE});
        } catch (error) {
            set({status: Status.FAILURE, error});
        }
    },
    updateGarden: async (garden?: Partial<UpdateGardenRequest>) => {
        try {
            if (!garden) return;
            set({status: Status.PENDING, error: undefined});
            const response = await updateGardenRequest(garden);
            if (response) set({status: Status.SUCCESS});
            else set({status: Status.FAILURE});
        } catch (error) {
            set({status: Status.FAILURE, error});
        }
    },
    deleteGarden: async (id?: string, groupId?: string) => {
        try {
            if (!id || !groupId) return;
            set({status: Status.PENDING, error: undefined});
            const response = await deleteGardenRequest(id, groupId);
            if (response) {
                const response = await getAllGardensRequest(groupId);
                set({status: Status.SUCCESS, gardens: response?.gardens});
            }
            else set({status: Status.FAILURE});
        } catch (error) {
            set({status: Status.FAILURE, error});
        }
    }
}))

export default useGarden;