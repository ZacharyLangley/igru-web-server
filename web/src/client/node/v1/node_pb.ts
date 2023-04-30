// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file node/v1/node.proto (package node.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Node } from "./schema_pb";

/**
 * @generated from message node.v1.UpdateNodeRequest
 */
export class UpdateNodeRequest extends Message<UpdateNodeRequest> {
  /**
   * @generated from field: string mac_address = 1;
   */
  macAddress = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: map<string, string> custom_labels = 4;
   */
  customLabels: { [key: string]: string } = {};

  constructor(data?: PartialMessage<UpdateNodeRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "node.v1.UpdateNodeRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "mac_address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "custom_labels", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateNodeRequest {
    return new UpdateNodeRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateNodeRequest {
    return new UpdateNodeRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateNodeRequest {
    return new UpdateNodeRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateNodeRequest | PlainMessage<UpdateNodeRequest> | undefined, b: UpdateNodeRequest | PlainMessage<UpdateNodeRequest> | undefined): boolean {
    return proto3.util.equals(UpdateNodeRequest, a, b);
  }
}

/**
 * @generated from message node.v1.UpdateNodeResponse
 */
export class UpdateNodeResponse extends Message<UpdateNodeResponse> {
  /**
   * @generated from field: node.v1.Node node = 1;
   */
  node?: Node;

  constructor(data?: PartialMessage<UpdateNodeResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "node.v1.UpdateNodeResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "node", kind: "message", T: Node },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateNodeResponse {
    return new UpdateNodeResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateNodeResponse {
    return new UpdateNodeResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateNodeResponse {
    return new UpdateNodeResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateNodeResponse | PlainMessage<UpdateNodeResponse> | undefined, b: UpdateNodeResponse | PlainMessage<UpdateNodeResponse> | undefined): boolean {
    return proto3.util.equals(UpdateNodeResponse, a, b);
  }
}

/**
 * @generated from message node.v1.GetNodesRequest
 */
export class GetNodesRequest extends Message<GetNodesRequest> {
  constructor(data?: PartialMessage<GetNodesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "node.v1.GetNodesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetNodesRequest {
    return new GetNodesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetNodesRequest {
    return new GetNodesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetNodesRequest {
    return new GetNodesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetNodesRequest | PlainMessage<GetNodesRequest> | undefined, b: GetNodesRequest | PlainMessage<GetNodesRequest> | undefined): boolean {
    return proto3.util.equals(GetNodesRequest, a, b);
  }
}

/**
 * @generated from message node.v1.GetNodesResponse
 */
export class GetNodesResponse extends Message<GetNodesResponse> {
  /**
   * @generated from field: repeated node.v1.Node nodes = 1;
   */
  nodes: Node[] = [];

  constructor(data?: PartialMessage<GetNodesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "node.v1.GetNodesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nodes", kind: "message", T: Node, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetNodesResponse {
    return new GetNodesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetNodesResponse {
    return new GetNodesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetNodesResponse {
    return new GetNodesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetNodesResponse | PlainMessage<GetNodesResponse> | undefined, b: GetNodesResponse | PlainMessage<GetNodesResponse> | undefined): boolean {
    return proto3.util.equals(GetNodesResponse, a, b);
  }
}

/**
 * @generated from message node.v1.GetNodeRequest
 */
export class GetNodeRequest extends Message<GetNodeRequest> {
  /**
   * @generated from field: string mac_address = 1;
   */
  macAddress = "";

  constructor(data?: PartialMessage<GetNodeRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "node.v1.GetNodeRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "mac_address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetNodeRequest {
    return new GetNodeRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetNodeRequest {
    return new GetNodeRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetNodeRequest {
    return new GetNodeRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetNodeRequest | PlainMessage<GetNodeRequest> | undefined, b: GetNodeRequest | PlainMessage<GetNodeRequest> | undefined): boolean {
    return proto3.util.equals(GetNodeRequest, a, b);
  }
}

/**
 * @generated from message node.v1.GetNodeResponse
 */
export class GetNodeResponse extends Message<GetNodeResponse> {
  /**
   * @generated from field: node.v1.Node node = 1;
   */
  node?: Node;

  constructor(data?: PartialMessage<GetNodeResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "node.v1.GetNodeResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "node", kind: "message", T: Node },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetNodeResponse {
    return new GetNodeResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetNodeResponse {
    return new GetNodeResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetNodeResponse {
    return new GetNodeResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetNodeResponse | PlainMessage<GetNodeResponse> | undefined, b: GetNodeResponse | PlainMessage<GetNodeResponse> | undefined): boolean {
    return proto3.util.equals(GetNodeResponse, a, b);
  }
}

