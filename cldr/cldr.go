package cldr

//go:generate go run embedded/generator/main.go

import (
	"github.com/freman/tzlib/cldr/embedded"
)

type SupplementalData interface {
	ToWindows(zoneName string) string
	ToOlson(zoneName string) string
}

var DefaultSupplimentalData SupplementalData = &embedded.SupplementalData{}

func ToWindows(zoneName string) string {
	return DefaultSupplimentalData.ToWindows(zoneName)
}
func ToOlson(zoneName string) string {
	return DefaultSupplimentalData.ToOlson(zoneName)
}
