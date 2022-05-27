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
	// Coverpage     *FB2Cover `xml:"coverpage"`
	Lang       string        `xml:"lang"`
	SrcLang    string        `xml:"src-lang"`
	Translator []FB2Author   `xml:"translator"`
	Sequence   []FB2Sequence `xml:"sequence"`
}

// NewFB2TitleInfo factory for FB2TitleInfo.
func NewFB2TitleInfo( //nolint:funlen,gocognit,cyclop
	tokenName string, reader xml.TokenReader,
) (res FB2TitleInfo, err error) {
	var token xml.Token

	var strVal string

	var author FB2Author

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
			case "genre":
				if strVal, err = GetContent(typedToken.Name.Local, reader); err == nil && strVal != "" {
					res.Genre = append(res.Genre, strVal)
				}
			case "author":
				if author, err = NewFB2Author(typedToken.Name.Local, reader); err == nil {
					res.Author = append(res.Author, author)
				}
			case "book-title":
				res.BookTitle, err = GetContent(typedToken.Name.Local, reader)
			case "annotation":
				var annotation FB2Annotation
				if annotation, err = NewFB2Annotation(typedToken.Name.Local, reader); err == nil {
					res.Annotation = &annotation
				}
			case "keywords":
				res.Keywords, err = GetContent(typedToken.Name.Local, reader)
			case "date":
				res.Date, err = GetContent(typedToken.Name.Local, reader)
			case "lang":
				res.Lang, err = GetContent(typedToken.Name.Local, reader)
			case "src-lang":
				res.SrcLang, err = GetContent(typedToken.Name.Local, reader)
			case "translator":
				if author, err = NewFB2Author(typedToken.Name.Local, reader); err == nil {
					res.Translator = append(res.Translator, author)
				}
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
