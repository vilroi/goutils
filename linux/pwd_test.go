package linux

import (
	"fmt"
	"os"
	"testing"

	"github.com/vilroi/goutils"
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
	username, err := UidToUsername(uint32(uid))
	check(err)

	fmt.Printf("%s: %d\n", username, uid)

	gid := os.Getgid()
	grpname, err := GidToName(uint32(gid))
	check(err)

	fmt.Printf("%s: %d\n", grpname, uid)
}

func TestErrors(t *testing.T) {
	_, err := UidToUsername(9999999)
	if err == nil {
		goutils.Err("UidToUsername should result in an error")
	}

	_, err = GidToName(9999999)
	if err == nil {
		goutils.Err("GidToName should result in an error")
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
