import {
    createPromiseClient,
    createConnectTransport,
} from '@bufbuild/connect-web';

import { CreateGardenRequest, UpdateGardenRequest } from '../../client/garden/v1/garden_pb';
import { GardenService } from '../../client/garden/v1/garden_connectweb';

const client = createPromiseClient(
    GardenService,
    createConnectTransport({
        baseUrl: '',
    })
);

export const getGardenRequest = async (id: string) => {
    return await client.getGarden({id});
};

export const getAllGardensRequest = async () => {
    return await client.getGardens({});
};

export const createGardenRequest = async (garden: Partial<CreateGardenRequest>) => {
    return await client.createGarden(garden);
};

export const updateGardenRequest = async (garden: Partial<UpdateGardenRequest>) => {
    return await client.updateGarden(garden);
};

export const deleteGardenRequest = async (id: string) => {
    return await client.deleteGarden({id});
};