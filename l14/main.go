package main
import "fmt"

func main(){
//refrences and simanticks
//pointers are shared but not coppied
//values are the opposute

// if a value cant be coppied than we share it 
// if a value is less than 64bytes consider cpooying thevale and returning it
//we sould use pointers or values consistantly !
//the stack is the more optimal way to refrence data  ... its faster than the heap


//returened range is always a copy
for i, thing := range []int{1,2,3,4,5,6}{
    fmt.Println(i,thing)
    //thing is always a coppy
}
// use index if you need to mutate element 
var slice = []int{1,2,3,4,5,6}
for i := range slice{
    slice[i] = 12
}
fmt.Println(slice)

//always append to the same verable



//getting pointer slices by index is risky








}
