package fb2parse

import (
	"encoding/xml"
	"errors"
	"io"
)

// FB2Author struct of fb2 author.
// http://www.fictionbook.org/index.php/Элемент_author
type FB2Author struct {
	FirstName  string   `xml:"first-name"`
	MiddleName string   `xml:"middle-name"`
	LastName   string   `xml:"last-name"`
	Nickname   string   `xml:"nickname"`
	HomePage   []string `xml:"home-page"`
	Email      []string `xml:"email"`
	ID         string   `xml:"id"`
}

// NewFB2Author factory for FB2Author.
func NewFB2Author(tokenName string, reader xml.TokenReader) (res FB2Author, err error) { //nolint:gocognit,cyclop
	var token xml.Token

	var strVal string

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
			case "first-name":
				res.FirstName, err = GetContent(typedToken.Name.Local, reader)
			case "middle-name":
				res.MiddleName, err = GetContent(typedToken.Name.Local, reader)
			case "last-name":
				res.LastName, err = GetContent(typedToken.Name.Local, reader)
			case "nickname":
				res.Nickname, err = GetContent(typedToken.Name.Local, reader)
			case "home-page":
				if strVal, err = GetContent(typedToken.Name.Local, reader); err == nil && strVal != "" {
					res.HomePage = append(res.HomePage, strVal)
				}
			case "email":
				if strVal, err = GetContent(typedToken.Name.Local, reader); err == nil && strVal != "" {
					res.Email = append(res.Email, strVal)
				}
			case "id":
				res.ID, err = GetContent(typedToken.Name.Local, reader)
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
