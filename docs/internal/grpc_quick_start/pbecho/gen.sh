#!/bin/bash
protoc --go_out=plugins=grpc:. ./echo.proto