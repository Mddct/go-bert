package tokenize

import (
	"go-bert/tokenize/vocab"
	"strings"
	"unicode"
	"os"
	"bufio"
)

type wenetTokenize struct {
	*Full
}
func newwenetTokenize(dict vocab.Dict) VocabTokenizer{
	return wenetTokenize{
		&Full{
			Basic: NewBasic(),
			Wordpiece: NewWordpiece(dict),
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
		word := scan.Text()
		segs := strings.Split(word, " ")
		if len(segs) < 1{
			continue
		}
		isHan := true
		for _ , r := range segs[0]{
			if !unicode.Is(unicode.Han ,r){
				isHan = false
				break
			}
		}
		if !isHan{
			if strings.HasPrefix(segs[0], "_"){
				segs[0] = segs[0][1:]
			}else{
				segs[0] = "##" + segs[0]
			}
		}
		table = append(table, segs[0])
	}
	return vocab.New(table)
}

func NewWenetTokenize(filename string) VocabTokenizer{
	voca := newVocab(filename)
	return newwenetTokenize(voca)
}
