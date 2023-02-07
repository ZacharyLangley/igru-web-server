// @generated by protoc-gen-connect-web v0.6.0 with parameter "target=ts"
// @generated from file authentication/v1/session.proto (package authentication.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CheckSessionPermissionsRequest, CheckSessionPermissionsResponse, CreateSessionRequest, CreateSessionResponse, DeleteSessionRequest, DeleteSessionResponse, GetSessionsRequest, GetSessionsResponse } from "./session_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service authentication.v1.SessionService
 */
export const SessionService = {
  typeName: "authentication.v1.SessionService",
  methods: {
    /**
     * @generated from rpc authentication.v1.SessionService.CreateSession
     */
    createSession: {
      name: "CreateSession",
      I: CreateSessionRequest,
      O: CreateSessionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc authentication.v1.SessionService.GetSessions
     */
    getSessions: {
      name: "GetSessions",
      I: GetSessionsRequest,
      O: GetSessionsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc authentication.v1.SessionService.DeleteSession
     */
    deleteSession: {
      name: "DeleteSession",
      I: DeleteSessionRequest,
      O: DeleteSessionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc authentication.v1.SessionService.CheckSessionPermissions
     */
    checkSessionPermissions: {
      name: "CheckSessionPermissions",
      I: CheckSessionPermissionsRequest,
      O: CheckSessionPermissionsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
