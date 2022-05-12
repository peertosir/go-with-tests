package main

import "fmt"

func main() {
	someArr := [2]int{1, 2}
	slSomeArr := someArr[:1]
	fmt.Println(len(slSomeArr))
	fmt.Println(cap(slSomeArr))
	slSomeArr = append(slSomeArr, 7)
	slSomeArr = append(slSomeArr, 12)
	slSomeArr[0] = 12
	fmt.Println(slSomeArr)
	fmt.Println(someArr)
}
