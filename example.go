package main

import (
	"bufio"
	"fmt"
	"go-bert/tokenize"
	"go-bert/tokenize/vocab"
	"os"
	"strings"
	"unicode"
)

type wenetTokenize struct {
	*tokenize.Full
}
func newwenetTokenize(dict vocab.Dict) tokenize.VocabTokenizer{
	return wenetTokenize{
		&tokenize.Full{
			Basic: tokenize.NewBasic(),
			Wordpiece: tokenize.NewWordpiece(dict),
		},
	}
}

func (f wenetTokenize) Tokenize(text string) []string {
	var toks []string
	for _, tok := range f.Basic.Tokenize(text) {
		for _ , wordpiece := range f.Wordpiece.Tokenize(tok){
			if strings.HasPrefix(wordpiece,"##"){
				wordpiece = wordpiece[len("##"):]
			}else {
				isHan := true
				for _, v := range wordpiece{
					if !unicode.Is(unicode.Han,v){
						isHan = false
					}
				}
				if !isHan{
					wordpiece = "_" + wordpiece
				}
			}
			toks = append(toks, wordpiece)
		}

	}
	return toks
}


func newVocab(filename string) vocab.Dict{
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	table := []string{}
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		table = append(table, scan.Text())
	}
	return vocab.New(table)
}

func main() {
	voca := newVocab("voca.txt")
	wenettokenize := newwenetTokenize(voca)
	toks := wenettokenize.Tokenize("good morningbug good你好吗")
	fmt.Println(strings.Join(toks, " "))
}