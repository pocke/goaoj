package goaoj

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type iint int

func (i *iint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	res, err := strconv.Atoi(strings.Trim(s, "\n"))
	if err != nil {
		return err
	}
	*i = iint(res)
	return nil
}

type sstring string

func (s *sstring) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ss string
	err := d.DecodeElement(&ss, &start)
	if err != nil {
		return err
	}
	res := strings.Trim(ss, "\n")
	*s = sstring(res)
	return nil
}
