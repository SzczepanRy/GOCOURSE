package main 

import (
	"fmt"
	"bufio"
	"sort"
	"os"
)
func main(){
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan(){
		words[scan.Text()]++
	}
	fmt.Println(len(words),"unique words")
//go run main.go < test.txt	


	type kv struct {
	
	key string
	val int

	}

	var ss[]kv

	for k,v := range words{
		ss = append(ss,kv{k,v})


	}

	sort.Slice(ss,func(i,j int) bool {
		return ss[i].val>ss[j].val
	})


	for _,s:= range ss[:10]{
		fmt.Println(s.key,"appears",s.val,"times")
	}	
}
