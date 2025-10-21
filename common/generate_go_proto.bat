@echo off
rem This script generates Go code from the .proto file.

pushd %~dp0

set PROTO_ROOT_DIR=.\proto
set GO_OUT_DIR=..\server\generated

echo Ensuring Go output directory exists: %GO_OUT_DIR%
if not exist "%GO_OUT_DIR%" (
    mkdir "%GO_OUT_DIR%"
)

echo Generating Go code for message.proto...
protoc --proto_path=%PROTO_ROOT_DIR% --go_out=paths=source_relative:%GO_OUT_DIR% message.proto

echo Generating Go code for json_api protos...
protoc --proto_path=%PROTO_ROOT_DIR% --go_out=paths=source_relative:%GO_OUT_DIR% %PROTO_ROOT_DIR%\json_api\*.proto

echo Go code generation complete.

popd