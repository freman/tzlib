package tzlib

import (
	"fmt"
	"time"
)

// LocalOlsonZoneName tries really hard to figure out your local zone name as the Olson (IANA) database expects it
func LocalOlsonZoneName() (string, error) {
	tmp := time.Local
	if tmp.String() != "Local" {
		return tmp.String(), nil
	}

	loc, err := platformLocalOlsonZoneName()
	if err != nil {
		return "", fmt.Errorf("please set your TZ variable %s (%w)", platformSymlinkErrorStr, err)
	}

	return loc, nil
}
