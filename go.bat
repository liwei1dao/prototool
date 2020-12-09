echo Start
set WorkPath=%cd%
set PROTOCEXE=%WorkPath%\tools\protoc\bin\protoc.exe
set PROTOCGOEXE=%WorkPath%\tools\go\protoc-gen-go.exe
set TARGET_PATH=./go

for /f "delims=" %%i in ('dir /b proto "proto/*.proto"') do (
    REM ::Go
    echo %PROTOCEXE% --plugin=protoc-gen-go=%PROTOCGOEXE% --go_out=%Go_TARGET_PATH% proto/%%i
    %PROTOCEXE% --plugin=protoc-gen-go=%PROTOCGOEXE% --go_out=%Go_TARGET_PATH% proto/%%i
)
echo End
pause