package fb2_test

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/egnd/go-xmlparse/fb2"
	"github.com/egnd/go-xmlparse/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_Author(t *testing.T) {
	cases := []struct {
		err   error
		mocks func(*mocks.TokenReader)
	}{
		{
			err: io.EOF,
			mocks: func(r *mocks.TokenReader) {
				r.On("Token").Return(nil, io.EOF)
			},
		},
		{
			err: errors.New("error"),
			mocks: func(r *mocks.TokenReader) {
				r.On("Token").Return(nil, errors.New("error"))
			},
		},
		{
			err: errors.New("error"),
			mocks: func(r *mocks.TokenReader) {
				r.On("Token").Return(xml.StartElement{Name: xml.Name{Local: "first-name"}}, nil).Once()
				r.On("Token").Return(nil, errors.New("error")).Once()
			},
		},
		{
			err: io.EOF,
			mocks: func(r *mocks.TokenReader) {
				r.On("Token").Return(xml.StartElement{Name: xml.Name{Local: "first-name"}}, nil).Once()
				r.On("Token").Return(nil, io.EOF).Once()
			},
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			reader := &mocks.TokenReader{}
			test.mocks(reader)
			_, err := fb2.NewAuthor("asdfds", reader)
			assert.EqualValues(t, test.err, err)
			reader.AssertExpectations(t)
		})
	}
}
