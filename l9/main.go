package main

import "fmt"


func fib () func () int{

	a, b:= 0, 1
	return func()int{
		a, b = b, a+b
		return b

	}
	
}

func main(){

	//closures
	//verable scome vs lifetime
	//escape ananlisis
	//in go  lifetime of the verable can exede the scope of it
	//a closure involves a function using verables from another scope
	//

	f:= fib()
	for x:=f(); x <100; x=f(){
		fmt.Println(x)
	}


	fmt.Println("###################################")

	g,h := fib(),fib()
	//notice the diffrent env pointer of the a & b values
	fmt.Println(g(),g(),g(),g())

	fmt.Println(h(),h(),h(),h())

}
