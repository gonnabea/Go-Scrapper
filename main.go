package main // main은 컴파일을 위한 특별한 패키지 이름
import (
	"fmt"
)

// func multiply(a int, b int) int {
// 	return a * b
// }

// // length와 uppercase는 리턴되는 값
// func lenAndUpper(name string) (length int, uppercase string) {
// 	defer fmt.Println("함수 수행 완료") // defer는 함수가 값을 리턴한 후에 실행됨
// 	length = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }

// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }

// func superAdd(numbers ...int) int {
// 	// for문의 첫번쨰 인자는 index, 두번째는 배열의 요소
// 	for index, number := range numbers {
// 		fmt.Println("number")

// 		fmt.Println(index, number)

// 	}

// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Println(numbers[i])
// 	}

// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}

// 	return total
// }

// func canIDrink(age int) bool {

// 	// if 문 생성과 동시에 선언되는 변수를 만들 수 있음
// 	// 다른 언어에서는, 어딘가에 변수를 선언해두고 if문에 적용해야함
// 	// 이것은 가독성과 유지보수성을 떨어뜨림, go는 이 단점을 해결
// 	if koreanAge := age + 2; koreanAge < 18 {
// 		return false
// 	}
// 	return true

// }

// func canIDrink(age int) bool {
// 	switch koreanAge := age + 2; koreanAge {

// 	case 10:
// 		return false
// 	case 18:
// 		return true
// 	}
// 	return false
// }

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	// nico := person{"nico", 18, favFood}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)

	// nico := map[string]string{
	// 	"name": "nico",
	// 	"age":  "12",
	// }

	// for key, value := range nico {
	// 	fmt.Println(key, value)
	// }

	// names := []string{"nico", "lynn", "dal"}
	// names = append(names, "12312")

	// fmt.Println(names)

	// fmt.Println(canIDrink(18))
	// fmt.Println(canIDrink(16))

	// result := superAdd(1, 2, 3, 4, 5, 6)
	// fmt.Println(result)

	// totalLength, uppercase := lenAndUpper("jiwon")
	// fmt.Println(totalLength, uppercase)

	// repeatMe("nico", "lynn", "dal", "marl", "flynn")

	// // go가 자동으로 타입을 찾아줌: 축약형
	// // 축약형 ( := )은 해당 함수 스코프에만 유효
	// name := "nico"
	// test := true
	// // var name string = "nico"
	// name = "lynn"
	// fmt.Println(multiply(2, 2))
}
