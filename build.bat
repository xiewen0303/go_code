@echo off
set GOROOT=G:\golang\tools\go

set GOPATH=G:\golang\tools\lib;G:\golang\workspace\demo\gentools

go build -o GenCodeTools.exe -a ./src

pause