package main

import (
	"fmt"
	"go-bert/tokenize"
)

func main(){
	tokn := tokenize.NewWenetTokenize("tools/dict_number.txt")
	fmt.Println(tokn.Tokenize("你好 中国bipkid fuck"))
}