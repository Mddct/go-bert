package main

import (
	"strings"
	"testing"
)

func TestWenetTokenzie(t *testing.T){
	voca := newVocab("voca.txt")
	wenettokenize := newwenetTokenize(voca)

	expected := "_good _morning bug _good 你 好 吗"
	toks := wenettokenize.Tokenize("good morningbug good你好吗")
	got := strings.Join(toks, " ")
	if got != expected{
		t.Logf("expect: %s, but got %s", expected, got)
	}

}