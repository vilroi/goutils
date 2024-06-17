package linux

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type MapEntry struct {
	StartAddr uint
	EndAddr   uint
	Perms     string
	Offset    uint
	DevMajor  int
	DevMinor  int
	Inode     uint
	Path      string
}

func (m *MapEntry) Size() uint {
	return m.EndAddr - m.StartAddr
}

func (m MapEntry) String() string {
	return fmt.Sprintf("%#x-%#x %s %x %x:%x %d %s", m.StartAddr, m.EndAddr, m.Perms,
		m.Offset, m.DevMajor, m.DevMinor, m.Inode, m.Path)
}

func ParseProcMaps(pid int) []MapEntry {
	path := fmt.Sprintf("/proc/%d/maps", pid)
	f, err := os.Open(path)
	check(err)

	scanner := bufio.NewScanner(f)

	var maps []MapEntry
	for scanner.Scan() {
		line := scanner.Text()

		var ent MapEntry
		_, err := fmt.Sscanf(line, "%x-%x %s %x %x:%x %d %s", &ent.StartAddr, &ent.EndAddr, &ent.Perms,
			&ent.Offset, &ent.DevMajor, &ent.DevMinor, &ent.Inode, &ent.Path)

		// memory regrion not backed by file
		if err == io.EOF {
			_, err = fmt.Sscanf(line, "%x-%x %s %x %x:%x %d", &ent.StartAddr, &ent.EndAddr, &ent.Perms,
				&ent.Offset, &ent.DevMajor, &ent.DevMinor, &ent.Inode)
		}
		check(err)

		maps = append(maps, ent)

	}

	return maps
}
