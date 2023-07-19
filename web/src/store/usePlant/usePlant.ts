import { CreatePlantRequest, UpdatePlantRequest } from 'client/garden/v1/plant_pb';
import {create} from 'zustand';
import { createPlantRequest, deletePlantRequest, getAllPlantsRequest, getPlantRequest, updatePlantRequest } from './requests';

export enum Status {
    IDLE = 'IDLE',
    SUCCESS = 'SUCCESS',
    PENDING = 'PENDING',
    FAILURE = 'FAILURE'
}

interface PlantState {
    plants?: any;
    selectedPlant: any;
    status: Status;
    error: any;
}

interface PlantActions {
    getPlant: (id?: string) => void;
    getAllPlants: () => void;
    createPlant: (plant: Partial<CreatePlantRequest>) => void;
    updatePlant: (plant: Partial<UpdatePlantRequest>) => void;
    deletePlant: (id?: string) => void;

}

const usePlant = create<PlantState & PlantActions>((set) => ({
    plants: undefined,
    selectedPlant: undefined,
    status: Status.IDLE,
    error: undefined,
    getPlant: async (id?: string) => {
        try {
            if (!id) return;
            set({status: Status.PENDING, error: undefined});
            const response = await getPlantRequest(id);
            if (response) set({status: Status.SUCCESS, selectedPlant: response});
            else set({status: Status.FAILURE});
        } catch (error) {
            set({status: Status.FAILURE, error});
        }
    },
    getAllPlants: async () => {
        try {
            set({status: Status.PENDING, error: undefined});
            const response = await getAllPlantsRequest();
            if (response) set({status: Status.SUCCESS, plants: response});
            else set({status: Status.FAILURE});
        } catch (error) {
            set({status: Status.FAILURE, error});
        }
    },
    createPlant: async (plant: Partial<CreatePlantRequest>) => {
        try {
            if (!plant) return;
            set({status: Status.PENDING, error: undefined});
            const response = await createPlantRequest(plant);
            if (response) set({status: Status.SUCCESS});
            else set({status: Status.FAILURE});
        } catch (error) {
            set({status: Status.FAILURE, error});
        }
    },
    updatePlant: async (plant: Partial<UpdatePlantRequest>) => {
        try {
            if (!plant) return;
            set({status: Status.PENDING, error: undefined});
            const response = await updatePlantRequest(plant);
            if (response) set({status: Status.SUCCESS});
            else set({status: Status.FAILURE});
        } catch (error) {
            set({status: Status.FAILURE, error});
        }
    },
    deletePlant: async (id?: string) => {
        try {
            if (!id) return;
            set({status: Status.PENDING, error: undefined});
            const response = await deletePlantRequest(id);
            if (response) set({status: Status.SUCCESS});
            else set({status: Status.FAILURE});
        } catch (error) {
            set({status: Status.FAILURE, error});
        }
    }
}))

export default usePlant;