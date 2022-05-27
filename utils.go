package fb2parse

import (
	"bytes"
	"encoding/xml"
	"io"
	"strings"

	"golang.org/x/net/html/charset"
)

// NewDecoder creates decoder for fb2 xml data.
func NewDecoder(reader io.Reader) *xml.Decoder {
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoder.Strict = false
	decoder.AutoClose = xml.HTMLAutoClose
	decoder.Entity = xml.HTMLEntity

	return decoder
}

// Unmarshal converts fb2 xml data some struct.
func Unmarshal(data []byte, v any) error {
	return NewDecoder(bytes.NewBuffer(data)).Decode(v) //nolint:wrapcheck
}

// GetContent returns xml token inner content.
func GetContent(tokenName string, reader xml.TokenReader) (res string, err error) {
	var buf strings.Builder

	var token xml.Token

	for {
		if token, err = reader.Token(); err != nil {
			break
		}

		switch typedToken := token.(type) {
		case xml.CharData:
			buf.Write(typedToken)
		case xml.EndElement:
			if typedToken.Name.Local == tokenName {
				res = strings.TrimSpace(buf.String())

				break
			}
		}
	}

	return
}
