package fb2parse

import (
	"encoding/xml"
	"errors"
	"io"
)

// FB2Publisher struct of fb2 publisher info.
// http://www.fictionbook.org/index.php/Элемент_publish-info
type FB2Publisher struct {
	BookName  string        `xml:"book-name"`
	Publisher string        `xml:"publisher"`
	City      string        `xml:"city"`
	Year      string        `xml:"year"`
	ISBN      string        `xml:"isbn"`
	Sequence  []FB2Sequence `xml:"sequence"`
}

// NewFB2Publisher factory for FB2Publisher.
func NewFB2Publisher(tokenName string, reader xml.TokenReader) (res FB2Publisher, err error) { //nolint:cyclop
	var token xml.Token

	var seq FB2Sequence

	for {
		if token, err = reader.Token(); err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
			}

			break
		}

		switch typedToken := token.(type) {
		case xml.StartElement:
			switch typedToken.Name.Local {
			case "book-name":
				res.BookName, err = GetContent(typedToken.Name.Local, reader)
			case "publisher":
				res.Publisher, err = GetContent(typedToken.Name.Local, reader)
			case "city":
				res.City, err = GetContent(typedToken.Name.Local, reader)
			case "year":
				res.Year, err = GetContent(typedToken.Name.Local, reader)
			case "isbn":
				res.ISBN, err = GetContent(typedToken.Name.Local, reader)
			case "sequence":
				if seq, err = NewFB2Sequence(typedToken); err == nil {
					res.Sequence = append(res.Sequence, seq)
				}
			}

			if err != nil {
				break
			}
		case xml.EndElement:
			if typedToken.Name.Local == tokenName {
				break
			}
		}
	}

	return res, err
}
