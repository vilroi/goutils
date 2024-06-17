package linux

import (
	"fmt"
	"os"
	"testing"
)

func TestProcMaps(t *testing.T) {
	pid := os.Getpid()

	maps := ParseProcMaps(pid)
	for _, ent := range maps {
		fmt.Println(ent)
	}
}
