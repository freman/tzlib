package xml

import (
	"encoding/xml"
	"io"
	"strings"
)

// SupplementalData makes decoding the structure of the xml data found at
// https://github.com/unicode-org/cldr/blob/master/common/supplemental/windowsZones.xml
// a good deal easier with a couple of helpers for shits and giggles
type SupplementalData struct {
	XMLName xml.Name `xml:"supplementalData"`
	Text    string   `xml:",chardata"`
	Version struct {
		Text   string `xml:",chardata"`
		Number string `xml:"number,attr"`
	} `xml:"version"`
	WindowsZones struct {
		Text         string `xml:",chardata"`
		MapTimezones struct {
			Text         string   `xml:",chardata"`
			OtherVersion string   `xml:"otherVersion,attr"`
			TypeVersion  string   `xml:"typeVersion,attr"`
			MapZone      mapZones `xml:"mapZone"`
		} `xml:"mapTimezones"`
	} `xml:"windowsZones"`
}

type mapZones []struct {
	Text      string `xml:",chardata"`
	Other     string `xml:"other,attr"`
	Territory string `xml:"territory,attr"`
	Type      string `xml:"type,attr"`
}

// ReadSupplementalData takes the given reader which hopefully is pointed at
// a xml stream, parses that xml stream and returns a SupplimentalData object
func ReadSupplementalData(r io.Reader) (supData SupplementalData, err error) {
	err = xml.NewDecoder(r).Decode(&supData)
	return
}

// UniqueTypeToOther returns a list of unique Olson types to windows types
func (m mapZones) UniqueTypeToOther() map[string]string {
	res := make(map[string]string)

	for _, zone := range m {
		for _, name := range strings.Fields(zone.Type) {
			if _, found := res[name]; found {
				continue
			}
			res[name] = zone.Other
		}
	}
	return res
}

// OtherToDefaultType returns a list of windows types to the "default" Olsen type
// as sadly windows lumps a whole bunch of zones together and I have no idea how
// to get "Territory" for windows.
func (m mapZones) OtherToDefaultType() map[string]string {
	res := make(map[string]string)
	for _, zone := range m {
		if zone.Territory == "001" {
			res[zone.Other] = zone.Type
		}
	}
	return res
}

// ToWindows will give you the Windows zone name for a given Olsen one
func (s SupplementalData) ToWindows(zoneName string) string {
	for _, zone := range s.WindowsZones.MapTimezones.MapZone {
		for _, name := range strings.Fields(zone.Type) {
			if zoneName == name {
				return zone.Other
			}
		}
	}
	return ""
}

// ToOlsen will return the "Default" Olsen type for a given Windows one
func (s SupplementalData) ToOlson(zoneName string) string {
	for _, zone := range s.WindowsZones.MapTimezones.MapZone {
		if zone.Other == zoneName && zone.Territory == "001" {
			return zone.Type
		}
	}
	return ""
}
