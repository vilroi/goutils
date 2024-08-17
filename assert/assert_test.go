package assert

import "testing"

func TestAssertEq(t *testing.T) {
	AssertEq("yo", "yo")
	AssertEq(10, 10)
	//AssertEq(10, "10")
	//AssertEq("re", "10")
}
