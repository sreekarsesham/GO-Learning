package main

import(

	"fmt"
)

func main(){

	fmt.Println("GO Channels")
	
	// creating a channel
	
	var c chan string = make(chan string)
	
	//c<-1998; // uncomment this to see failing code 
	
	//fmt.Println(<-c) //trying to print before creating routine but it doesnt becoz we need atleast 2 go routines to communicate
	
	// creating a go routine 
	
	go func(){
	
		c<-"sreekar"
	}()
	
	fmt.Println(<-c)
	
	
}
