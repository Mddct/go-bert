package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-bert/tokenize"
	"log"
	"os"
	"strings"
)

var(
	unit = flag.String("unit", "", "e2e model unit" )
	rawLexicon = flag.String("raw_lexicon", "", "raw lexicon eg: biddict")
	//tlgLexicon = flag.String("tlg_lexicon", "", "TLG lexicon")
)

func readOneSeg(filename string)(set map[string]struct{}){
	f, err := os.Open(filename)
	if err != nil{
		log.Fatalf("unit file: %w", err)
	}
	defer f.Close()

	set = make(map[string]struct{})

	scan := bufio.NewScanner(f)
	for scan.Scan(){
		// token id
		line := scan.Text()
		segs := strings.Split(line, " ")

		set[segs[0]] = struct{}{}
	}
	return set
}
func e2eUnitSet(filename string)(unitTable map[string]struct{}){
	return readOneSeg(filename)
}
func toTlGLexicon(rawLexiconFile string,  m map[string]struct{}, bpeVoca string, sep string){
	if len(m ) == 0{
		return
	}
	rf , err:= os.Open(rawLexiconFile)
	if err != nil{
		log.Fatalf("raw lexicon: %w", err)
	}
	defer rf.Close()

	scan := bufio.NewScanner(rf)
	tokn :=  tokenize.NewWenetTokenize(bpeVoca)
	for scan.Scan(){
		line := scan.Text()
		segs := strings.Split(line, sep)
		if len(segs) < 2{
			continue
		}
		if segs[0] == "SIL" || segs[0] == "<SPOKEN_NOISE>"{
			continue
		}

		wordpieces := tokn.Tokenize(segs[0])
		for i := range wordpieces{
			if _, ok := m[wordpieces[i]]; !ok {
				wordpieces[i] = "<unk>"
			}
		}
		fmt.Println(segs[0], strings.Join(wordpieces, " "))
	}
}
func main(){
	//token := tokenize.NewWenetTokenize("../voca.txt")
	//fmt.Println(token.Tokenize("你好中国vipkid english ipad苹果very bad"))
	flag.Parse()
	unitSet := e2eUnitSet(*unit)
	toTlGLexicon(*rawLexicon, unitSet, "../voca.txt", "\t")

}