// @generated by protoc-gen-connect-web v0.6.0 with parameter "target=ts"
// @generated from file gardens/v1/strains.proto (package gardens.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateStrainRequest, CreateStrainResponse, DeleteStrainRequest, DeleteStrainResponse, GetStrainRequest, GetStrainResponse, GetStrainsRequest, GetStrainsResponse, UpdateStrainRequest, UpdateStrainResponse } from "./strains_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service gardens.v1.StrainsService
 */
export const StrainsService = {
  typeName: "gardens.v1.StrainsService",
  methods: {
    /**
     * @generated from rpc gardens.v1.StrainsService.CreateStrain
     */
    createStrain: {
      name: "CreateStrain",
      I: CreateStrainRequest,
      O: CreateStrainResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc gardens.v1.StrainsService.DeleteStrain
     */
    deleteStrain: {
      name: "DeleteStrain",
      I: DeleteStrainRequest,
      O: DeleteStrainResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc gardens.v1.StrainsService.UpdateStrain
     */
    updateStrain: {
      name: "UpdateStrain",
      I: UpdateStrainRequest,
      O: UpdateStrainResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc gardens.v1.StrainsService.GetStrains
     */
    getStrains: {
      name: "GetStrains",
      I: GetStrainsRequest,
      O: GetStrainsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc gardens.v1.StrainsService.GetStrain
     */
    getStrain: {
      name: "GetStrain",
      I: GetStrainRequest,
      O: GetStrainResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

