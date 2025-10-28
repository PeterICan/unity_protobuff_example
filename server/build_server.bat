@echo off
SET SCRIPT_DIR=%~dp0

pushd %SCRIPT_DIR%

ECHO Building Go server executable...
go build -o server_executable.exe internal/main.go

IF %ERRORLEVEL% NEQ 0 (
    ECHO Error: Go build failed!
    popd
    EXIT /B 1
)

ECHO Go server executable built successfully: %SCRIPT_DIR%server_executable.exe
popd
EXIT /B 0