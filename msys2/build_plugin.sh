mkdir -p build
tinygo build -o build/vocals.wasm -target wasi vocals.go
