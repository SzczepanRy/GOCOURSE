package main
import "fmt"

func main(){
	//slices in depth
	
	var s []int
	fmt.Println(len(s), cap(s),s==nil)
	t:= []int{}
	fmt.Println(len(t), cap(t),t==nil)
	u := make([]int,5)
	fmt.Println(len(u), cap(u),u==nil)
	v:= make([]int,0 ,5)
	fmt.Println(len(v), cap(v),v==nil)

	
	//can append to a nil slice i n contrast to maps

	a:= [3]int{1,2,3}
	b:= a[0:1]	
	fmt.Println("a = ",a) 

	fmt.Println("b = ",b) 
	
	c := b[0:2]//???
	//slices the cap of b witch is tch cap a whitch is filled

	fmt.Println("c = ",c) 

	fmt.Println("len c ",len(c)) 
	fmt.Println("cap c ",cap(c)) 
	fmt.Println("len b ",len(b)) 
	fmt.Println("cap b ",cap(b)) 

	//d := [0:1:1]//controls length and cap to prevent the example above
	//fmt.Println(d)
	//appending to a slice can mutate the underelying array
	//when not controling the cap  of the slice ofc	
}
