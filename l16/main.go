package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getOne(i int) []byte {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "stopped : %s\n", err)

		os.Exit(-1)

	}
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK{
        fmt.Println("skipping" , i)
        return nil
    }
    body , err := io.ReadAll(resp.Body)
    	if err != nil {
		fmt.Fprintf(os.Stderr, "cant read : %s\n", err)

		os.Exit(-1)

	}


    return body


}

func main() {
	var (
		output io.WriteCloser = os.Stdout
		err    error
		cnt    int
		data   []byte
		fails  int
	)
	if len(os.Args) > 1 {
		output, err = os.Create(os.Args[1])
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(-1)
		}
		defer output.Close()
		//the output will be uin thr form of a array
		//so add the brackets before and after
		fmt.Fprint(output,"[")
		defer fmt.Fprint(output,"]")
		// stop if er get 2 404s in a row
		for i := 1; fails < 2; i++ {
			if data = getOne(i); data == nil {
				fails++
				continue
			}
			if cnt > 0 {
				fmt.Fprint(output, ",")
			}

			_, err = io.Copy(output, bytes.NewBuffer(data))
			if err != nil {
				fmt.Fprintf(os.Stderr, "stopped : %s\n", err)

				os.Exit(-1)

			}
			fails = 0
			cnt++
		}

		fmt.Fprintf(os.Stderr, "reag %d comics\n", cnt)
	}
}
