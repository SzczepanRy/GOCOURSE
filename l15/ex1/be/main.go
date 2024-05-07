package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const url = "https://jsonplaceholder.typicode.com/"

var form = `
    <h1>Todo #{{.ID}}</h1>
    <div>{{printf "%s (completed: %t)" .Title .Completed}}</div>`

func handler(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get(url + r.URL.Path[1:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {



        var item todo

		//err = json.Unmarshal(body, &item)
		err = json.NewDecoder(resp.Body).Decode(&item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

        tmpl := template.New("mine")
        tmpl.Parse(form)
        tmpl.Execute(w,item)


    }
	fmt.Fprintf(w, " Hello from %s\n", r.URL.Path[1:])
}

func main() {
	//http
	http.HandleFunc("/", handler)

	//opens a tcp/http connection
	log.Fatal(http.ListenAndServe(":3000", nil))
}
