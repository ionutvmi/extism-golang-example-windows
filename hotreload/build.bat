set CGO_CFLAGS=-I %~dp0libs\extism-0.1.0
set CGO_LDFLAGS=-L %~dp0libs\extism-0.1.0
go build -o out/program.exe main.go

build_plugin.bat

copy libs\extism-0.1.0\extism.dll out\extism.dll
