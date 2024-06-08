package linux

import (
	"strings"

	"github.com/vilroi/goutils"
)

const (
	passwdPath string = "/etc/passwd"
	groupPath  string = "/etc/group"
)

type Passwd struct {
	Name    string
	Passwd  string
	Uid     uint32
	Gid     uint32
	Gecos   string
	Homedir string
	Shell   string
}

type Group struct {
	Name    string
	Passwd  string
	Gid     uint32
	Members []string
}

func UidToUsername(uid uint32) string {
	users := Getpwents()

	for _, user := range users {
		if user.Uid == uid {
			return user.Name
		}
	}
	return ""
}

func GidToName(gid uint32) string {
	groups := Getgrents()

	for _, grp := range groups {
		if grp.Gid == gid {
			return grp.Name
		}
	}
	return ""
}

func Getpwents() []Passwd {
	var entries []Passwd

	scanner := goutils.MakeFileScanner(passwdPath)
	for scanner.Scan() {
		ent := parsepwent(scanner.Text())
		entries = append(entries, ent)
	}

	return entries
}

func parsepwent(s string) Passwd {
	parts := strings.Split(s, ":")

	var passwd Passwd
	passwd.Name = parts[0]
	passwd.Passwd = parts[1]
	passwd.Uid = uint32(goutils.Atoi(parts[2]))
	passwd.Gid = uint32(goutils.Atoi(parts[3]))
	passwd.Gecos = parts[4]
	passwd.Homedir = parts[5]
	passwd.Shell = parts[6]

	return passwd
}

func Getgrents() []Group {
	var entries []Group

	scanner := goutils.MakeFileScanner(groupPath)
	for scanner.Scan() {
		group := parsegrent(scanner.Text())
		entries = append(entries, group)
	}

	return entries
}

func parsegrent(s string) Group {
	parts := strings.Split(s, ":")

	var group Group
	group.Name = parts[0]
	group.Passwd = parts[1]
	group.Gid = uint32(goutils.Atoi(parts[2]))
	group.Members = strings.Split(parts[3], ",")

	return group
}
