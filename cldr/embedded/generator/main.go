package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/freman/tzlib/cldr/xml"
)

func main() {
	c := &http.Client{
		Timeout: 30 * time.Second,
	}

	err := func() error {
		resp, err := c.Get("https://raw.githubusercontent.com/unicode-org/cldr/master/common/supplemental/windowsZones.xml")
		if err != nil {
			return fmt.Errorf("Failure to download windowsZones from cldr github: %w", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("Failure to download windowsZones from cldr github: %v", resp.Status)
		}

		supData, err := xml.ReadSupplementalData(resp.Body)

		embedded, err := os.Create("embedded/embedded.go")
		if err != nil {
			return fmt.Errorf("Can't write to embedded.go: %w", err)
		}
		defer func() {
			cerr := embedded.Close()
			if err == nil && cerr != nil {
				err = cerr
			}
		}()

		if err := template.Must(template.New("embedded.gotmpl").Funcs(template.FuncMap{
			"maxLen": func(in map[string]string) (max int) {
				for i := range in {
					if l := len(i); l > max {
						max = l
					}
				}
				return max
			},
			"repeat": strings.Repeat,
			"sub": func(from int, is ...int) int {
				for _, i := range is {
					from = from - i
				}
				return from
			},
		}).ParseFiles("embedded/generator/embedded.gotmpl")).Execute(embedded, supData); err != nil {
			return fmt.Errorf("Failed to write to embedded.go: %w", err)
		}

		return nil
	}()

	if err != nil {
		os.Remove("embedded/embedded.go")
		log.Fatal(err)
	}
}
