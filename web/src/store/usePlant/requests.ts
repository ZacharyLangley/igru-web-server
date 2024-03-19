import {
    createPromiseClient,
    createConnectTransport,
} from '@bufbuild/connect-web';

import { CreatePlantRequest, UpdatePlantRequest } from '../../client/garden/v1/plant_pb';
import { PlantService } from '../../client/garden/v1/plant_connectweb';
import { getCredentials } from '../utils/auth';

const client = createPromiseClient(PlantService, createConnectTransport({ baseUrl: '' }));

export const getPlantRequest = async (id: string) => {
    const options = await getCredentials();
    if (!options) return;

    return await client.getPlant({id}, options);
};

export const getAllPlantsRequest = async (groupId?: string) => {
    const options = await getCredentials();
    if (!options) return;

    return await client.getPlants({ groupId }, options);
};

export const createPlantRequest = async (plant: Partial<CreatePlantRequest>) => {
    const options = await getCredentials();
    if (!options) return;

    return await client.createPlant(plant, options);
};

export const updatePlantRequest = async (plant: Partial<UpdatePlantRequest>) => {
    const options = await getCredentials();
    if (!options) return;

    return await client.updatePlant(plant, options);
};

export const deletePlantRequest = async (id: string) => {
    const options = await getCredentials();
    if (!options) return;

    return await client.deletePlant({id}, options)
};