package ergast

import (
	"encoding/xml"
	"time"
)

type ergastDate struct {
	time.Time
}

// Ripped basically whole cloth from https://stackoverflow.com/questions/17301149/golang-xml-unmarshal-and-time-time-fields
func (e *ergastDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const format = "2006-01-02"
	var v string

	d.DecodeElement(&v, &start)

	parse, err := time.Parse(format, v)
	if err != nil {
		return err
	}

	*e = ergastDate{parse}
	return nil
}

type ergastTime struct {
	time.Time
}

func (e *ergastTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const format = "15:04:05Z"
	var v string

	d.DecodeElement(&v, &start)

	parse, err := time.Parse(format, v)
	if err != nil {
		return err
	}

	*e = ergastTime{parse}
	return nil
}
