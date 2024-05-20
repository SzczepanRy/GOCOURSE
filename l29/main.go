package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float32 // Bad

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database struct {
	mu sync.Mutex
	db map[string]dollars
}

func (d *database) list(w http.ResponseWriter, req *http.Request) {
	//dont need pointer refrende because db already points to a hash table
	d.mu.Lock()
	defer d.mu.Unlock()
	for item, price := range d.db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (d *database) update(w http.ResponseWriter, req *http.Request) {
	d.mu.Lock()
	defer d.mu.Unlock()
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := d.db[item]; !ok {
		msg := fmt.Sprintf("no sutch item  item :%q", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("duplicate item :%q", price)
		http.Error(w, msg, http.StatusBadRequest)
		return

	}

	d.db[item] = dollars(p)

	fmt.Fprintf(w, "updated sucesfully { name : %s , price: %s}", item, d.db[item])

}

func (d *database) add(w http.ResponseWriter, req *http.Request) {
	d.mu.Lock()
	defer d.mu.Unlock()

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := d.db[item]; ok {
		msg := fmt.Sprintf("duplicate item :%q", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("duplicate item :%q", price)
		http.Error(w, msg, http.StatusBadRequest)
		return

	}

	d.db[item] = dollars(p)

	fmt.Fprintf(w, "added sucesfully { name : %s , price: %s}", item, d.db[item])

}

func (d *database) fetch(w http.ResponseWriter, req *http.Request) {
	d.mu.Lock()
	defer d.mu.Unlock()

	item := req.URL.Query().Get("item")

	if _, ok := d.db[item]; !ok {
		msg := fmt.Sprintf("no sutch item  item :%q", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "item %s has price %s\n", item, d.db[item])
}

func (d *database) drop(w http.ResponseWriter, req *http.Request) {
	d.mu.Lock()
	defer d.mu.Unlock()

	item := req.URL.Query().Get("item")

	if _, ok := d.db[item]; !ok {
		msg := fmt.Sprintf("no sutch item  item :%q", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	////db is the wraper around the map thats why we use delete
	delete(d.db, item)

	fmt.Fprintf(w, "deleted  %s\n", item)
}

func main() {
	db := database{
		db:map[string]dollars {
			"a": 21,
			"b": 3,
		},
	}

	http.HandleFunc("/list", db.list)

	http.HandleFunc("/create", db.add)

	http.HandleFunc("/update", db.update)
	http.HandleFunc("/read", db.fetch)
	http.HandleFunc("/delete", db.drop)
	log.Fatal(http.ListenAndServe(":3000", nil))

}

/*
oop in http




*/
