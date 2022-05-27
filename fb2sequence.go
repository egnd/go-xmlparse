package fb2parse

import (
	"encoding/xml"
	"strconv"
)

// FB2Sequence struct of fb2 sequence info.
// http://www.fictionbook.org/index.php/Элемент_sequence
type FB2Sequence struct {
	Number int    `xml:"number,attr"`
	Name   string `xml:"name,attr"`
}

// NewFB2Sequence factory for FB2Sequence.
func NewFB2Sequence(token xml.StartElement) (res FB2Sequence, err error) {
	for _, attr := range token.Attr {
		switch attr.Name.Local {
		case "name":
			res.Name = attr.Value
		case "number":
			if res.Number, err = strconv.Atoi(attr.Value); err != nil {
				return
			}
		}
	}

	return
}