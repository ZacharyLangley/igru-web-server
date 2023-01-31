// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file authentication/v1/groups.proto (package authentication.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Group, GroupMember, GroupRole } from "./schema_pb.js";
import { PaginationRequest } from "../../common/v1/pagination_pb.js";

/**
 * @generated from message authentication.v1.CreateGroupRequest
 */
export class CreateGroupRequest extends Message<CreateGroupRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<CreateGroupRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.CreateGroupRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateGroupRequest {
    return new CreateGroupRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateGroupRequest {
    return new CreateGroupRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateGroupRequest {
    return new CreateGroupRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateGroupRequest | PlainMessage<CreateGroupRequest> | undefined, b: CreateGroupRequest | PlainMessage<CreateGroupRequest> | undefined): boolean {
    return proto3.util.equals(CreateGroupRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.CreateGroupResponse
 */
export class CreateGroupResponse extends Message<CreateGroupResponse> {
  /**
   * @generated from field: authentication.v1.Group group = 1;
   */
  group?: Group;

  constructor(data?: PartialMessage<CreateGroupResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.CreateGroupResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "group", kind: "message", T: Group },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateGroupResponse {
    return new CreateGroupResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateGroupResponse {
    return new CreateGroupResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateGroupResponse {
    return new CreateGroupResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateGroupResponse | PlainMessage<CreateGroupResponse> | undefined, b: CreateGroupResponse | PlainMessage<CreateGroupResponse> | undefined): boolean {
    return proto3.util.equals(CreateGroupResponse, a, b);
  }
}

/**
 * @generated from message authentication.v1.UpdateGroupRequest
 */
export class UpdateGroupRequest extends Message<UpdateGroupRequest> {
  /**
   * @generated from field: authentication.v1.Group group = 1;
   */
  group?: Group;

  constructor(data?: PartialMessage<UpdateGroupRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.UpdateGroupRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "group", kind: "message", T: Group },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateGroupRequest {
    return new UpdateGroupRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateGroupRequest {
    return new UpdateGroupRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateGroupRequest {
    return new UpdateGroupRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateGroupRequest | PlainMessage<UpdateGroupRequest> | undefined, b: UpdateGroupRequest | PlainMessage<UpdateGroupRequest> | undefined): boolean {
    return proto3.util.equals(UpdateGroupRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.UpdateGroupResponse
 */
export class UpdateGroupResponse extends Message<UpdateGroupResponse> {
  /**
   * @generated from field: authentication.v1.Group group = 1;
   */
  group?: Group;

  constructor(data?: PartialMessage<UpdateGroupResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.UpdateGroupResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "group", kind: "message", T: Group },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateGroupResponse {
    return new UpdateGroupResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateGroupResponse {
    return new UpdateGroupResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateGroupResponse {
    return new UpdateGroupResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateGroupResponse | PlainMessage<UpdateGroupResponse> | undefined, b: UpdateGroupResponse | PlainMessage<UpdateGroupResponse> | undefined): boolean {
    return proto3.util.equals(UpdateGroupResponse, a, b);
  }
}

/**
 * @generated from message authentication.v1.DeleteGroupRequest
 */
export class DeleteGroupRequest extends Message<DeleteGroupRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<DeleteGroupRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.DeleteGroupRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteGroupRequest {
    return new DeleteGroupRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteGroupRequest {
    return new DeleteGroupRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteGroupRequest {
    return new DeleteGroupRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteGroupRequest | PlainMessage<DeleteGroupRequest> | undefined, b: DeleteGroupRequest | PlainMessage<DeleteGroupRequest> | undefined): boolean {
    return proto3.util.equals(DeleteGroupRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.DeleteGroupResponse
 */
export class DeleteGroupResponse extends Message<DeleteGroupResponse> {
  constructor(data?: PartialMessage<DeleteGroupResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.DeleteGroupResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteGroupResponse {
    return new DeleteGroupResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteGroupResponse {
    return new DeleteGroupResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteGroupResponse {
    return new DeleteGroupResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteGroupResponse | PlainMessage<DeleteGroupResponse> | undefined, b: DeleteGroupResponse | PlainMessage<DeleteGroupResponse> | undefined): boolean {
    return proto3.util.equals(DeleteGroupResponse, a, b);
  }
}

/**
 * @generated from message authentication.v1.GetGroupRequest
 */
export class GetGroupRequest extends Message<GetGroupRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<GetGroupRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.GetGroupRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetGroupRequest {
    return new GetGroupRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetGroupRequest {
    return new GetGroupRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetGroupRequest {
    return new GetGroupRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetGroupRequest | PlainMessage<GetGroupRequest> | undefined, b: GetGroupRequest | PlainMessage<GetGroupRequest> | undefined): boolean {
    return proto3.util.equals(GetGroupRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.GetGroupResponse
 */
export class GetGroupResponse extends Message<GetGroupResponse> {
  /**
   * @generated from field: authentication.v1.Group group = 1;
   */
  group?: Group;

  constructor(data?: PartialMessage<GetGroupResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.GetGroupResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "group", kind: "message", T: Group },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetGroupResponse {
    return new GetGroupResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetGroupResponse {
    return new GetGroupResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetGroupResponse {
    return new GetGroupResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetGroupResponse | PlainMessage<GetGroupResponse> | undefined, b: GetGroupResponse | PlainMessage<GetGroupResponse> | undefined): boolean {
    return proto3.util.equals(GetGroupResponse, a, b);
  }
}

/**
 * @generated from message authentication.v1.GetGroupsRequest
 */
export class GetGroupsRequest extends Message<GetGroupsRequest> {
  /**
   * @generated from field: common.v1.PaginationRequest pagination = 1;
   */
  pagination?: PaginationRequest;

  constructor(data?: PartialMessage<GetGroupsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.GetGroupsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pagination", kind: "message", T: PaginationRequest },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetGroupsRequest {
    return new GetGroupsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetGroupsRequest {
    return new GetGroupsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetGroupsRequest {
    return new GetGroupsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetGroupsRequest | PlainMessage<GetGroupsRequest> | undefined, b: GetGroupsRequest | PlainMessage<GetGroupsRequest> | undefined): boolean {
    return proto3.util.equals(GetGroupsRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.GetGroupsResponse
 */
export class GetGroupsResponse extends Message<GetGroupsResponse> {
  /**
   * @generated from field: repeated authentication.v1.Group groups = 1;
   */
  groups: Group[] = [];

  constructor(data?: PartialMessage<GetGroupsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.GetGroupsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "groups", kind: "message", T: Group, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetGroupsResponse {
    return new GetGroupsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetGroupsResponse {
    return new GetGroupsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetGroupsResponse {
    return new GetGroupsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetGroupsResponse | PlainMessage<GetGroupsResponse> | undefined, b: GetGroupsResponse | PlainMessage<GetGroupsResponse> | undefined): boolean {
    return proto3.util.equals(GetGroupsResponse, a, b);
  }
}

/**
 * @generated from message authentication.v1.AddGroupMemberRequest
 */
export class AddGroupMemberRequest extends Message<AddGroupMemberRequest> {
  /**
   * @generated from field: string group_id = 1;
   */
  groupId = "";

  /**
   * @generated from field: string user_id = 2;
   */
  userId = "";

  constructor(data?: PartialMessage<AddGroupMemberRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.AddGroupMemberRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "group_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddGroupMemberRequest {
    return new AddGroupMemberRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddGroupMemberRequest {
    return new AddGroupMemberRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddGroupMemberRequest {
    return new AddGroupMemberRequest().fromJsonString(jsonString, options);
  }

  static equals(a: AddGroupMemberRequest | PlainMessage<AddGroupMemberRequest> | undefined, b: AddGroupMemberRequest | PlainMessage<AddGroupMemberRequest> | undefined): boolean {
    return proto3.util.equals(AddGroupMemberRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.AddGroupMemberResponse
 */
export class AddGroupMemberResponse extends Message<AddGroupMemberResponse> {
  constructor(data?: PartialMessage<AddGroupMemberResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.AddGroupMemberResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddGroupMemberResponse {
    return new AddGroupMemberResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddGroupMemberResponse {
    return new AddGroupMemberResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddGroupMemberResponse {
    return new AddGroupMemberResponse().fromJsonString(jsonString, options);
  }

  static equals(a: AddGroupMemberResponse | PlainMessage<AddGroupMemberResponse> | undefined, b: AddGroupMemberResponse | PlainMessage<AddGroupMemberResponse> | undefined): boolean {
    return proto3.util.equals(AddGroupMemberResponse, a, b);
  }
}

/**
 * @generated from message authentication.v1.UpdateGroupMemberRequest
 */
export class UpdateGroupMemberRequest extends Message<UpdateGroupMemberRequest> {
  /**
   * @generated from field: string group_id = 1;
   */
  groupId = "";

  /**
   * @generated from field: string user_id = 2;
   */
  userId = "";

  /**
   * @generated from field: authentication.v1.GroupRole role = 3;
   */
  role = GroupRole.UNSPECIFIED;

  constructor(data?: PartialMessage<UpdateGroupMemberRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.UpdateGroupMemberRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "group_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "role", kind: "enum", T: proto3.getEnumType(GroupRole) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateGroupMemberRequest {
    return new UpdateGroupMemberRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateGroupMemberRequest {
    return new UpdateGroupMemberRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateGroupMemberRequest {
    return new UpdateGroupMemberRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateGroupMemberRequest | PlainMessage<UpdateGroupMemberRequest> | undefined, b: UpdateGroupMemberRequest | PlainMessage<UpdateGroupMemberRequest> | undefined): boolean {
    return proto3.util.equals(UpdateGroupMemberRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.UpdateGroupMemberResponse
 */
export class UpdateGroupMemberResponse extends Message<UpdateGroupMemberResponse> {
  constructor(data?: PartialMessage<UpdateGroupMemberResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.UpdateGroupMemberResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateGroupMemberResponse {
    return new UpdateGroupMemberResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateGroupMemberResponse {
    return new UpdateGroupMemberResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateGroupMemberResponse {
    return new UpdateGroupMemberResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateGroupMemberResponse | PlainMessage<UpdateGroupMemberResponse> | undefined, b: UpdateGroupMemberResponse | PlainMessage<UpdateGroupMemberResponse> | undefined): boolean {
    return proto3.util.equals(UpdateGroupMemberResponse, a, b);
  }
}

/**
 * @generated from message authentication.v1.RemoveGroupMemberRequest
 */
export class RemoveGroupMemberRequest extends Message<RemoveGroupMemberRequest> {
  /**
   * @generated from field: string group_id = 1;
   */
  groupId = "";

  /**
   * @generated from field: string user_id = 2;
   */
  userId = "";

  constructor(data?: PartialMessage<RemoveGroupMemberRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.RemoveGroupMemberRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "group_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RemoveGroupMemberRequest {
    return new RemoveGroupMemberRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RemoveGroupMemberRequest {
    return new RemoveGroupMemberRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RemoveGroupMemberRequest {
    return new RemoveGroupMemberRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RemoveGroupMemberRequest | PlainMessage<RemoveGroupMemberRequest> | undefined, b: RemoveGroupMemberRequest | PlainMessage<RemoveGroupMemberRequest> | undefined): boolean {
    return proto3.util.equals(RemoveGroupMemberRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.RemoveGroupMemberResponse
 */
export class RemoveGroupMemberResponse extends Message<RemoveGroupMemberResponse> {
  constructor(data?: PartialMessage<RemoveGroupMemberResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.RemoveGroupMemberResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RemoveGroupMemberResponse {
    return new RemoveGroupMemberResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RemoveGroupMemberResponse {
    return new RemoveGroupMemberResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RemoveGroupMemberResponse {
    return new RemoveGroupMemberResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RemoveGroupMemberResponse | PlainMessage<RemoveGroupMemberResponse> | undefined, b: RemoveGroupMemberResponse | PlainMessage<RemoveGroupMemberResponse> | undefined): boolean {
    return proto3.util.equals(RemoveGroupMemberResponse, a, b);
  }
}

/**
 * @generated from message authentication.v1.GetGroupMembersRequest
 */
export class GetGroupMembersRequest extends Message<GetGroupMembersRequest> {
  /**
   * @generated from field: common.v1.PaginationRequest pagination = 1;
   */
  pagination?: PaginationRequest;

  /**
   * @generated from field: string group_id = 2;
   */
  groupId = "";

  constructor(data?: PartialMessage<GetGroupMembersRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.GetGroupMembersRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pagination", kind: "message", T: PaginationRequest },
    { no: 2, name: "group_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetGroupMembersRequest {
    return new GetGroupMembersRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetGroupMembersRequest {
    return new GetGroupMembersRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetGroupMembersRequest {
    return new GetGroupMembersRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetGroupMembersRequest | PlainMessage<GetGroupMembersRequest> | undefined, b: GetGroupMembersRequest | PlainMessage<GetGroupMembersRequest> | undefined): boolean {
    return proto3.util.equals(GetGroupMembersRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.GetGroupMembersResponse
 */
export class GetGroupMembersResponse extends Message<GetGroupMembersResponse> {
  /**
   * @generated from field: repeated authentication.v1.GroupMember members = 1;
   */
  members: GroupMember[] = [];

  constructor(data?: PartialMessage<GetGroupMembersResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.GetGroupMembersResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "members", kind: "message", T: GroupMember, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetGroupMembersResponse {
    return new GetGroupMembersResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetGroupMembersResponse {
    return new GetGroupMembersResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetGroupMembersResponse {
    return new GetGroupMembersResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetGroupMembersResponse | PlainMessage<GetGroupMembersResponse> | undefined, b: GetGroupMembersResponse | PlainMessage<GetGroupMembersResponse> | undefined): boolean {
    return proto3.util.equals(GetGroupMembersResponse, a, b);
  }
}

/**
 * @generated from message authentication.v1.GetUserGroupsRequest
 */
export class GetUserGroupsRequest extends Message<GetUserGroupsRequest> {
  /**
   * @generated from field: common.v1.PaginationRequest pagination = 1;
   */
  pagination?: PaginationRequest;

  /**
   * @generated from field: string user_id = 2;
   */
  userId = "";

  constructor(data?: PartialMessage<GetUserGroupsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.GetUserGroupsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pagination", kind: "message", T: PaginationRequest },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserGroupsRequest {
    return new GetUserGroupsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserGroupsRequest {
    return new GetUserGroupsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserGroupsRequest {
    return new GetUserGroupsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserGroupsRequest | PlainMessage<GetUserGroupsRequest> | undefined, b: GetUserGroupsRequest | PlainMessage<GetUserGroupsRequest> | undefined): boolean {
    return proto3.util.equals(GetUserGroupsRequest, a, b);
  }
}

/**
 * @generated from message authentication.v1.GetUserGroupsResponse
 */
export class GetUserGroupsResponse extends Message<GetUserGroupsResponse> {
  /**
   * @generated from field: repeated authentication.v1.Group groups = 1;
   */
  groups: Group[] = [];

  constructor(data?: PartialMessage<GetUserGroupsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "authentication.v1.GetUserGroupsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "groups", kind: "message", T: Group, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserGroupsResponse {
    return new GetUserGroupsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserGroupsResponse {
    return new GetUserGroupsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserGroupsResponse {
    return new GetUserGroupsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserGroupsResponse | PlainMessage<GetUserGroupsResponse> | undefined, b: GetUserGroupsResponse | PlainMessage<GetUserGroupsResponse> | undefined): boolean {
    return proto3.util.equals(GetUserGroupsResponse, a, b);
  }
}

