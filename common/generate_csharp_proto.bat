@echo off
rem This script generates C# code from the .proto file.

pushd %~dp0

set PROTO_ROOT_DIR=.\proto
set CSHARP_OUT_DIR=..\client\Assets\Generated

echo Ensuring C# output directory exists: %CSHARP_OUT_DIR%
if not exist "%CSHARP_OUT_DIR%" (
    mkdir "%CSHARP_OUT_DIR%"
)

echo Generating C# code for message.proto...
protoc --proto_path=%PROTO_ROOT_DIR% --csharp_out=%CSHARP_OUT_DIR% message.proto

echo Generating C# code for json_api protos...
protoc --proto_path=%PROTO_ROOT_DIR% --csharp_out=%CSHARP_OUT_DIR% %PROTO_ROOT_DIR%\json_api\*.proto

echo C# code generation complete.

popd