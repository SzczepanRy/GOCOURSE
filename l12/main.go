package main 

import (
    "encoding/json"
    "fmt"
    "time"
)






type Employee struct {
	Name string
	Number int
	Boss *Employee
	Hired  time.Time
}


type Response struct{
    Page int `json:"page"`
    Words []string `json:"words,omitempty"`//if wors are nil than words do not return
}
// if words  ar not capped than it funcks up the encoding
func main (){
	// structs, struct tags , JSON 

	
    e :=  Employee{
        Name: "tom",
        Number:12,
        Hired:time.Now(),
    }

    
    b:= Employee{Name:"jim",Boss:&e}


    e.Name = "mat"
    e.Hired = time.Now()

    
	fmt.Printf("%T %+[1]v\n" , e)
	fmt.Printf("%T %+[1]v\n" , b)



    
    c:= map[string]*Employee{}

    c["mat"]= &Employee{Name:"mat",Number:13}
    c["mat"].Number++
    c["jake"]= &Employee{Name:"jake",Boss:c["mat"],Number:133}
    //cant do that  c["jake"]= &Employee{Name:"jake",Boss:&c["mat"],Number:133}
    fmt.Println(c["jake"])
    fmt.Println(c["mat"])



   ///maps canit take a value od the entry of a struct
   //

    fmt.Println("$$#################################")

    var album1 = struct{
        title string
    }{
        "the goofd album",
    }

    var album2 = struct{
        title string
    }{
        "the bad album",
    }

    album1 = album2 //copy values + structular comp
    fmt.Println(album1,album2)
    //cannot work when types are named types
    //if using named type s convert them as usual 
    //ex : structType album1 = album1(someStruct)

    //structs are comparible when their values are 



    ///struct tags are ex `json:"foo"`
    

    // how we do json

    r := &Response{Page:1 , Words:[]string{"up","in","out"}}

    fmt.Println("%#v\n",r)

    j,_ := json.Marshal(r)
    fmt.Printf(string(j))

    var r2 Response

    _ = json.Unmarshal(j , &r2)
    


    fmt.Printf("%#v\n",r2)





    














}
