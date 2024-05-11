package main

import "fmt"


type errFoo struct {
    err error
    path string
}


func (e errFoo) Error() string{
    return fmt.Sprintf("%s: %s", e.path , e.err)
}


func XYZ(a int) error {// bad *errFoo{
    return nil
}


func main(){
    var err error = XYZ(1)// BAD : interface gets a nil concrete ptr
    if err!= nil {
        fmt.Println("oops err ")
    }else{
        fmt.Println("OK")}

}
/*

an iterface has 2 parts
- somway to refare to the concrete type of a value
- what type that value is
can think of it as two pointers

interface -> type,pointer(value) -> object
r -> nil, nil -> r == null
r -> Buffer , nil -> r == not null


p1 := new(Point)//*Point (0,0)
p2:= Point{1,1}

p1.OffsetOf(p2) //same as (*p).OffsetOf (derefrence the pointer)
p2.Add(3,4)//same as (&p2).Add (refreance of pointer)


L=R

var p Point

p.Add(1,2)//OK
Point{1,2}.Add{2,3} // NOT OK


if one method takes a recever type than all of the methods should take it
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

CURRYING - takes a function that has say 3 params  bind one of them and turn it to a 2 param function

func Add(a, b int) int{
    return a+b
}

func AddToA(a int) func(int) int {//binds value to a
    return func(b int) int{
        return Add(a,b)
    }
}

AddTo1 := AddToA(1)

fmt.Println(Add(1,2) == AddTo1(2)) // true
//////////////////////////////////////////////////////////////////////////////////////////////////////////
METHOD VALUE


func (p Point) Distance(q Point) float64{//notise that the method takes a value of p not a pointer whitch means that the p is coppied
    return math.Hypot(p.X-q.X , q.Y-p.Y)
}

p := Point{1,2}
q := Point{4,6}

distanceFromP := p.Distance//this is the method value it closes over p

fmt.Println(distanceFromP(q))//and can be called later
//////////////////////////////////////////////////////////////////////////////////////////////////////////
interfaces

only specify the min required behavior


reuse standard interfaces

kep the interfaces small

compose if needed

accept interfaces return concrete types



*/

