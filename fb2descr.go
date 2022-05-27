package fb2parse

import (
	"encoding/xml"
	"errors"
	"io"
)

// FB2Description struct of fb2 description.
// http://www.fictionbook.org/index.php/Элемент_description
type FB2Description struct {
	TitleInfo    FB2TitleInfo  `xml:"title-info"`
	SrcTitleInfo *FB2TitleInfo `xml:"src-title-info"`
	// DocInfo      FB2DocInfo      `xml:"document-info"`
	PublishInfo *FB2Publisher `xml:"publish-info"`
	// CustomInfo   []FB2CustomInfo `xml:"custom-info"`
}

// NewFB2Description factory for FB2Description.
func NewFB2Description(tokenName string, reader xml.TokenReader) (res FB2Description, err error) { //nolint:cyclop
	var token xml.Token

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
			case "title-info":
				res.TitleInfo, err = NewFB2TitleInfo(typedToken.Name.Local, reader)
			case "src-title-info":
				var item FB2TitleInfo

				if item, err = NewFB2TitleInfo(typedToken.Name.Local, reader); err == nil {
					res.SrcTitleInfo = &item
				}
			case "publish-info":
				var item FB2Publisher

				if item, err = NewFB2Publisher(typedToken.Name.Local, reader); err == nil {
					res.PublishInfo = &item
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
