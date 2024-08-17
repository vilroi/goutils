package assert

import (
	"fmt"
	"os"
	"reflect"
)

func AssertEq(expected, got any) {
	if reflect.TypeOf(expected) != reflect.TypeOf(got) {
		fmt.Fprintf(os.Stderr, "Cannot compare %v(%T) with %v(%T)\n", expected, expected, got, got)
		os.Exit(1)
	}

	if expected != got {
		fmt.Fprintf(os.Stderr, "Expected:%v\nGot:%v\n", expected, got)
		os.Exit(1)
	}
}
