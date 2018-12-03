package parser

import (
	"gocc/lexer"
	"io"
	"io/ioutil"
	"os"
)

type Parser struct {
	l *lexer.BaseLexer
}

func ParseFile(file string) bool {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return ParseReader(f)
}
func ParseReader(r io.Reader) bool {
	source, error := ioutil.ReadAll(r)
	if error!=nil {
		panic(error)
	}
	return ParseBytes(source)
}
func ParseBytes(source []byte) bool {
	p := &Parser{
		l : lexer.NewLexer(source),
	}
	if p!=nil {
		return true
	}
	return false
}
