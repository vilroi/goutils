package goutils

import (
	"fmt"
	"os"
	"reflect"
)

func Err(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func SizeOf(a any) uintptr {
	t := reflect.TypeOf(a)
	return t.Size()
}
