package main

import "fmt"

func main(){
	println("hello world")

	i:=1
	for i<10 {
		println(i)
		i++
	}
	for j:=1;j<5;j++{
		println("heloo world")
	}
	for {
		fmt.Println("loop")
		break
	}
	fmt.Print("vfast")
	boolean:=true
	for boolean {
		boolean:=false
	}
}
