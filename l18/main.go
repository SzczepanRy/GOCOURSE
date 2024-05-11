package main

import (
    "fmt"
    "strings"
    "strconv"
)
type IntSlice []int

//method for Int Slice
func (is IntSlice) String() string{
    var strs []string


    for _,v := range is{
        strs = append(strs,strconv.Itoa(v))

    }
    return "["+strings.Join(strs, ";")+"]"
}



func main (){
    var v IntSlice = []int{1,2,3,4,5}
    for i,x := range v {
        fmt.Printf("%d: %d\n",i,x)
    }
    fmt.Printf("%d: %[1]v\n" , v)
}

/*

without  interfaces we would have to write many functions tor many concerte types

we can define our funcrions in terms of abstract behavour


an interface sepecifies req behavior as a method set

we dont have to have a struct to have methods

canot put methods on int and sting and bool
you can wrap them via your own type



methoods have to be inside the same filr a the type


*/
