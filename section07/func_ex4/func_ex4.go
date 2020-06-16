//함수 심화(4)
package main

import "fmt"

func main() {
	//함수 고급
	//익명 함수(Anonymous Functions)

	//예제1
	func() {
		fmt.Println("ex1 : Anonymous func!")
	}()

	fmt.Println()

	//예제2
	msg := "Hello Golang!"

	func(m string) {
		fmt.Println("ex2 : ", m)
	}(msg)

	fmt.Println()

	//예제3
	func(x, y int) {
		fmt.Println("ex3 : ", x+y)
	}(10, 20)

	fmt.Println()

	//예제4
	r := func(x, y int) int {
		return x * y
	}(10, 10)
	fmt.Println("ex4 : ", r)
}
