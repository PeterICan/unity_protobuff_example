@echo off
SET SCRIPT_DIR=%~dp0

pushd %SCRIPT_DIR%

echo Building Docker Compose services...
set COMPOSE_PROJECT_NAME=ws_server_example
docker-compose -f .\docker-compose.yml up -d --build
IF %ERRORLEVEL% NEQ 0 (
    ECHO Error: Docker Compose build failed!
    popd
    EXIT /B 1
)
ECHO Docker Compose services built successfully.
popd
EXIT /B 0