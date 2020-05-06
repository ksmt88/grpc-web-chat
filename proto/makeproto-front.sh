#!/bin/bash

DIRNAME=`dirname $0`

protoc $DIRNAME/chat.proto --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=typescript,mode=grpcwebtext:.
