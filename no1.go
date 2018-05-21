package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// 1
	// start1 := time.Now()
	// _, err1 := http.Get("https://fb.com")
	// if err1 != nil {
	// 	fmt.Println("[Error] Get Http", err1)
	// }
	// elapsed1 := time.Now() - start1
	// fmt.Println("elapse 1 =", elapsed1)

	// answer : invalid operation, because time.Now() not calculation with the same data types.

	// 2
	start2 := time.Now()
	_, err2 := http.Get("https://fb.com")
	if err2 != nil {
		fmt.Println("[Error] Get Http", err2)
	}
	elapsed2 := time.Since(start2).Seconds()
	fmt.Println("elapse 2 =", elapsed2)

	// 3
	start3 := time.Now()
	_, err3 := http.Get("https://fb.com")
	if err3 != nil {
		fmt.Println("[Error] Get Http", err3)
	}
	elapsed3 := time.Now().After(start3)
	fmt.Println("elapse 3 =", elapsed3)

	// answer : tidak bisa menghitung time consume, karena time.Now().After() return value boolean (true or false), sehingga tidak bisa

	// 4
	start4 := time.Now()
	_, err4 := http.Get("https://fb.com")
	if err4 != nil {
		fmt.Println("[Error] Get Http", err4)
	}
	elapsed4 := time.Now().Sub(start4)
	fmt.Println("elapse 4 =", elapsed4)

}
