package main

import (
	"fmt"
	"github.com/gfphoenix78/goc/lexer"
	"github.com/gfphoenix78/goc/token"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("lexer_input.txt")
	if err != nil {
		panic("input file not exists")
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic("read failed")
	}
	fmt.Println("data=", string(bytes), bytes)
	l := lexer.NewLexer(bytes)
	for tk := l.Next(); tk.Kind != token.EOF; tk = l.Next() {
		fmt.Println(tk)
	}
}
