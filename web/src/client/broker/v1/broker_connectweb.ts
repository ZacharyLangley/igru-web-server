// @generated by protoc-gen-connect-web v0.6.0 with parameter "target=ts"
// @generated from file broker/v1/broker.proto (package broker.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { PingRequest, PingResponse } from "./broker_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service broker.v1.BrokerService
 */
export const BrokerService = {
  typeName: "broker.v1.BrokerService",
  methods: {
    /**
     * @generated from rpc broker.v1.BrokerService.Ping
     */
    ping: {
      name: "Ping",
      I: PingRequest,
      O: PingResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
