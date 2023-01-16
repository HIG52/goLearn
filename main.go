package main

import (
	"fmt"
	"strings"
)

/*
func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(world ...string) { //여러개의 argument값을 받아오려면 타입 앞에 ... 찍어주기
	fmt.Println(world)
}
*/

func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	//축약형은 오로지 func안에서만 가능하고 변수에만 적용 가능하다
	/*
			name := "jun" //== v var name string = "jun"
			name = "yeong"
			fmt.Println(name)


		 fmt.Println(multiply(2, 2))

		totalLenght, upperName := lenAndUpper("jun")
		totalLenght, _ := lenAndUpper("jun") //_를 사용하면 무시하고 그냥 하나만 리턴받아서 쓸스도 있음
		fmt.Println(totalLenght)

		repeatMe("hong", "jun", "yeong", "good", "guy")
	*/

	totalLenght, up := lenAndUpper("jun")
	fmt.Println(totalLenght, up)

}
