SCRIPT=$(realpath "$0")
SCRIPTPATH=$(dirname "$SCRIPT")

# replace drive letter
SCRIPTPATH=${SCRIPTPATH/\/c/C:}
SCRIPTPATH=${SCRIPTPATH/\/d/D:}

export CGO_CFLAGS="-I $SCRIPTPATH/libs/extism-0.1.0"
export CGO_LDFLAGS="-L $SCRIPTPATH/libs/extism-0.1.0"


go build -o out/program.exe main.go

tinygo build -o out/vowels.wasm -target wasi vowels.go

cp libs/extism-0.1.0/extism.dll out/extism.dll
