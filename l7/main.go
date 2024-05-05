package main 

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main (){
	//fileio
	//gom main.go *.txt
	/*for _, filename := range os.Args[1:]{
		file,errr:=os.Open(filename)
		if errr !=nil{
			fmt.Fprintln(os.Stderr,errr)
			continue
		}
		if _,err:= io.Copy(os.Stdout,file);err!=nil{
			fmt.Fprintln(os.Stderr,err)

		}
		file.Close()
	}*/



	opt := os.Args[1:2][0]

	for _, filename := range os.Args[2:]{
		var lc,wc,cc int
		file,err:=os.Open(filename)
		if err!=nil{
			fmt.Fprintln(os.Stderr,err)
			continue
		}


		scan := bufio.NewScanner(file)

		for scan.Scan(){
			s:=scan.Text()
			cc += len(s)
			lc++
			wc += len(strings.Fields(s))
		}
		
	 	 switch opt{
		 case "wc":
			fmt.Println("word count" , wc)
		 case "cc":
			fmt.Println("char count" , cc)

		 case "lc":
			fmt.Println("line count" , lc)
		 default:
			fmt.Println(" line count ",lc,"word count ", wc, "char count" ,cc)
		 }

		file.Close()
	}
	
}
