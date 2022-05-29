package fb2parse

import (
	"encoding/xml"
	"errors"
	"io"
)

// FB2Description struct of fb2 description.
// http://www.fictionbook.org/index.php/Элемент_description
type FB2Description struct {
	TitleInfo    FB2TitleInfo    `xml:"title-info"`
	SrcTitleInfo *FB2TitleInfo   `xml:"src-title-info"`
	DocInfo      FB2DocInfo      `xml:"document-info"`
	PublishInfo  *FB2Publisher   `xml:"publish-info"`
	CustomInfo   []FB2CustomInfo `xml:"custom-info"`
}

// NewFB2Description factory for FB2Description.
func NewFB2Description(
	tokenName string, reader xml.TokenReader, rules []HandlingRule,
) (res FB2Description, err error) {
	var token xml.Token

	handler := buildChain(rules, getFB2DescriptionHandler(rules))

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
func getFB2DescriptionHandler(rules []HandlingRule) TokenHandler {
	var customInfo FB2CustomInfo

	return func(res interface{}, node xml.StartElement, reader xml.TokenReader) (err error) {
		switch node.Name.Local {
		case "title-info":
			res.(*FB2Description).TitleInfo, err = NewFB2TitleInfo(node.Name.Local, reader, rules)
		case "src-title-info":
			var item FB2TitleInfo

			if item, err = NewFB2TitleInfo(node.Name.Local, reader, rules); err == nil {
				res.(*FB2Description).SrcTitleInfo = &item
			}
		case "document-info":
			res.(*FB2Description).DocInfo, err = NewFB2DocInfo(node.Name.Local, reader, rules)
		case "publish-info":
			var item FB2Publisher

			if item, err = NewFB2Publisher(node.Name.Local, reader, rules); err == nil {
				res.(*FB2Description).PublishInfo = &item
			}
		case "custom-info":
			if customInfo, err = NewFB2CustomInfo(node, reader); err == nil {
				res.(*FB2Description).CustomInfo = append(res.(*FB2Description).CustomInfo, customInfo)
			}
		}

		return
	}
}
