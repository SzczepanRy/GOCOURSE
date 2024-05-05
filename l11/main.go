package main

import (
	"fmt"
	"os"
	"strings"
	"bytes"


	"golang.org/x/net/html"
)


	var raw = `
	<!DOCTYPE html>
	<html>
		<body>
		<h1>My First page </h1>
		<p>MY first paragraph</p>
		<p>HTML imafes are defined within the img tag:</p>
		<img src="xxx.jpg" width="104" height="142">
		<img src="xxx.jpg" width="104" height="142">
		</body>
	</html>`
	
func main(){

	//homework 
	doc ,err := html.Parse(bytes.NewReader([]byte(raw)))
	if err!= nil{
		fmt.Println(os.Stderr,"pase false")
	os.Exit(-1)
	}
	words , pics := CountWordsAndImages(doc)

	fmt.Println("words ", words, "pics ", pics)

}



func visit(n *html.Node , words ,pics *int ){

	//is is ana element node , what tag does it have
	if  n.Type == html.TextNode{
		*words += len(strings.Fields(n.Data))
	}else if n.Type == html.ElementNode && n.Data == "img"{

		*pics++
	}




	for c:= n.FirstChild; c!=nil;c =  c.NextSibling{
	 	visit(c, words , pics)

	}

	

}
func CountWordsAndImages(doc *html.Node)(int ,int ){

	var words , pics int
	

	visit(doc , &words , &pics)

	return words , pics


}
