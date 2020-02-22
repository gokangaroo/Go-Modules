package main

import "syscall/js"

// GOARCH=wasm GOOS=js go build -o test.wasm main.go
// cp $(go env GOROOT)/misc/wasm/wasm_exec.{html,js} .
// window.sumNative(1, 2,3)
func main() {
	registerCallbacks()
	select {}
}

func sum(this js.Value, params []js.Value) interface{} {
	result := 0
	for _, value := range params {
		result += value.Int()
	}
	js.Global().Call("wamsCallback", result)
	return result
}

func registerCallbacks() {
	js.Global().Set("sumNative", js.FuncOf(sum))
}
