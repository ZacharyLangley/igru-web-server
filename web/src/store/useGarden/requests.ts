import {
    createPromiseClient,
    createConnectTransport,
} from '@bufbuild/connect-web';

import { CreateGardenRequest, UpdateGardenRequest } from '../../client/garden/v1/garden_pb';
import { GardenService } from '../../client/garden/v1/garden_connectweb';
import { getCredentials } from '../utils/auth';

const client = createPromiseClient(GardenService, createConnectTransport({ baseUrl: '' }));

export const getGardenRequest = async (id: string) => {
    const options = await getCredentials();
    if (!options) return;

    return await client.getGarden({id}, options);
};

export const getAllGardensRequest = async (groupId?: string) => {
    const options = await getCredentials();
    if (!options) return;

    return await client.getGardens({groupId}, options);
};

export const createGardenRequest = async (garden: Partial<CreateGardenRequest>) => {
    const options = await getCredentials();
    if (!options) return;

    return await client.createGarden(garden, options);
};

export const updateGardenRequest = async (garden: Partial<UpdateGardenRequest>) => {
    const options = await getCredentials();
    if (!options) return

    return await client.updateGarden(garden, options);
};

export const deleteGardenRequest = async (id?: string, groupId?: string) => {
    const options = await getCredentials();
    if (!options) return;

    return await client.deleteGarden({id, groupId}, options);
};