echo Start
set WorkPath=%cd%
set PROTOCEXE=%WorkPath%\tools\protoc\bin\protoc.exe
set PROTOCGOEXE=%WorkPath%\tools\go\protoc-gen-go.exe
set PROTOCGOEXE_GOGO=%WorkPath%\tools\go\protoc-gen-gogo.exe
set TARGET_PATH=./go

for /f "delims=" %%i in ('dir /b proto "proto/*.proto"') do (
    REM ::Go
    echo %PROTOCEXE% --plugin=protoc-gen-go=%PROTOCGOEXE% --plugin=protoc-gen-gogo=%PROTOCGOEXE_GOGO% -I=%WorkPath%\tools\protoc\include -I=%WorkPath%\tools\protoc\frankee\protobuf -I=./proto  --gogo_out=%TARGET_PATH% proto/%%i 
    %PROTOCEXE% --plugin=protoc-gen-go=%PROTOCGOEXE% --plugin=protoc-gen-gogo=%PROTOCGOEXE_GOGO% -I=%WorkPath%\tools\protoc\include -I=%WorkPath%\tools\protoc\frankee\protobuf -I=./proto  --gogo_out=%TARGET_PATH% proto/%%i 
)
echo End
pause