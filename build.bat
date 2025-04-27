@echo off
REM build.bat - Windows build script for go-jocy

REM Uncomment these lines if you want to build the frontend
REM cd path\to\your\jocy-web
REM call yarn build
REM cd %~dp0
REM 
REM REM Delete existing frontend files
REM if exist web\dist rmdir /s /q web\dist
REM 
REM REM Copy frontend files to project directory
REM xcopy /E /I path\to\your\jocy-web\dist web\dist
REM 
REM REM Package frontend files
REM cd web
REM go-bindata -o=bindata/bindata.go -pkg=bindata -prefix "dist" dist/...
REM cd ..

REM Build for Linux ARM64
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=arm64
go build -o Jocy-linux-arm64 -ldflags "-s -w -extldflags -static"

REM Build for Linux AMD64
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o Jocy-linux-amd64 -ldflags "-s -w -extldflags -static"

REM Build for Windows AMD64
set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -o Jocy-windows-amd64.exe -ldflags "-s -w -extldflags -static"

REM Compress executables with UPX (if installed)
where upx >nul 2>&1
if %ERRORLEVEL% EQU 0 (
    upx Jocy-*
) else (
    echo UPX not found. Skipping compression.
    echo Install UPX from https://github.com/upx/upx/releases if needed.
)

echo Build completed.
