package utils

import (
	"os"
)

func RecoverPanic() {
	if r := recover(); r != nil {
		f, err := os.OpenFile("panic.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer func(f *os.File) {
			_ = f.Close()
		}(f)

		if _, err := f.WriteString(r.(string)); err != nil {
			panic(err)
		}

		os.Exit(1)
	}
}

func HandlePanicAsync(fn func()) {
	go func() {
		defer RecoverPanic()
		fn()
	}()
}
