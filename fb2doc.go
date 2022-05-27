package fb2parse

import (
	"encoding/xml"
	"errors"
	"io"
)

// FB2DocInfo struct of fb2 document info.
// http://www.fictionbook.org/index.php/Элемент_document-info
type FB2DocInfo struct {
	Authors []FB2Author `xml:"author"`
	// program-used - 0..1 (один, опционально) @TODO:
	// date - 1 (один, обязателен) @TODO:
	SrcURL []string `xml:"src-url"`
	// src-ocr - 0..1 (один, опционально) @TODO:
	ID      string `xml:"id"`
	Version string `xml:"version"`
	// history - 0..1 (один, опционально) @TODO:
	Publishers []FB2Author `xml:"publisher"`
}

// NewFB2DocInfo factory for NewFB2DocInfo.
func NewFB2DocInfo(tokenName string, reader xml.TokenReader) (res FB2DocInfo, err error) { //nolint:gocognit,cyclop
	var token xml.Token

	var strVal string

	var author FB2Author

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
			case "author":
				if author, err = NewFB2Author(typedToken.Name.Local, reader); err == nil {
					res.Authors = append(res.Authors, author)
				}
			case "src-url":
				if strVal, err = GetContent(typedToken.Name.Local, reader); err == nil && strVal != "" {
					res.SrcURL = append(res.SrcURL, strVal)
				}
			case "id":
				res.ID, err = GetContent(typedToken.Name.Local, reader)
			case "version":
				res.Version, err = GetContent(typedToken.Name.Local, reader)
			case "publisher":
				if author, err = NewFB2Author(typedToken.Name.Local, reader); err == nil {
					res.Publishers = append(res.Publishers, author)
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
