echo Start
set WorkPath=%cd%
set PROTOCEXE=%WorkPath%\tools\protoc\bin\protoc.exe
set PROTOCLuaEXEBat=%WorkPath%\tools\lua\protoc-gen-lua-master\plugin\protoc-gen-lua.bat
set TARGET_PATH=./lua

for /f "delims=" %%i in ('dir /b proto "proto/*.proto"') do (
    REM ::Lua

    REM echo %PROTOCEXE% --plugin=protoc-gen-lua=%PROTOCLuaEXEBat% --lua_out=%Lua_TARGET_PATH% proto/%%i
    REM %PROTOCEXE% --plugin=protoc-gen-lua=%PROTOCLuaEXEBat% --lua_out=%Lua_TARGET_PATH% proto/%%i

    echo %PROTOCEXE% --plugin=protoc-gen-lua=%PROTOCLuaEXEBat% -I=%WorkPath%\tools\protoc\include -I=%WorkPath%\tools\protoc\frankee\protobuf -I=./proto  --lua_out=%TARGET_PATH% proto/%%i 
    %PROTOCEXE% --plugin=protoc-gen-lua=%PROTOCLuaEXEBat% -I=%WorkPath%\tools\protoc\include -I=%WorkPath%\tools\protoc\frankee\protobuf -I=./proto  --lua_out=%TARGET_PATH% proto/%%i 
)
echo End
pause