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
func NewFB2Publisher(
	tokenName string, reader xml.TokenReader, rules []HandlingRule,
) (res FB2Publisher, err error) {
	var token xml.Token

	handler := buildChain(rules, getFB2PublisherHandler(rules))

loop:
	for {
		if token, err = reader.Token(); err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
			}

			break
		}

		switch typedToken := token.(type) {
		case xml.StartElement:
			if err = handler(&res, typedToken, reader); err != nil {
				break loop
			}
		case xml.EndElement:
			if typedToken.Name.Local == tokenName {
				break loop
			}
		}
	}

	return res, err
}

//nolint:forcetypeassert
func getFB2PublisherHandler(_ []HandlingRule) TokenHandler {
	var seq FB2Sequence

	return func(res interface{}, node xml.StartElement, reader xml.TokenReader) (err error) {
		switch node.Name.Local {
		case "book-name":
			res.(*FB2Publisher).BookName, err = GetContent(node.Name.Local, reader)
		case "publisher":
			res.(*FB2Publisher).Publisher, err = GetContent(node.Name.Local, reader)
		case "city":
			res.(*FB2Publisher).City, err = GetContent(node.Name.Local, reader)
		case "year":
			res.(*FB2Publisher).Year, err = GetContent(node.Name.Local, reader)
		case "isbn":
			res.(*FB2Publisher).ISBN, err = GetContent(node.Name.Local, reader)
		case "sequence":
			if seq, err = NewFB2Sequence(node); err == nil {
				res.(*FB2Publisher).Sequence = append(res.(*FB2Publisher).Sequence, seq)
			}
		}

		return
	}
}
