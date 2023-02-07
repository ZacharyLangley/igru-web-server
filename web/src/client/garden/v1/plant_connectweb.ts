// @generated by protoc-gen-connect-web v0.6.0 with parameter "target=ts"
// @generated from file garden/v1/plant.proto (package garden.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreatePlantRequest, CreatePlantResponse, DeletePlantRequest, DeletePlantResponse, GetPlantRequest, GetPlantResponse, GetPlantsRequest, GetPlantsResponse, UpdatePlantRequest, UpdatePlantResponse } from "./plant_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service garden.v1.PlantService
 */
export const PlantService = {
  typeName: "garden.v1.PlantService",
  methods: {
    /**
     * @generated from rpc garden.v1.PlantService.CreatePlant
     */
    createPlant: {
      name: "CreatePlant",
      I: CreatePlantRequest,
      O: CreatePlantResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc garden.v1.PlantService.DeletePlant
     */
    deletePlant: {
      name: "DeletePlant",
      I: DeletePlantRequest,
      O: DeletePlantResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc garden.v1.PlantService.UpdatePlant
     */
    updatePlant: {
      name: "UpdatePlant",
      I: UpdatePlantRequest,
      O: UpdatePlantResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc garden.v1.PlantService.GetPlants
     */
    getPlants: {
      name: "GetPlants",
      I: GetPlantsRequest,
      O: GetPlantsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc garden.v1.PlantService.GetPlant
     */
    getPlant: {
      name: "GetPlant",
      I: GetPlantRequest,
      O: GetPlantResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
