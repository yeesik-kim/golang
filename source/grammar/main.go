// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var num1 int
var num2, num3 int
var num4, num5, str1 = 4, 5, "example"

// errorvar := str1
var (
	i int
	b bool
	s string
)

func main() {
	// 변수
	fmt.Println("①", num1)
	fmt.Println("②", num2, num3)
	fmt.Println("③", num4, num5, str1)
	num6 := 6
	fmt.Println("④", num6)
	fmt.Println("⑤", i, b, s)

	// 문자열
	byteString := []byte("A String")
	normalString := "normal"
	fmt.Printf("%s", byteString)
	fmt.Println()
	fmt.Println(string(byteString))
	fmt.Println(normalString)
	// 룬 형태로 출력
	for _, v := range byteString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()
	// 문자 형태로 출력
	for _, v := range byteString {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	// if
	v1 := 0
	if v1 == 0 {
		fmt.Println("V1 is zero")
	}
	if v2 := 1; v2 > 0 {
		fmt.Println("V2 is none zero")
	}

	// for
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// 세미콜론 없이, C의 while과 비슷하게 사용 가능.
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// 배열
	// ① 배열선언과 원소 초기화를 따로
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("a[0], a[1]:", a[0], a[1])
	fmt.Println("a:", a)
	// ② 배열선언과 초기화를 동시에
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes:", primes)

	// 맵
	// make 사용
	mymap := make(map[string]int)
	mymap["key1"] = 123
	mymap["key2"] = 456
	fmt.Println("① mymap[\"key1\"]: ", mymap["key1"])
	fmt.Println("① mymap[\"key2\"]: ", mymap["key2"])
	// 리터럴
	m := map[string]int{
		"key1": -1,
		"key2": 0,
	}
	fmt.Println("① m[\"key1\"]: ", m["key1"])
	fmt.Println("① m[\"key2\"]: ", m["key2"])

	// 포인터
	i, j := 42, 2701
	p := &i         // i를 가리키는 포인터
	fmt.Println(*p) // 포인터를 통해 i 값을 읽습니다.
	*p = 21         // 포인터를 통해 i값을 설정합니다.
	fmt.Println(i)
	p = &j       // j를 가리킵니다.
	*p = *p / 37 // 포인터를 통해 j를 나눕니다.
	fmt.Println(j)

	// 구조체
	structures()

	// 함수
	fmt.Println("add1(x int, y int)의 결과: ", add1(42, 13))
	fmt.Println("add2(x, y int)의 결과: ", add2(42, 13))

	// ①로 한 번에 여러개의 결과를 return받는 부분
	var quotient, remainder int
	quotient, remainder = divide1(10, 3)
	fmt.Println("①의 결과:", quotient, remainder)
	// ②로 한 번에 여러개의 결과를 return받는 부분
	quotient, remainder = divide2(10, 3)
	fmt.Println("②의 결과:", quotient, remainder)

}

func structures() {
	fmt.Println("v1.X값:", v1.X)
	v1.X = 4
	fmt.Println("v1.X = 4로 바꾼 v1.X값:", v1.X)
	// new 키워드 : 구조체의 포인터를 반환
	ps := new(Vertex)
	fmt.Println("ps:", ps)
}

// 상수
// 상수도 선언과 동시에 초기화하면 타입을 지정하지 않아도 됩니다.
const Pi1 float32 = 3.14
const Pi2 = 3.14

// 괄호로 묶으면 상수 키워드를 한 번만 명시합니다.
const (
	Big_const   = 1 << 100
	Small_const = Big_const >> 99
)

// 구조체
type Vertex struct {
	X int
	Y int
}

// 구조체 인스턴스 선언 방법
var (
	//① 일반적인 선언방식
	v1 = Vertex{1, 2}
	//② X만 값을 지정해주고, Y는 int에 zero value로 설정
	v2 = Vertex{X: 1}
	//③ X, Y모두 int에 zero value로 설정
	v3 = Vertex{}
)

// 함수
// ① 매개변수 타입, 리턴 타입은 이름 뒤에 지정해줍니다
func add1(x int, y int) int {
	return x + y
}

// ② 매개변수 x, y가 같은 타입일 때에는 타입을 한 번만 명시해 줄 수 있습니다.
func add2(x, y int) int {
	return x + y
}

// ①
func divide1(dividend, divisor int) (int, int) {
	var quotient = (int)(dividend / divisor)
	var remainder = dividend % divisor
	return quotient, remainder
}

// ②
func divide2(dividend, divisor int) (quotient, remainder int) {
	quotient = (int)(dividend / divisor)
	remainder = dividend % divisor
	return //return이라고만 적으면 미리 return값으로 정해 놓은 quotient와 remainder를 return합니다.
}
