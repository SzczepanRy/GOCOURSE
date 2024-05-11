package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type xkcd struct {
	Num        int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func main() {
	if len(os.Args) < 2 {

		fmt.Fprintln(os.Stderr, "no file givven ")
		os.Exit(-1)
	}

	fn := os.Args[1]
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "no search item")
		os.Exit(-1)
	}
	fmt.Println(fn)
	var (
		items []xkcd
		terms []string
		input io.ReadCloser
		cnt   int
		err   error
	)

	input, err = os.Open(fn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bad filr")
		os.Exit(-1)

	} //decode file

    //:bytes, _ := input.Read(input)

	if err = json.NewDecoder(input).Decode(&items); err != nil {
		fmt.Fprintf(os.Stderr, "bad file %s  \n",err)
		os.Exit(-1)

	}
	fmt.Println(items)
	fmt.Fprintf(os.Stdout, "read %d comics\n", len(items))
	//search params
	for _, t := range os.Args[2:] {
		terms = append(terms, strings.ToLower(t))
	}

	//search
outer:
	for _, item := range items {
		title := strings.ToLower(item.Title)
		transcript := strings.ToLower(item.Transcript)

		for _, term := range terms {
			if !strings.Contains(title, term) && !strings.Contains(transcript, term) {
				continue outer
			}
		}
		fmt.Printf("https://xkcd.com/%d/ %s/%s/%s %s\n", item.Num, item.Month, item.Day, item.Year, item.Title)
		cnt++
	}

	fmt.Fprintf(os.Stdout, "found %d comicks\n", cnt)

}
