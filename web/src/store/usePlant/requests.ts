import {
    createPromiseClient,
    createConnectTransport,
} from '@bufbuild/connect-web';

import { CreatePlantRequest, UpdatePlantRequest } from '../../client/garden/v1/plant_pb';
import { PlantService } from '../../client/garden/v1/plant_connectweb';

const client = createPromiseClient(PlantService, createConnectTransport({ baseUrl: '' }));

export const getPlantRequest = async (id: string) => {
    return await client.getPlant({id});
};

export const getAllPlantsRequest = async () => {
    return await client.getPlants({});
};

export const createPlantRequest = async (plant: Partial<CreatePlantRequest>) => {
    return await client.createPlant(plant);
};

export const updatePlantRequest = async (plant: Partial<UpdatePlantRequest>) => {
    return await client.updatePlant(plant);
};

export const deletePlantRequest = async (id: string) => {
    return await client.deletePlant({id})
};