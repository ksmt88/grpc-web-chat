/**
 * @fileoverview gRPC-Web generated client stub for chat
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';

import {
  Message,
  Result} from './chat_pb';

export class ChatClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoGetMessages = new grpcWeb.AbstractClientBase.MethodInfo(
    Message,
    (request: google_protobuf_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    Message.deserializeBinary
  );

  getMessages(
    request: google_protobuf_empty_pb.Empty,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/chat.Chat/GetMessages',
      request,
      metadata || {},
      this.methodInfoGetMessages);
  }

  methodInfoPostMessage = new grpcWeb.AbstractClientBase.MethodInfo(
    Result,
    (request: Message) => {
      return request.serializeBinary();
    },
    Result.deserializeBinary
  );

  postMessage(
    request: Message,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Result) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/chat.Chat/PostMessage',
      request,
      metadata || {},
      this.methodInfoPostMessage,
      callback);
  }

}

