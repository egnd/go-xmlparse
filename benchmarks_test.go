package fb2parse_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/egnd/go-fb2parse"
)

func Benchmark_Decoding(b *testing.B) {
	fb2data, err := getBookData("examples/big.xml")
	if err != nil {
		b.Error(err)
	}

	b.Run("xml", func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			if err := fb2parse.NewDecoder(bytes.NewBuffer(fb2data)).Decode(&fb2parse.FB2File{}); err != nil {
				bb.Error(err)
			}
		}
	})

	b.Run("fb2", func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			if err := fb2parse.NewDecoder(bytes.NewBuffer(fb2data)).Decode(&fb2parse.FB2File{}); err != nil {
				bb.Error(err)
			}
		}
	})
}

func Benchmark_Parsing(b *testing.B) {
	table := []int{0, 1, 10, 100}

	fb2data, err := getBookData("examples/big.xml")
	if err != nil {
		b.Error(err)
	}

	for _, rulesCnt := range table {

		rules := make([]fb2parse.HandlingRule, 0, rulesCnt)
		for i := 0; i < rulesCnt; i++ {
			rules = append(rules, getBenchParsingRule)
		}

		b.Run(fmt.Sprintf("rules_%d", len(rules)), func(bb *testing.B) {
			for k := 0; k < bb.N; k++ {
				if _, err := fb2parse.NewFB2File(fb2parse.NewDecoder(bytes.NewBuffer(fb2data)), rules...); err != nil {
					bb.Error(err)
				}
			}
		})
	}
}

func getBenchParsingRule(next fb2parse.TokenHandler) fb2parse.TokenHandler {
	return func(obj interface{}, node xml.StartElement, r xml.TokenReader) (err error) {
		if _, ok := obj.(*fb2parse.FB2TitleInfo); ok && node.Name.Local == "keywords" {
			obj.(*fb2parse.FB2TitleInfo).Keywords = "test keyword"
		}

		return next(obj, node, r)
	}
}
