// +build windows

package tzlib

import (
	"strings"
	"syscall"
	"unicode/utf16"

	"github.com/freman/tzlib/cldr"
)

const platformSymlinkErrorStr = `(eg: set TZ="Australia/Brisbane") or use linux :P`

func platformLocalOlsonZoneName() (string, error) {
	var i syscall.Timezoneinformation
	if _, err := syscall.GetTimeZoneInformation(&i); err != nil {
		return "", err
	}

	name := strings.TrimRight(string(utf16.Decode(i.StandardName[:])), "\x00")
	return cldr.ToOlson(name), nil
}
