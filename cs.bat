echo Start
set WorkPath=%cd%
set PROTOCEXE=%WorkPath%\tools\protoc\bin\protoc.exe
set TARGET_PATH=./cs

for /f "delims=" %%i in ('dir /b proto "proto/*.proto"') do (
    REM ::Cs
    echo %PROTOCEXE%  -I=%WorkPath%\tools\protoc\include -I=%WorkPath%\tools\protoc\frankee\protobuf -I=./proto  --csharp_out=%TARGET_PATH% proto/%%i 
    %PROTOCEXE% -I=%WorkPath%\tools\protoc\include -I=%WorkPath%\tools\protoc\frankee\protobuf -I=./proto  --csharp_out=%TARGET_PATH% proto/%%i 
)
echo End
pause