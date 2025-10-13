@echo off
rem This script generates Go code from the .proto file.

set PROTO_DIR=..\proto
set GO_OUT_DIR=..\server\generated

echo Ensuring Go output directory exists: %GO_OUT_DIR%
if not exist "%GO_OUT_DIR%" (
    mkdir "%GO_OUT_DIR%"
)

echo Generating Go code...
protoc --proto_path=%PROTO_DIR% --go_out=%GO_OUT_DIR% message.proto

echo Go code generation complete.