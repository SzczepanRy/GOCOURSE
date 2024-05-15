package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)



type dollars float32 // Bad


func (d dollars) String() string {
    return fmt.Sprintf("$%.2f",d )
}

type database map[string]dollars

func(db database) list(w http.ResponseWriter, req *http.Request){
//dont need pointer refrende because db already points to a hash table

    for item, price := range db{
        fmt.Fprintf(w,"%s: %s\n", item,price)
    }
}


func(db database) update(w http.ResponseWriter, req *http.Request){
    item := req.URL.Query().Get("item")
    price := req.URL.Query().Get("price")


    if _, ok := db[item] ;!ok{
        msg:= fmt.Sprintf("no sutch item  item :%q", item)
        http.Error(w,msg, http.StatusNotFound)
        return
    }
    p,err := strconv.ParseFloat(price ,32 );
    if err !=nil{
        msg:= fmt.Sprintf("duplicate item :%q", price)
        http.Error(w,msg, http.StatusBadRequest)
        return

    }


    db[item] = dollars(p )

    fmt.Fprintf(w , "updated sucesfully { name : %s , price: %s}",item , db[item])



}



func(db database) add(w http.ResponseWriter, req *http.Request){
    item := req.URL.Query().Get("item")
    price := req.URL.Query().Get("price")


    if _, ok := db[item] ;ok{
        msg:= fmt.Sprintf("duplicate item :%q", item)
        http.Error(w,msg, http.StatusBadRequest)
        return
    }
    p,err := strconv.ParseFloat(price ,32 );
    if err !=nil{
        msg:= fmt.Sprintf("duplicate item :%q", price)
        http.Error(w,msg, http.StatusBadRequest)
        return

    }


    db[item] = dollars(p )

    fmt.Fprintf(w , "added sucesfully { name : %s , price: %s}",item , db[item])



}


func(db database) fetch(w http.ResponseWriter, req *http.Request){
    item := req.URL.Query().Get("item")

    if _, ok := db[item] ;!ok{
        msg:= fmt.Sprintf("no sutch item  item :%q", item)
        http.Error(w,msg, http.StatusNotFound)
        return
    }
    fmt.Fprintf(w, "item %s has price %s\n" , item , db[item])
}

func(db database) drop(w http.ResponseWriter, req *http.Request){
    item := req.URL.Query().Get("item")

    if _, ok := db[item] ;!ok{
        msg:= fmt.Sprintf("no sutch item  item :%q", item)
        http.Error(w,msg, http.StatusNotFound)
        return
    }

    //db is the wraper around the map thats why we use delete
    delete(db, item)

    fmt.Fprintf(w, "deleted  %s\n" , item )
}





func main(){
    db:= database{
        "a": 21,
        "b":3,
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
