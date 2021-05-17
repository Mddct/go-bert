## 
```bash
 go run prepare_dict.go  -raw_lexicon big_dict.txt -unit ../voca.txt > lexicon.txt
 grep -v  unk  lexicon.txt lexicon_unk.txt
```
