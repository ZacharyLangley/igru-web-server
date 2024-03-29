// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file garden/v1/recipe.proto (package garden.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Recipe } from "./schema_pb";

/**
 * @generated from message garden.v1.CreateRecipeRequest
 */
export class CreateRecipeRequest extends Message<CreateRecipeRequest> {
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
   * @generated from field: string ingredients = 4;
   */
  ingredients = "";

  /**
   * @generated from field: string instructions = 5;
   */
  instructions = "";

  /**
   * @generated from field: double ph = 6;
   */
  ph = 0;

  /**
   * @generated from field: double mix_time_milliseconds = 7;
   */
  mixTimeMilliseconds = 0;

  /**
   * @generated from field: string tags = 8;
   */
  tags = "";

  constructor(data?: PartialMessage<CreateRecipeRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.CreateRecipeRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "group_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "comment", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "ingredients", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "instructions", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "ph", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 7, name: "mix_time_milliseconds", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 8, name: "tags", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRecipeRequest {
    return new CreateRecipeRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRecipeRequest {
    return new CreateRecipeRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRecipeRequest {
    return new CreateRecipeRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateRecipeRequest | PlainMessage<CreateRecipeRequest> | undefined, b: CreateRecipeRequest | PlainMessage<CreateRecipeRequest> | undefined): boolean {
    return proto3.util.equals(CreateRecipeRequest, a, b);
  }
}

/**
 * @generated from message garden.v1.CreateRecipeResponse
 */
export class CreateRecipeResponse extends Message<CreateRecipeResponse> {
  /**
   * @generated from field: garden.v1.Recipe recipe = 1;
   */
  recipe?: Recipe;

  constructor(data?: PartialMessage<CreateRecipeResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.CreateRecipeResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "recipe", kind: "message", T: Recipe },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRecipeResponse {
    return new CreateRecipeResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRecipeResponse {
    return new CreateRecipeResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRecipeResponse {
    return new CreateRecipeResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateRecipeResponse | PlainMessage<CreateRecipeResponse> | undefined, b: CreateRecipeResponse | PlainMessage<CreateRecipeResponse> | undefined): boolean {
    return proto3.util.equals(CreateRecipeResponse, a, b);
  }
}

/**
 * @generated from message garden.v1.DeleteRecipeRequest
 */
export class DeleteRecipeRequest extends Message<DeleteRecipeRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string group_id = 2;
   */
  groupId = "";

  constructor(data?: PartialMessage<DeleteRecipeRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.DeleteRecipeRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "group_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteRecipeRequest {
    return new DeleteRecipeRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteRecipeRequest {
    return new DeleteRecipeRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteRecipeRequest {
    return new DeleteRecipeRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteRecipeRequest | PlainMessage<DeleteRecipeRequest> | undefined, b: DeleteRecipeRequest | PlainMessage<DeleteRecipeRequest> | undefined): boolean {
    return proto3.util.equals(DeleteRecipeRequest, a, b);
  }
}

/**
 * @generated from message garden.v1.DeleteRecipeResponse
 */
export class DeleteRecipeResponse extends Message<DeleteRecipeResponse> {
  /**
   * @generated from field: garden.v1.Recipe recipe = 1;
   */
  recipe?: Recipe;

  constructor(data?: PartialMessage<DeleteRecipeResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.DeleteRecipeResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "recipe", kind: "message", T: Recipe },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteRecipeResponse {
    return new DeleteRecipeResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteRecipeResponse {
    return new DeleteRecipeResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteRecipeResponse {
    return new DeleteRecipeResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteRecipeResponse | PlainMessage<DeleteRecipeResponse> | undefined, b: DeleteRecipeResponse | PlainMessage<DeleteRecipeResponse> | undefined): boolean {
    return proto3.util.equals(DeleteRecipeResponse, a, b);
  }
}

/**
 * @generated from message garden.v1.UpdateRecipeRequest
 */
export class UpdateRecipeRequest extends Message<UpdateRecipeRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string group_id = 2;
   */
  groupId = "";

  /**
   * @generated from field: string name = 3;
   */
  name = "";

  /**
   * @generated from field: string comment = 4;
   */
  comment = "";

  /**
   * @generated from field: string ingredients = 5;
   */
  ingredients = "";

  /**
   * @generated from field: string instructions = 6;
   */
  instructions = "";

  /**
   * @generated from field: double ph = 7;
   */
  ph = 0;

  /**
   * @generated from field: double mix_time_milliseconds = 8;
   */
  mixTimeMilliseconds = 0;

  /**
   * @generated from field: string tags = 9;
   */
  tags = "";

  constructor(data?: PartialMessage<UpdateRecipeRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.UpdateRecipeRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "group_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "comment", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "ingredients", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "instructions", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "ph", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 8, name: "mix_time_milliseconds", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 9, name: "tags", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateRecipeRequest {
    return new UpdateRecipeRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateRecipeRequest {
    return new UpdateRecipeRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateRecipeRequest {
    return new UpdateRecipeRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateRecipeRequest | PlainMessage<UpdateRecipeRequest> | undefined, b: UpdateRecipeRequest | PlainMessage<UpdateRecipeRequest> | undefined): boolean {
    return proto3.util.equals(UpdateRecipeRequest, a, b);
  }
}

/**
 * @generated from message garden.v1.UpdateRecipeResponse
 */
export class UpdateRecipeResponse extends Message<UpdateRecipeResponse> {
  /**
   * @generated from field: garden.v1.Recipe recipe = 1;
   */
  recipe?: Recipe;

  constructor(data?: PartialMessage<UpdateRecipeResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.UpdateRecipeResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "recipe", kind: "message", T: Recipe },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateRecipeResponse {
    return new UpdateRecipeResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateRecipeResponse {
    return new UpdateRecipeResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateRecipeResponse {
    return new UpdateRecipeResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateRecipeResponse | PlainMessage<UpdateRecipeResponse> | undefined, b: UpdateRecipeResponse | PlainMessage<UpdateRecipeResponse> | undefined): boolean {
    return proto3.util.equals(UpdateRecipeResponse, a, b);
  }
}

/**
 * @generated from message garden.v1.GetRecipesRequest
 */
export class GetRecipesRequest extends Message<GetRecipesRequest> {
  /**
   * @generated from field: string group_id = 1;
   */
  groupId = "";

  constructor(data?: PartialMessage<GetRecipesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.GetRecipesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "group_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRecipesRequest {
    return new GetRecipesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRecipesRequest {
    return new GetRecipesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRecipesRequest {
    return new GetRecipesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetRecipesRequest | PlainMessage<GetRecipesRequest> | undefined, b: GetRecipesRequest | PlainMessage<GetRecipesRequest> | undefined): boolean {
    return proto3.util.equals(GetRecipesRequest, a, b);
  }
}

/**
 * @generated from message garden.v1.GetRecipesResponse
 */
export class GetRecipesResponse extends Message<GetRecipesResponse> {
  /**
   * @generated from field: repeated garden.v1.Recipe recipes = 1;
   */
  recipes: Recipe[] = [];

  constructor(data?: PartialMessage<GetRecipesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.GetRecipesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "recipes", kind: "message", T: Recipe, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRecipesResponse {
    return new GetRecipesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRecipesResponse {
    return new GetRecipesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRecipesResponse {
    return new GetRecipesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetRecipesResponse | PlainMessage<GetRecipesResponse> | undefined, b: GetRecipesResponse | PlainMessage<GetRecipesResponse> | undefined): boolean {
    return proto3.util.equals(GetRecipesResponse, a, b);
  }
}

/**
 * @generated from message garden.v1.GetRecipeRequest
 */
export class GetRecipeRequest extends Message<GetRecipeRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string group_id = 2;
   */
  groupId = "";

  constructor(data?: PartialMessage<GetRecipeRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.GetRecipeRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "group_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRecipeRequest {
    return new GetRecipeRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRecipeRequest {
    return new GetRecipeRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRecipeRequest {
    return new GetRecipeRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetRecipeRequest | PlainMessage<GetRecipeRequest> | undefined, b: GetRecipeRequest | PlainMessage<GetRecipeRequest> | undefined): boolean {
    return proto3.util.equals(GetRecipeRequest, a, b);
  }
}

/**
 * @generated from message garden.v1.GetRecipeResponse
 */
export class GetRecipeResponse extends Message<GetRecipeResponse> {
  /**
   * @generated from field: garden.v1.Recipe recipe = 1;
   */
  recipe?: Recipe;

  constructor(data?: PartialMessage<GetRecipeResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "garden.v1.GetRecipeResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "recipe", kind: "message", T: Recipe },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRecipeResponse {
    return new GetRecipeResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRecipeResponse {
    return new GetRecipeResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRecipeResponse {
    return new GetRecipeResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetRecipeResponse | PlainMessage<GetRecipeResponse> | undefined, b: GetRecipeResponse | PlainMessage<GetRecipeResponse> | undefined): boolean {
    return proto3.util.equals(GetRecipeResponse, a, b);
  }
}

