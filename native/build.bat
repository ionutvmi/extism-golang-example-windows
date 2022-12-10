set CGO_CFLAGS=-I %~dp0libs\extism-0.1.0
set CGO_LDFLAGS=-L %~dp0libs\extism-0.1.0
go build -o build/program.exe main.go

tinygo build -o build/vocals.wasm -target wasi vocals.go

copy libs\extism-0.1.0\extism.dll build\extism.dll
