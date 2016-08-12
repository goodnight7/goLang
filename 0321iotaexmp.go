package main

import "fmt"
const (
	i=1<<iota  //i=1<<0
    j=3<<iota	//j=3<<1
    k	//k=3<<2
    l	//l=3<<3
)

func main() {
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
}