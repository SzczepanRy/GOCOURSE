package main

import "fmt"

func main(){
	//arrays 
//	var c =[...]int{0 ,0 ,0}	//sized by initializer
     	var a =[3]int{1,2,3}

//	a=c//elements coppied

	//slices 

//	var a =[]int //nil , no storage
	var b =[3]int{2,2,8}//initialized

		//appending ofc

	a = b //overwirtes a (a is now a pointer to the b values) 
	
	d := make([]int,5)//[]int{0,0,0,0,0}
	e := a // same storage (alias) points to the same storage
		
        fmt.Println("aaa",a,b,d,e)	
        fmt.Println("###################################")	

	
	t:= []byte("string")
	fmt.Println(len(t),t)
	fmt.Println(t[2])
	fmt.Println(t[:2])
	fmt.Println(t[2:])
	fmt.Println(t[3:5],len(t[3:5]))


	//maps

        fmt.Println("###################################")	
	var m map[string]int //nil , no storage
	p := make(map[string]int)//non-nil but empty

	//m["and"] = 1 //panic - nil map
	p["and"] = 1 //ok
	m = p//m now points to the same thing that p points to
	m["and"]++

	c := p["and"]
	fmt.Println(c)
	fmt.Println(p["the"])
	fmt.Println(m["the"])

	//f:= map[string]int{}
	//a :=p["the"]

	//t,ok := p["and"]
	//fmt.Println(t,ok)

	p["the"]++
	l,ok := p["the"]
	fmt.Println(l,ok)


	if w ,ok := p["the"];ok{
		fmt.Println("the is in the map",w)
	}




}
