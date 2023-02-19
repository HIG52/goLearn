package main

import (
	"fmt"

	"github.com/HIG52/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)

	/*
		word := "hello"
		definition := "Greeting"

		err := dictionary.Add(word, definition)
		if err != nil {
			fmt.Println(err)
		}
		hello, _ := dictionary.Search(word)
		fmt.Println(hello)

		err2 := dictionary.Add("test", "testResult")
		if err2 != nil {
			fmt.Println(err2)
		}
		test, _ := dictionary.Search("test")
		fmt.Println(test)
	*/
	/*
		definition, err := dictionary.Search("first")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(definition)
	*/
}

/*
// main.go:6:2: no required module provides package github.com/HIG52/learngo/banking: go.mod file not found in current directory or any parent directory; see 'go help modules'
// 라는 에러가 떠서
// go mod init github.com/HIG52/learngo 이걸로 고쳐줌
// error 체크 method 생성 등등
func main() {
	account := accounts.NewAccount("Hongjun")
	account.Deposit(20)

	account.ChangeOwner("hong")
	fmt.Println(account.Balance())
	err := account.Withdraw(10)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(account)
}
*/
/*
// struct
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"김치찌개", "볶음김치"}
	hong := person{name: "hong", age: 27, favFood: favFood}

	//fmt.Println(hong)
	fmt.Println(hong.name)

}
*/
/*
//map
func main() {
	hong := map[string]string{"name": "jun", "age": "27"}

	for key, value := range hong {
		fmt.Println(key, value)
	}
}
*/
/*
// array & slice
func main() {
	names := []string{"hong", "jun", "yeong"} //slice (크기를 넣지 않음)
	names = append(names, "test1")

	fmt.Println(names)
}
*/
/*
//pointer
func main() { //메모리 주소 볼때 앞에 & 붙이기, *는 살펴보거나 훑어본다는 느낌으로 사용한다
	a := 2
	b := &a
	*b = 20
	//fmt.Println(*b)
	fmt.Println(a)
}
*/
//--------기본 문법 등등 if, for, switch naked return, defer 등등
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
/*
func lenAndUpper(name string) (length int, uppercase string) { //naked return
	defer fmt.Println("I'm done") //func가 끝난후 실행해되는 코드
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
*/
/*
func superAdd(numbers ...int) int { //여러개의 arguments를 받아오기위해 ...을 찍어줌
	//range 란 array에 loop를 적용할 수 있도록 해줌
	total := 0
	for _, numbers := range numbers { //numbers안에서 조건에 따라 반복 실행을 하도록 해줌
		total += numbers
	}
	return total
}
*/
/*
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}
*/
/*
func canIDrink(age int) (bool, int) {
	koreanAge := age + 2
	//switch koreanAge := age + 2; koreanAge {
	switch koreanAge {
	case 18:
		return true, koreanAge
	case 10:
		return false, koreanAge
	}
	return false, koreanAge
}
*/

//func main() {
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
/*
	totalLenght, up := lenAndUpper("jun")
	fmt.Println(totalLenght, up)
*/
/*
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
*/
/*
	fmt.Println(canIDrink(18))
*/
/*
	fmt.Println(canIDrink(18))
*/
//}
