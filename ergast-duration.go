package ergast

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

type ergastDuration struct {
	time.Duration
}

func (e *ergastDuration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string

	d.DecodeElement(&v, &start)

	p, err := parseErgastDuration(v)
	if err != nil {
		return err
	}

	*e = ergastDuration{p}
	return nil
}

func parseErgastDuration(v string) (time.Duration, error) {
	parts := strings.Split(v, ":")

	// Minutes will be parts[0]

	m := parts[0]

	parts = strings.Split(parts[1], ".")

	// Seconds will be parts[0], Microseconds will be parts[1]

	s := parts[0]
	ms := parts[1]

	return time.ParseDuration(fmt.Sprintf("%vm%vs%vms", m, s, ms))
}
