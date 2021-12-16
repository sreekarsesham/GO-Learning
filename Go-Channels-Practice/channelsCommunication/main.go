package main

import(

	"fmt"
	"time"
)

func main(){

	fmt.Println("channels communicate\n\n ")
	
	var c chan string = make(chan string)
	
	go send(c)
	go send2(c)
	receive(c)
	
	
}

func send(c chan string){
	start := time.Now()
	fmt.Println("started send at : ",start)
	fmt.Println("Iam sending")
	c<-"routine1"
}
func receive(c chan string){
	start := time.Now()
	fmt.Println("started receive at : ",start)
	fmt.Println("I received")
	fmt.Println("from send 1 : ",<-c)
	fmt.Println("from send 2 : ",<-c)
}

func send2(c chan string){

	start := time.Now()
	fmt.Println("started send at : ",start)
	fmt.Println("Iam sending")
	c<-"routine2"
	

}
