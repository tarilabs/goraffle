

## Building

```
GOOS=js GOARCH=wasm go build -o build/raffle.wasm
cp $(go env GOROOT)/misc/wasm/wasm_exec.js build/
```

