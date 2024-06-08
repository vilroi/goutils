package goutils

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Err(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func SizeOf(a any) uintptr {
	t := reflect.TypeOf(a)
	return t.Size()
}

func MakeFileScanner(path string) *bufio.Scanner {
	f, err := os.Open(path)
	check(err)

	return bufio.NewScanner(f)
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	check(err)

	return i
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
