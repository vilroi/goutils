package linux

import (
	"fmt"
	"os"
	"testing"
)

func TestGetPwents(t *testing.T) {
	ents := Getpwents()
	for _, passwd := range ents {
		fmt.Println(passwd)
	}
}

func TestGetrents(t *testing.T) {
	ents := Getgrents()
	for _, grp := range ents {
		fmt.Println(grp)
	}
}

func Test(t *testing.T) {
	uid := os.Getuid()
	username := UidToUsername(uint32(uid))
	fmt.Printf("%s: %d\n", username, uid)

	gid := os.Getgid()
	grpname := GidToName(uint32(gid))
	fmt.Printf("%s: %d\n", grpname, uid)
}
