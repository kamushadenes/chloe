package utils

import (
	"fmt"
	"os"
	"runtime"
)

func HandlePanic(r interface{}) error {
	var f *os.File
	f, err := os.OpenFile("panic.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error opening panic.log file:", err.Error())
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	switch e := r.(type) {
	case string:
		if _, err = f.WriteString(e); err != nil {
			fmt.Println("error writing to panic.log file:", e)
		}
	case runtime.Error:
		if _, err = f.WriteString(e.Error()); err != nil {
			fmt.Println("error writing to panic.log file:", e.Error())
		}
	case error:
		if _, err = f.WriteString(e.Error()); err != nil {
			fmt.Println("error writing to panic.log file:", e.Error())
		}
	default:
		if _, err = f.WriteString(fmt.Sprintf("%+v", e)); err != nil {
			fmt.Println("error writing to panic.log file:", e)
		}
	}
	_, _ = f.WriteString("\n\n")

	return err
}
