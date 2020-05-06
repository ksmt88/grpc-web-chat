#!/bin/bash

DIRNAME=`dirname $0`

protoc $DIRNAME/chat.proto --go_out=plugins=grpc:.
