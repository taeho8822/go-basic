//패키지 접근제어(2)
package main

import (
	_ "basic/section04/lib"        //빈 식별자 사용
	testlib "basic/section04/lib2" //별칭 사용
	"fmt"
)

func main() {
	//패키지 접근제어
	//별칭 사용
	//빈 식별자 사용

	fmt.Print("100 보다 큰 수? :  ", testlib.CheckNum1(101))
}
