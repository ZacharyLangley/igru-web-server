import {
    createPromiseClient,
    createConnectTransport,
    CallOptions
} from '@bufbuild/connect-web';

import { CreateGardenRequest, UpdateGardenRequest } from '../../client/garden/v1/garden_pb';
import { GardenService } from '../../client/garden/v1/garden_connectweb';
import { getUserCookie } from '../utils/cookies';

const client = createPromiseClient(GardenService, createConnectTransport({ baseUrl: '' }));

export const getGardenRequest = async (id: string) => {
    return await client.getGarden({id});
};

export const getAllGardensRequest = async (groupId?: string) => {
    const token = await getUserCookie();
    if (!token) return;

    const options: CallOptions = {headers: {session: token}};

    return await client.getGardens({groupId}, options);
};

export const createGardenRequest = async (garden: Partial<CreateGardenRequest>) => {
    const token = await getUserCookie();
    if (!token) return;

    const options: CallOptions = {headers: {session: token}}
    return await client.createGarden(garden, options);
};

export const updateGardenRequest = async (garden: Partial<UpdateGardenRequest>) => {
    return await client.updateGarden(garden);
};

export const deleteGardenRequest = async (id: string) => {
    return await client.deleteGarden({id});
};