@echo off
rem This script generates C# code from the .proto file.

set PROTO_DIR=.\proto
set CSHARP_OUT_DIR=..\client\Assets\Generated

echo Ensuring C# output directory exists: %CSHARP_OUT_DIR%
if not exist "%CSHARP_OUT_DIR%" (
    mkdir "%CSHARP_OUT_DIR%"
)

echo Generating C# code...
protoc --proto_path=%PROTO_DIR% --csharp_out=%CSHARP_OUT_DIR% message.proto

echo C# code generation complete.