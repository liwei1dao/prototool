echo Start
set WorkPath=%cd%
set PROTOCEXE=%WorkPath%\tools\protoc\bin\protoc.exe
set TARGET_PATH=./java

for /f "delims=" %%i in ('dir /b proto "proto/*.proto"') do (
    REM ::Java
    echo %PROTOCEXE%  -I=%WorkPath%\tools\protoc\include -I=%WorkPath%\tools\protoc\frankee\protobuf -I=./proto  --java_out=%TARGET_PATH% proto/%%i 
    %PROTOCEXE% -I=%WorkPath%\tools\protoc\include -I=%WorkPath%\tools\protoc\frankee\protobuf -I=./proto  --java_out=%TARGET_PATH% proto/%%i 
)
echo End
pause