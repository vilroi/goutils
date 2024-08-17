package assert

import "testing"

func TestAssertEq(t *testing.T) {
	AssertEq("yo", "yo")
	AssertEq(10, 10)
	//AssertEq(10, "10")
	//AssertEq("re", "10")
}

func TestAssertNotEq(t *testing.T) {
	//AssertNotEq(10, "10")
	//AssertNotEq("yo", "yo")
	AssertNotEq("re", "10")
}
