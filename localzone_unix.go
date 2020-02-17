// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

package tzlib

import (
	"errors"
	"os"
	"strings"
	"time"
)

const platformSymlinkErrorStr = `(eg: export TZ="Australia/Brisbane") or fix the /etc/timezone symlink`

func platformLocalOlsonZoneName() (string, error) {
	originFile, err := os.Readlink("/etc/localtime")
	if err != nil {
		return "", err
	}
	paths := strings.Split(originFile, "/zoneinfo/")
	if len(paths) < 2 {
		return "", errors.New("invalid symlink path")
	}
	tz, err := time.LoadLocation(paths[1])
	if err != nil {
		return "", err
	}
	return tz.String(), nil
}
