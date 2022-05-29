package fb2parse

import (
	"encoding/xml"
	"errors"
	"io"
)

// FB2TitleInfo struct of fb2 title info.
// http://www.fictionbook.org/index.php/Элемент_title-info
type FB2TitleInfo struct {
	Genre      []string       `xml:"genre"`
	Author     []FB2Author    `xml:"author"`
	BookTitle  string         `xml:"book-title"`
	Annotation *FB2Annotation `xml:"annotation"`
	Keywords   string         `xml:"keywords"`
	Date       string         `xml:"date"`
	Coverpage  *FB2Cover      `xml:"coverpage"`
	Lang       string         `xml:"lang"`
	SrcLang    string         `xml:"src-lang"`
	Translator []FB2Author    `xml:"translator"`
	Sequence   []FB2Sequence  `xml:"sequence"`
}

// NewFB2TitleInfo factory for FB2TitleInfo.
func NewFB2TitleInfo(
	tokenName string, reader xml.TokenReader, rules []HandlingRule,
) (res FB2TitleInfo, err error) {
	var token xml.Token

	handler := buildChain(rules, getFB2TitleInfoHandler(rules))

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
func getFB2TitleInfoHandler(_ []HandlingRule) TokenHandler { //nolint:cyclop
	var strVal string

	var author FB2Author

	var seq FB2Sequence

	return func(res interface{}, node xml.StartElement, reader xml.TokenReader) (err error) {
		switch node.Name.Local {
		case "genre":
			if strVal, err = GetContent(node.Name.Local, reader); err == nil && strVal != "" {
				res.(*FB2TitleInfo).Genre = append(res.(*FB2TitleInfo).Genre, strVal)
			}
		case "author":
			if author, err = NewFB2Author(node.Name.Local, reader); err == nil {
				res.(*FB2TitleInfo).Author = append(res.(*FB2TitleInfo).Author, author)
			}
		case "book-title":
			res.(*FB2TitleInfo).BookTitle, err = GetContent(node.Name.Local, reader)
		case "annotation":
			var annotation FB2Annotation
			if annotation, err = NewFB2Annotation(node.Name.Local, reader); err == nil {
				res.(*FB2TitleInfo).Annotation = &annotation
			}
		case "keywords":
			res.(*FB2TitleInfo).Keywords, err = GetContent(node.Name.Local, reader)
		case "date":
			res.(*FB2TitleInfo).Date, err = GetContent(node.Name.Local, reader)
		case "coverpage":
			var cover FB2Cover
			if cover, err = NewFB2Cover(node.Name.Local, reader); err == nil {
				res.(*FB2TitleInfo).Coverpage = &cover
			}
		case "lang":
			res.(*FB2TitleInfo).Lang, err = GetContent(node.Name.Local, reader)
		case "src-lang":
			res.(*FB2TitleInfo).SrcLang, err = GetContent(node.Name.Local, reader)
		case "translator":
			if author, err = NewFB2Author(node.Name.Local, reader); err == nil {
				res.(*FB2TitleInfo).Translator = append(res.(*FB2TitleInfo).Translator, author)
			}
		case "sequence":
			if seq, err = NewFB2Sequence(node); err == nil {
				res.(*FB2TitleInfo).Sequence = append(res.(*FB2TitleInfo).Sequence, seq)
			}
		}

		return
	}
}
