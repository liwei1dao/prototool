@echo off

echo Start

set WorkPath=%cd%

::protoc tool
set PROTOCEXE=%WorkPath%\Tools\protoc\protoc.exe
set PROTOCGOEXE=%WorkPath%\Tools\go\protoc-gen-go.exe
set PROTOCLuaEXEBat=%WorkPath%\Tools\lua\protoc-gen-lua-master\plugin\protoc-gen-lua.bat
::Go
set Go_TARGET_PATH=./Go
::Cs
set Cs_TARGET_PATH=./Cs
::Java
set JAVA_TARGET_PATH=./Java
::Lua 
set Lua_TARGET_PATH=./Lua

for /f "delims=" %%i in ('dir /b proto "proto/*.proto"') do (
    
    REM REM ::Go
    REM echo %PROTOCEXE% --plugin=protoc-gen-go=%PROTOCGOEXE% --go_out=%Go_TARGET_PATH% proto/%%i
    REM %PROTOCEXE% --plugin=protoc-gen-go=%PROTOCGOEXE% --go_out=%Go_TARGET_PATH% proto/%%i
    
    REM ::Cs
    echo %PROTOCEXE%  --csharp_out=%Cs_TARGET_PATH% proto/%%i -I./proto
    %PROTOCEXE%  --csharp_out=%Cs_TARGET_PATH% proto/%%i -I./proto

    ::Lua
    REM echo %PROTOCEXE% --plugin=protoc-gen-lua=%PROTOCLuaEXEBat% --lua_out=%Lua_TARGET_PATH% proto/%%i
    REM %PROTOCEXE% --plugin=protoc-gen-lua=%PROTOCLuaEXEBat% --lua_out=%Lua_TARGET_PATH% proto/%%i

    REM ::JAVA
    REM echo %PROTOCEXE% --java_out=%JAVA_TARGET_PATH% proto/%%i
    REM %PROTOCEXE% --java_out=%JAVA_TARGET_PATH% proto/%%i

)

echo End

pause