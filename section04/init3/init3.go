//GO 초기화 함수(3)
package main

import (
	"basic/section04/lib"
	"fmt"
)

var num int32

//변수 초기화
func init() {
	num = 30
}

func main() {
	fmt.Print("10 보다 큰 수? :  ", lib.CheckNum(num))
}
