package fb2parse_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/egnd/go-fb2parse"
	"github.com/stretchr/testify/assert"
)

func Test_Parsing(t *testing.T) {
	fb2data, err := getBookData("examples/small.xml")
	assert.NoError(t, err)

	var marshRes fb2parse.FB2File
	err = fb2parse.NewDecoder(bytes.NewReader(fb2data)).Decode(&marshRes)
	assert.NoError(t, err)

	parseRes, err := fb2parse.NewFB2File(fb2parse.NewDecoder(bytes.NewReader(fb2data)))
	assert.NoError(t, err)

	assert.EqualValues(t, marshRes, parseRes)
}

func getBookData(path string) ([]byte, error) {
	book, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer book.Close()

	return io.ReadAll(book)
}
