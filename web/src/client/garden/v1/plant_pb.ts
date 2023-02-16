// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file garden/v1/plant.proto (package garden.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { Plant } from "./schema_pb.js";

/**
 * @generated from message garden.v1.CreatePlantRequest
 */
export class CreatePlantRequest extends Message<CreatePlantRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: string group_id = 2;
   */
  groupId = "";

  /**
   * @generated from field: string comment = 3;
   */
  comment = "";

  /**
   * @generated from field: string notes = 4;
   */
  notes = "";

  /**
   * @generated from field: string grow_cycle_length = 5;
   */
  growCycleLength = "";

  /**
   * @generated from field: string parentage = 6;
   */
  parentage = "";

  /**
   * @generated from field: string origin = 7;
   */
  origin = "";

  /**
   * @generated from field: string gender = 8;
   */
  gender = "";

  /**
   * @generated from field: double days_flowering = 9;
   */
  daysFlowering = 0;

  /**
   * @generated from field: double days_cured = 10;
   */
  daysCured = 0;

  /**
   * @generated from field: string harvested_weight = 11;
   */
  harvestedWeight = "";

  /**
   * @generated from field: double bud_density = 12;
   */
  budDensity = 0;

  /**
   * @generated from field: bool bud_pistils = 13;
   */
  budPistils = false;

  /**
   * @generated from field: string tags = 14;
   */
  tags = "";

  /**
   * @generated from field: google.protobuf.Timestamp acquired_at = 15;
   */
  acquiredAt?: Timestamp;

  constructor(data?: PartialMessage<CreatePlantRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.CreatePlantRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "group_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "comment", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "notes", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "grow_cycle_length", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "parentage", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "origin", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "gender", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "days_flowering", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 10, name: "days_cured", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 11, name: "harvested_weight", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "bud_density", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 13, name: "bud_pistils", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 14, name: "tags", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 15, name: "acquired_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreatePlantRequest {
    return new CreatePlantRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreatePlantRequest {
    return new CreatePlantRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreatePlantRequest {
    return new CreatePlantRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreatePlantRequest | PlainMessage<CreatePlantRequest> | undefined, b: CreatePlantRequest | PlainMessage<CreatePlantRequest> | undefined): boolean {
    return proto3.util.equals(CreatePlantRequest, a, b);
  }
}

/**
 * @generated from message garden.v1.CreatePlantResponse
 */
export class CreatePlantResponse extends Message<CreatePlantResponse> {
  /**
   * @generated from field: garden.v1.Plant plant = 1;
   */
  plant?: Plant;

  constructor(data?: PartialMessage<CreatePlantResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.CreatePlantResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "plant", kind: "message", T: Plant },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreatePlantResponse {
    return new CreatePlantResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreatePlantResponse {
    return new CreatePlantResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreatePlantResponse {
    return new CreatePlantResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreatePlantResponse | PlainMessage<CreatePlantResponse> | undefined, b: CreatePlantResponse | PlainMessage<CreatePlantResponse> | undefined): boolean {
    return proto3.util.equals(CreatePlantResponse, a, b);
  }
}

/**
 * @generated from message garden.v1.DeletePlantRequest
 */
export class DeletePlantRequest extends Message<DeletePlantRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<DeletePlantRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.DeletePlantRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeletePlantRequest {
    return new DeletePlantRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeletePlantRequest {
    return new DeletePlantRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeletePlantRequest {
    return new DeletePlantRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeletePlantRequest | PlainMessage<DeletePlantRequest> | undefined, b: DeletePlantRequest | PlainMessage<DeletePlantRequest> | undefined): boolean {
    return proto3.util.equals(DeletePlantRequest, a, b);
  }
}

/**
 * @generated from message garden.v1.DeletePlantResponse
 */
export class DeletePlantResponse extends Message<DeletePlantResponse> {
  /**
   * @generated from field: garden.v1.Plant plant = 1;
   */
  plant?: Plant;

  constructor(data?: PartialMessage<DeletePlantResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.DeletePlantResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "plant", kind: "message", T: Plant },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeletePlantResponse {
    return new DeletePlantResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeletePlantResponse {
    return new DeletePlantResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeletePlantResponse {
    return new DeletePlantResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeletePlantResponse | PlainMessage<DeletePlantResponse> | undefined, b: DeletePlantResponse | PlainMessage<DeletePlantResponse> | undefined): boolean {
    return proto3.util.equals(DeletePlantResponse, a, b);
  }
}

/**
 * @generated from message garden.v1.UpdatePlantRequest
 */
export class UpdatePlantRequest extends Message<UpdatePlantRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string comment = 3;
   */
  comment = "";

  /**
   * @generated from field: string notes = 4;
   */
  notes = "";

  /**
   * @generated from field: string grow_cycle_length = 5;
   */
  growCycleLength = "";

  /**
   * @generated from field: string parentage = 6;
   */
  parentage = "";

  /**
   * @generated from field: string origin = 7;
   */
  origin = "";

  /**
   * @generated from field: string gender = 8;
   */
  gender = "";

  /**
   * @generated from field: double days_flowering = 9;
   */
  daysFlowering = 0;

  /**
   * @generated from field: double days_cured = 10;
   */
  daysCured = 0;

  /**
   * @generated from field: string harvested_weight = 11;
   */
  harvestedWeight = "";

  /**
   * @generated from field: double bud_density = 12;
   */
  budDensity = 0;

  /**
   * @generated from field: bool bud_pistils = 13;
   */
  budPistils = false;

  /**
   * @generated from field: string tags = 14;
   */
  tags = "";

  /**
   * @generated from field: google.protobuf.Timestamp acquired_at = 15;
   */
  acquiredAt?: Timestamp;

  constructor(data?: PartialMessage<UpdatePlantRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.UpdatePlantRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "comment", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "notes", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "grow_cycle_length", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "parentage", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "origin", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "gender", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "days_flowering", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 10, name: "days_cured", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 11, name: "harvested_weight", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "bud_density", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 13, name: "bud_pistils", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 14, name: "tags", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 15, name: "acquired_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePlantRequest {
    return new UpdatePlantRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePlantRequest {
    return new UpdatePlantRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePlantRequest {
    return new UpdatePlantRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePlantRequest | PlainMessage<UpdatePlantRequest> | undefined, b: UpdatePlantRequest | PlainMessage<UpdatePlantRequest> | undefined): boolean {
    return proto3.util.equals(UpdatePlantRequest, a, b);
  }
}

/**
 * @generated from message garden.v1.UpdatePlantResponse
 */
export class UpdatePlantResponse extends Message<UpdatePlantResponse> {
  /**
   * @generated from field: garden.v1.Plant plant = 1;
   */
  plant?: Plant;

  constructor(data?: PartialMessage<UpdatePlantResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.UpdatePlantResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "plant", kind: "message", T: Plant },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePlantResponse {
    return new UpdatePlantResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePlantResponse {
    return new UpdatePlantResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePlantResponse {
    return new UpdatePlantResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePlantResponse | PlainMessage<UpdatePlantResponse> | undefined, b: UpdatePlantResponse | PlainMessage<UpdatePlantResponse> | undefined): boolean {
    return proto3.util.equals(UpdatePlantResponse, a, b);
  }
}

/**
 * @generated from message garden.v1.GetPlantsRequest
 */
export class GetPlantsRequest extends Message<GetPlantsRequest> {
  /**
   * @generated from field: string group_id = 1;
   */
  groupId = "";

  constructor(data?: PartialMessage<GetPlantsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.GetPlantsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "group_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPlantsRequest {
    return new GetPlantsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPlantsRequest {
    return new GetPlantsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPlantsRequest {
    return new GetPlantsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetPlantsRequest | PlainMessage<GetPlantsRequest> | undefined, b: GetPlantsRequest | PlainMessage<GetPlantsRequest> | undefined): boolean {
    return proto3.util.equals(GetPlantsRequest, a, b);
  }
}

/**
 * @generated from message garden.v1.GetPlantsResponse
 */
export class GetPlantsResponse extends Message<GetPlantsResponse> {
  /**
   * @generated from field: repeated garden.v1.Plant plants = 1;
   */
  plants: Plant[] = [];

  constructor(data?: PartialMessage<GetPlantsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.GetPlantsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "plants", kind: "message", T: Plant, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPlantsResponse {
    return new GetPlantsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPlantsResponse {
    return new GetPlantsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPlantsResponse {
    return new GetPlantsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetPlantsResponse | PlainMessage<GetPlantsResponse> | undefined, b: GetPlantsResponse | PlainMessage<GetPlantsResponse> | undefined): boolean {
    return proto3.util.equals(GetPlantsResponse, a, b);
  }
}

/**
 * @generated from message garden.v1.GetPlantRequest
 */
export class GetPlantRequest extends Message<GetPlantRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<GetPlantRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.GetPlantRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPlantRequest {
    return new GetPlantRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPlantRequest {
    return new GetPlantRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPlantRequest {
    return new GetPlantRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetPlantRequest | PlainMessage<GetPlantRequest> | undefined, b: GetPlantRequest | PlainMessage<GetPlantRequest> | undefined): boolean {
    return proto3.util.equals(GetPlantRequest, a, b);
  }
}

/**
 * @generated from message garden.v1.GetPlantResponse
 */
export class GetPlantResponse extends Message<GetPlantResponse> {
  /**
   * @generated from field: garden.v1.Plant plant = 1;
   */
  plant?: Plant;

  constructor(data?: PartialMessage<GetPlantResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.GetPlantResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "plant", kind: "message", T: Plant },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPlantResponse {
    return new GetPlantResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPlantResponse {
    return new GetPlantResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPlantResponse {
    return new GetPlantResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetPlantResponse | PlainMessage<GetPlantResponse> | undefined, b: GetPlantResponse | PlainMessage<GetPlantResponse> | undefined): boolean {
    return proto3.util.equals(GetPlantResponse, a, b);
  }
}

