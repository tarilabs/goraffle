

## Building

GOOS=js GOARCH=wasm go build -o raffle.wasm
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .

