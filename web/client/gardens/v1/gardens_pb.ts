// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file gardens/v1/gardens.proto (package gardens.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Garden } from "./schema_pb.js";

/**
 * @generated from message gardens.v1.CreateGardenRequest
 */
export class CreateGardenRequest extends Message<CreateGardenRequest> {
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
   * @generated from field: string location = 4;
   */
  location = "";

  /**
   * @generated from field: string grow_type = 5;
   */
  growType = "";

  /**
   * @generated from field: string grow_size = 6;
   */
  growSize = "";

  /**
   * @generated from field: string grow_style = 7;
   */
  growStyle = "";

  /**
   * @generated from field: string container_size = 8;
   */
  containerSize = "";

  /**
   * @generated from field: string tags = 9;
   */
  tags = "";

  constructor(data?: PartialMessage<CreateGardenRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "gardens.v1.CreateGardenRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "group_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "comment", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "location", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "grow_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "grow_size", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "grow_style", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "container_size", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "tags", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateGardenRequest {
    return new CreateGardenRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateGardenRequest {
    return new CreateGardenRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateGardenRequest {
    return new CreateGardenRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateGardenRequest | PlainMessage<CreateGardenRequest> | undefined, b: CreateGardenRequest | PlainMessage<CreateGardenRequest> | undefined): boolean {
    return proto3.util.equals(CreateGardenRequest, a, b);
  }
}

/**
 * @generated from message gardens.v1.CreateGardenResponse
 */
export class CreateGardenResponse extends Message<CreateGardenResponse> {
  /**
   * @generated from field: gardens.v1.Garden garden = 1;
   */
  garden?: Garden;

  constructor(data?: PartialMessage<CreateGardenResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "gardens.v1.CreateGardenResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "garden", kind: "message", T: Garden },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateGardenResponse {
    return new CreateGardenResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateGardenResponse {
    return new CreateGardenResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateGardenResponse {
    return new CreateGardenResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateGardenResponse | PlainMessage<CreateGardenResponse> | undefined, b: CreateGardenResponse | PlainMessage<CreateGardenResponse> | undefined): boolean {
    return proto3.util.equals(CreateGardenResponse, a, b);
  }
}

/**
 * @generated from message gardens.v1.DeleteGardenRequest
 */
export class DeleteGardenRequest extends Message<DeleteGardenRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<DeleteGardenRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "gardens.v1.DeleteGardenRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteGardenRequest {
    return new DeleteGardenRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteGardenRequest {
    return new DeleteGardenRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteGardenRequest {
    return new DeleteGardenRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteGardenRequest | PlainMessage<DeleteGardenRequest> | undefined, b: DeleteGardenRequest | PlainMessage<DeleteGardenRequest> | undefined): boolean {
    return proto3.util.equals(DeleteGardenRequest, a, b);
  }
}

/**
 * @generated from message gardens.v1.DeleteGardenResponse
 */
export class DeleteGardenResponse extends Message<DeleteGardenResponse> {
  /**
   * @generated from field: gardens.v1.Garden garden = 1;
   */
  garden?: Garden;

  constructor(data?: PartialMessage<DeleteGardenResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "gardens.v1.DeleteGardenResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "garden", kind: "message", T: Garden },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteGardenResponse {
    return new DeleteGardenResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteGardenResponse {
    return new DeleteGardenResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteGardenResponse {
    return new DeleteGardenResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteGardenResponse | PlainMessage<DeleteGardenResponse> | undefined, b: DeleteGardenResponse | PlainMessage<DeleteGardenResponse> | undefined): boolean {
    return proto3.util.equals(DeleteGardenResponse, a, b);
  }
}

/**
 * @generated from message gardens.v1.UpdateGardenRequest
 */
export class UpdateGardenRequest extends Message<UpdateGardenRequest> {
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
   * @generated from field: string location = 4;
   */
  location = "";

  /**
   * @generated from field: string grow_type = 5;
   */
  growType = "";

  /**
   * @generated from field: string grow_size = 6;
   */
  growSize = "";

  /**
   * @generated from field: string grow_style = 7;
   */
  growStyle = "";

  /**
   * @generated from field: string container_size = 8;
   */
  containerSize = "";

  /**
   * @generated from field: string tags = 9;
   */
  tags = "";

  constructor(data?: PartialMessage<UpdateGardenRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "gardens.v1.UpdateGardenRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "comment", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "location", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "grow_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "grow_size", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "grow_style", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "container_size", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "tags", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateGardenRequest {
    return new UpdateGardenRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateGardenRequest {
    return new UpdateGardenRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateGardenRequest {
    return new UpdateGardenRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateGardenRequest | PlainMessage<UpdateGardenRequest> | undefined, b: UpdateGardenRequest | PlainMessage<UpdateGardenRequest> | undefined): boolean {
    return proto3.util.equals(UpdateGardenRequest, a, b);
  }
}

/**
 * @generated from message gardens.v1.UpdateGardenResponse
 */
export class UpdateGardenResponse extends Message<UpdateGardenResponse> {
  /**
   * @generated from field: gardens.v1.Garden garden = 1;
   */
  garden?: Garden;

  constructor(data?: PartialMessage<UpdateGardenResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "gardens.v1.UpdateGardenResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "garden", kind: "message", T: Garden },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateGardenResponse {
    return new UpdateGardenResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateGardenResponse {
    return new UpdateGardenResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateGardenResponse {
    return new UpdateGardenResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateGardenResponse | PlainMessage<UpdateGardenResponse> | undefined, b: UpdateGardenResponse | PlainMessage<UpdateGardenResponse> | undefined): boolean {
    return proto3.util.equals(UpdateGardenResponse, a, b);
  }
}

/**
 * @generated from message gardens.v1.GetGardensRequest
 */
export class GetGardensRequest extends Message<GetGardensRequest> {
  /**
   * @generated from field: string group_id = 1;
   */
  groupId = "";

  constructor(data?: PartialMessage<GetGardensRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "gardens.v1.GetGardensRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "group_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetGardensRequest {
    return new GetGardensRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetGardensRequest {
    return new GetGardensRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetGardensRequest {
    return new GetGardensRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetGardensRequest | PlainMessage<GetGardensRequest> | undefined, b: GetGardensRequest | PlainMessage<GetGardensRequest> | undefined): boolean {
    return proto3.util.equals(GetGardensRequest, a, b);
  }
}

/**
 * @generated from message gardens.v1.GetGardensResponse
 */
export class GetGardensResponse extends Message<GetGardensResponse> {
  /**
   * @generated from field: repeated gardens.v1.Garden gardens = 1;
   */
  gardens: Garden[] = [];

  constructor(data?: PartialMessage<GetGardensResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "gardens.v1.GetGardensResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "gardens", kind: "message", T: Garden, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetGardensResponse {
    return new GetGardensResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetGardensResponse {
    return new GetGardensResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetGardensResponse {
    return new GetGardensResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetGardensResponse | PlainMessage<GetGardensResponse> | undefined, b: GetGardensResponse | PlainMessage<GetGardensResponse> | undefined): boolean {
    return proto3.util.equals(GetGardensResponse, a, b);
  }
}

/**
 * @generated from message gardens.v1.GetGardenRequest
 */
export class GetGardenRequest extends Message<GetGardenRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<GetGardenRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "gardens.v1.GetGardenRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetGardenRequest {
    return new GetGardenRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetGardenRequest {
    return new GetGardenRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetGardenRequest {
    return new GetGardenRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetGardenRequest | PlainMessage<GetGardenRequest> | undefined, b: GetGardenRequest | PlainMessage<GetGardenRequest> | undefined): boolean {
    return proto3.util.equals(GetGardenRequest, a, b);
  }
}

/**
 * @generated from message gardens.v1.GetGardenResponse
 */
export class GetGardenResponse extends Message<GetGardenResponse> {
  /**
   * @generated from field: gardens.v1.Garden garden = 1;
   */
  garden?: Garden;

  constructor(data?: PartialMessage<GetGardenResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "gardens.v1.GetGardenResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "garden", kind: "message", T: Garden },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetGardenResponse {
    return new GetGardenResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetGardenResponse {
    return new GetGardenResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetGardenResponse {
    return new GetGardenResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetGardenResponse | PlainMessage<GetGardenResponse> | undefined, b: GetGardenResponse | PlainMessage<GetGardenResponse> | undefined): boolean {
    return proto3.util.equals(GetGardenResponse, a, b);
  }
}

