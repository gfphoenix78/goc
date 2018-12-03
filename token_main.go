package main

import (
	"fmt"
)

import "gocc/token"

func main()  {
	for i:=0; i<=int(token.EOF); i++ {
		tk := token.TokenKind(i)
		fmt.Println(i, tk)
	}
}
