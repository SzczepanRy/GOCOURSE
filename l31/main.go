package main

import "fmt"

//enumerated types
//iota operator

//func sum(x ...int , <- this wll not work nums ...int) int{
func sum(nums ...int) int{
    var total int

    for _,num := range nums {
        total += num
    }
    return total
}

func main(){
    fmt.Println(sum())
    fmt.Println(sum(1))
    fmt.Println(sum(1,2,3,4))

    s := []int{1,2,3,4}
    //fmt.Println(sum(s))//does not wirk

    fmt.Println(sum(s...))//will work
    // s = append(s , s) // will not work

    s = append(s , s... ) // will  work

    fmt.Println(sum(s...))//will work

    //shift operator << shifts the bits operator a << 3 shift by 3 b >> 2 unshift by 2
}

