package something

import "fmt"

// 소문자로 시작할 시 외부로 export 불가
func sayBye() {
	fmt.Println("Bye")
}

// 대문자로 시작할 시 외부로 export 가능
func SayHello() {
	fmt.Println("Hello")
}
