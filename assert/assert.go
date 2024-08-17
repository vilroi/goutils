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
		fmt.Fprintf(os.Stderr, "AssertEq failed\n")
		fmt.Fprintf(os.Stderr, "Expected:%v\nGot:%v\n", expected, got)
		os.Exit(1)
	}
}

func AssertNotEq(x, y any) {
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		fmt.Fprintf(os.Stderr, "Cannot compare %v(%T) with %v(%T)\n", x, x, y, y)
		os.Exit(1)
	}

	if x == y {
		fmt.Fprintf(os.Stderr, "AssertNotEq failed\n")
		fmt.Fprintf(os.Stderr, "%v == %v\n", x, y)
		os.Exit(1)
	}
}
