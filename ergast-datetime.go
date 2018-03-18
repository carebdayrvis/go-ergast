package ergast

import (
	"encoding/xml"
	"time"
)

type ErgastDate struct {
	time.Time
}

// Ripped basically whole cloth from https://stackoverflow.com/questions/17301149/golang-xml-unmarshal-and-time-time-fields
func (e *ErgastDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const format = "2006-01-02"
	var v string

	d.DecodeElement(&v, &start)

	parse, err := time.Parse(format, v)
	if err != nil {
		return err
	}

	*e = ErgastDate{parse}
	return nil
}

type ErgastTime struct {
	time.Time
}

func (e *ErgastTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const format = "15:04:05Z"
	var v string

	d.DecodeElement(&v, &start)

	parse, err := time.Parse(format, v)
	if err != nil {
		return err
	}

	*e = ErgastTime{parse}
	return nil
}
