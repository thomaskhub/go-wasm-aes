package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("encrypt", js.FuncOf(EncryptJs))
	js.Global().Set("decrypt", js.FuncOf(DecryptJs))

	select {}
}

func EncryptJs(this js.Value, args []js.Value) interface{} {

	password := args[0].String()
	salt := args[1].String()
	data := args[2].String()
	encData := Encrypt(password, salt, data)

	return js.ValueOf(encData)
}

func DecryptJs(this js.Value, args []js.Value) interface{} {

	password := args[0].String()
	salt := args[1].String()
	data := args[2].String()
	decData := Decrypt(password, salt, data)

	return js.ValueOf(decData)
}
