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
		println("hello world")
	}
	for {
		fmt.Println("loop")
		break
	}
	fmt.Print("vfast")
	boolean:=true
	for boolean {
		fmt.Println("aaaaaaaa")
		break
	}
	sum :=0
	for i:=1;i<=1000;i++{
		sum +=i

	}
	println(sum)
	for i,j := 20,20 ;i<30 && j < 30;i++{
		println(j," ",i)
		j += 2
	}
}
