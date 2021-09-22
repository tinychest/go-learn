@echo off

set GOARCH=amd64
set GOOS=windows

set ts=%date:~2,2%%date:~5,2%%date:~8,2%.%time:~0,2%%time:~3,2%%time:~6,2%
set ts=%ts: =0%

set proj_name=main
set exe_name=%proj_name%-%ts%.exe

go build -o %exe_name%