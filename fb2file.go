// Package fb2parse contains tools for parsing fb2-files
package fb2parse

import (
	"encoding/xml"
	"errors"
	"io"
)

// FB2File struct of fb2 file.
// http://www.fictionbook.org/index.php/Элемент_FictionBook
// http://www.fictionbook.org/index.php/Описание_формата_FB2_от_Sclex
type FB2File struct {
	Description FB2Description `xml:"description"`
	// Body        []FB2Body      `xml:"body"`
	// Binary      []FB2Binary    `xml:"binary"`
}

// NewFB2File factory for FB2File.
func NewFB2File(doc *xml.Decoder) (res FB2File, err error) {
	var token xml.Token

	for {
		if token, err = doc.Token(); err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
			}

			break
		}

		switch typedToken := token.(type) {
		case xml.StartElement:
			switch typedToken.Name.Local { //nolint:gocritic
			case "description":
				res.Description, err = NewFB2Description(typedToken.Name.Local, doc)
			}

			if err != nil {
				break
			}
		}
	}

	return
}
