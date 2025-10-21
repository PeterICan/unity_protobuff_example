@echo off
:: 步驟 0: 設定專案根目錄並移除結尾的反斜線
SET PROJECT_ROOT_WITH_SLASH=%~dp0

:: 使用 Batch 語法移除最後一個字元 (\)
SET PROJECT_ROOT=%PROJECT_ROOT_WITH_SLASH:~0,-1%

:: 1. 設定 Unity 版本設定檔的路徑
:: 注意：因為 PROJECT_ROOT 結尾已無反斜線，這裡需要手動補上 \
SET VERSION_FILE="%PROJECT_ROOT%\ProjectSettings\ProjectVersion.txt"


ECHO ----------------------------------------------------
ECHO Step 1: Check and read the project version
ECHO ----------------------------------------------------

IF NOT EXIST %VERSION_FILE% (
    ECHO.
    ECHO Error: Cannot find the Unity version configuration file %VERSION_FILE%. Please verify the script location.
    EXIT /B 1
)

:: Read the version number
SET UNITY_VERSION=
FOR /F "tokens=2 delims= " %%i IN ('TYPE %VERSION_FILE% ^| FINDSTR /R "m_EditorVersion:" ^| FINDSTR /V "Revision"') DO (
    SET UNITY_VERSION=%%i
)

IF "%UNITY_VERSION%"=="" (
    ECHO.
    ECHO Error: Unable to parse the Unity version number from the configuration file.
    EXIT /B 1
)

ECHO Successfully read the required project version: %UNITY_VERSION%

:: 2. Construct the Unity editor path and check the executable
SET UNITY_PATH_DIR=C:\Program Files\Unity\Hub\Editor\%UNITY_VERSION%\Editor
SET UNITY_PATH="%UNITY_PATH_DIR%\Unity.exe"

ECHO.
ECHO Step 2: Check if the Unity editor executable exists
ECHO Editor path: %UNITY_PATH%

IF NOT EXIST %UNITY_PATH% (
    ECHO.
    ECHO Critical Error: Cannot find the specified Unity editor executable!
    ECHO Path: %UNITY_PATH%
    ECHO Please install version %UNITY_VERSION% via Unity Hub.
    EXIT /B 1
)

:: 3. 設定其他參數
SET BUILD_METHOD=BuildScript.PerformBuild
:: 注意：LOG_FILE 也需要手動補上 \
SET LOG_FILE="%PROJECT_ROOT%\build_log.txt"

ECHO.
ECHO Step 3: Execute the Unity build command
ECHO Log output: %LOG_FILE%
ECHO Project path: "%PROJECT_ROOT%"
ECHO ----------------------------------------------------


:: Ensure all paths with spaces are enclosed in double quotes
%UNITY_PATH% ^
-batchmode ^
-quit ^
-nographics ^
-projectPath "%PROJECT_ROOT%" ^
-executeMethod %BUILD_METHOD% ^
-logFile "%LOG_FILE%"

:: Check the error code (ErrorLevel) from the Unity execution result
IF %ERRORLEVEL% NEQ 0 (
    ECHO.
    ECHO Error: Unity build process failed! Please check %LOG_FILE% (if it exists) for more details.
    EXIT /B 1
)

ECHO.
ECHO Unity build completed successfully.
EXIT /B 0