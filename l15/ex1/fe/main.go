package main

import (

    "encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type todo struct{
    ID int `json:"id"`
    Title string `json:"title"`
    Completed bool `json:"completed"`
}

const url = "https://jsonplaceholder.typicode.com"

func main() {

    resp, err := http.Get(url + "/todos/1")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)

			os.Exit(-1)
		}

        var item todo

        err = json.Unmarshal(body,&item)
        if err != nil{
            fmt.Fprintln(os.Stderr,err)
            os.Exit(-1)
        }
        fmt.Printf("%#v\n",item)
    }
}
