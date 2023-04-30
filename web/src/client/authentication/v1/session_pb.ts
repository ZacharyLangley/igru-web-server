// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file authentication/v1/session.proto (package authentication.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { PermissionRequest, PermissionResponse, Session, User } from "./schema_pb";
import { PaginationRequest } from "../../common/v1/pagination_pb";

/**
 * @generated from message authentication.v1.CreateSessionRequest
 */
export class CreateSessionRequest extends Message<CreateSessionRequest> {
  /**
   * @generated from field: string email = 1;
   */
  email = "";

  /**
   * @generated from field: string password = 2;
   */
  password = "";

  /**
   * @generated from field: string fullName = 3;
   */
  fullName = "";

  constructor(data?: PartialMessage<CreateSessionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.CreateSessionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "fullName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSessionRequest {
    return new CreateSessionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSessionRequest {
    return new CreateSessionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSessionRequest {
    return new CreateSessionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSessionRequest | PlainMessage<CreateSessionRequest> | undefined, b: CreateSessionRequest | PlainMessage<CreateSessionRequest> | undefined): boolean {
    return proto3.util.equals(CreateSessionRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.CreateSessionResponse
 */
export class CreateSessionResponse extends Message<CreateSessionResponse> {
  /**
   * @generated from field: authentication.v1.User user = 1;
   */
  user?: User;

  constructor(data?: PartialMessage<CreateSessionResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.CreateSessionResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user", kind: "message", T: User },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSessionResponse {
    return new CreateSessionResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSessionResponse {
    return new CreateSessionResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSessionResponse {
    return new CreateSessionResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSessionResponse | PlainMessage<CreateSessionResponse> | undefined, b: CreateSessionResponse | PlainMessage<CreateSessionResponse> | undefined): boolean {
    return proto3.util.equals(CreateSessionResponse, a, b);
  }
}

/**
 * @generated from message authentication.v1.ValidateSessionRequest
 */
export class ValidateSessionRequest extends Message<ValidateSessionRequest> {
  constructor(data?: PartialMessage<ValidateSessionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.ValidateSessionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ValidateSessionRequest {
    return new ValidateSessionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ValidateSessionRequest {
    return new ValidateSessionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ValidateSessionRequest {
    return new ValidateSessionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ValidateSessionRequest | PlainMessage<ValidateSessionRequest> | undefined, b: ValidateSessionRequest | PlainMessage<ValidateSessionRequest> | undefined): boolean {
    return proto3.util.equals(ValidateSessionRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.ValidateSessionResponse
 */
export class ValidateSessionResponse extends Message<ValidateSessionResponse> {
  /**
   * @generated from field: authentication.v1.User user = 1;
   */
  user?: User;

  constructor(data?: PartialMessage<ValidateSessionResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.ValidateSessionResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user", kind: "message", T: User },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ValidateSessionResponse {
    return new ValidateSessionResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ValidateSessionResponse {
    return new ValidateSessionResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ValidateSessionResponse {
    return new ValidateSessionResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ValidateSessionResponse | PlainMessage<ValidateSessionResponse> | undefined, b: ValidateSessionResponse | PlainMessage<ValidateSessionResponse> | undefined): boolean {
    return proto3.util.equals(ValidateSessionResponse, a, b);
  }
}

/**
 * @generated from message authentication.v1.GetSessionsRequest
 */
export class GetSessionsRequest extends Message<GetSessionsRequest> {
  /**
   * @generated from field: common.v1.PaginationRequest pagination = 1;
   */
  pagination?: PaginationRequest;

  constructor(data?: PartialMessage<GetSessionsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.GetSessionsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pagination", kind: "message", T: PaginationRequest },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSessionsRequest {
    return new GetSessionsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSessionsRequest {
    return new GetSessionsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSessionsRequest {
    return new GetSessionsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetSessionsRequest | PlainMessage<GetSessionsRequest> | undefined, b: GetSessionsRequest | PlainMessage<GetSessionsRequest> | undefined): boolean {
    return proto3.util.equals(GetSessionsRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.GetSessionsResponse
 */
export class GetSessionsResponse extends Message<GetSessionsResponse> {
  /**
   * @generated from field: repeated authentication.v1.Session sessions = 1;
   */
  sessions: Session[] = [];

  constructor(data?: PartialMessage<GetSessionsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.GetSessionsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sessions", kind: "message", T: Session, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSessionsResponse {
    return new GetSessionsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSessionsResponse {
    return new GetSessionsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSessionsResponse {
    return new GetSessionsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetSessionsResponse | PlainMessage<GetSessionsResponse> | undefined, b: GetSessionsResponse | PlainMessage<GetSessionsResponse> | undefined): boolean {
    return proto3.util.equals(GetSessionsResponse, a, b);
  }
}

/**
 * @generated from message authentication.v1.DeleteSessionRequest
 */
export class DeleteSessionRequest extends Message<DeleteSessionRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string user_id = 2;
   */
  userId = "";

  constructor(data?: PartialMessage<DeleteSessionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.DeleteSessionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteSessionRequest {
    return new DeleteSessionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteSessionRequest {
    return new DeleteSessionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteSessionRequest {
    return new DeleteSessionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteSessionRequest | PlainMessage<DeleteSessionRequest> | undefined, b: DeleteSessionRequest | PlainMessage<DeleteSessionRequest> | undefined): boolean {
    return proto3.util.equals(DeleteSessionRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.DeleteSessionResponse
 */
export class DeleteSessionResponse extends Message<DeleteSessionResponse> {
  constructor(data?: PartialMessage<DeleteSessionResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.DeleteSessionResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteSessionResponse {
    return new DeleteSessionResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteSessionResponse {
    return new DeleteSessionResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteSessionResponse {
    return new DeleteSessionResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteSessionResponse | PlainMessage<DeleteSessionResponse> | undefined, b: DeleteSessionResponse | PlainMessage<DeleteSessionResponse> | undefined): boolean {
    return proto3.util.equals(DeleteSessionResponse, a, b);
  }
}

/**
 * @generated from message authentication.v1.CheckSessionPermissionsRequest
 */
export class CheckSessionPermissionsRequest extends Message<CheckSessionPermissionsRequest> {
  /**
   * @generated from field: repeated authentication.v1.PermissionRequest requests = 1;
   */
  requests: PermissionRequest[] = [];

  constructor(data?: PartialMessage<CheckSessionPermissionsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.CheckSessionPermissionsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "requests", kind: "message", T: PermissionRequest, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CheckSessionPermissionsRequest {
    return new CheckSessionPermissionsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CheckSessionPermissionsRequest {
    return new CheckSessionPermissionsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CheckSessionPermissionsRequest {
    return new CheckSessionPermissionsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CheckSessionPermissionsRequest | PlainMessage<CheckSessionPermissionsRequest> | undefined, b: CheckSessionPermissionsRequest | PlainMessage<CheckSessionPermissionsRequest> | undefined): boolean {
    return proto3.util.equals(CheckSessionPermissionsRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.CheckSessionPermissionsResponse
 */
export class CheckSessionPermissionsResponse extends Message<CheckSessionPermissionsResponse> {
  /**
   * @generated from field: repeated authentication.v1.PermissionResponse responses = 2;
   */
  responses: PermissionResponse[] = [];

  constructor(data?: PartialMessage<CheckSessionPermissionsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.CheckSessionPermissionsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 2, name: "responses", kind: "message", T: PermissionResponse, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CheckSessionPermissionsResponse {
    return new CheckSessionPermissionsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CheckSessionPermissionsResponse {
    return new CheckSessionPermissionsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CheckSessionPermissionsResponse {
    return new CheckSessionPermissionsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CheckSessionPermissionsResponse | PlainMessage<CheckSessionPermissionsResponse> | undefined, b: CheckSessionPermissionsResponse | PlainMessage<CheckSessionPermissionsResponse> | undefined): boolean {
    return proto3.util.equals(CheckSessionPermissionsResponse, a, b);
  }
}

