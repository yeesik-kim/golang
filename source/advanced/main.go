package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var slice1 []int            // 크기를 지정하지 않으면 slice. nil
	// slice2 := []int{}           // 길이와 용량은 0으로 지정됨. (nil 아님)
	// slice3 := []int{1, 2, 3}    // 선언과 동시에 초기화 하면 길이와 용량은 요소의 개수로 지정됨
	// slice4 := make([]int, 2, 5) // 길이가 2이고 용량은 5인 슬라이스 선언

	words := make([]string, 3, 5)
	words[0] = "A"
	words[1] = "B"
	words[2] = "C"
	fmt.Printf("Length=%d Capacity=%d\n", len(words), cap(words))

	fruits := make([]string, 3, 3)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fmt.Printf("1: %p\t%p\n", &fruits, &fruits[0])
	printAddress(fruits)

	slice1()
	reslice1()
	slicesize()
	slicedel()
	slicedel2()

	method()
	method2()
	method3()

	interface1()
	interface2()
	interface3()
	interface4()

	fmt.Println(sum(1, 2, 3, 4, 5)) //5개의 인수를 사용합니다.
	fmt.Println(sum(10, 20))        //2개의 인수를 사용합니다.
	fmt.Println(sum())              //0개의 인수를 사용합니다.

	fmt.Println(2, "hello", 3.14)

	// defer
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer fmt.Println("반드시 호출됩니다.") // 1
	defer f.Close()                 // 2
	defer fmt.Println("파일을 닫았습니다")  // 3
	fmt.Println("파일에 Hello World를 씁니다.")
	fmt.Fprintln(f, "Hello World")

	//int 타입 인수 2개를 받아서 int 타입 반환을 하는 함수 타입 변수
	var operator func(int, int) int
	operator = getOperator("*")
	var result = operator(3, 4) //함수 타입 변수를 사용해서 함수 호출
	fmt.Println(result)

	fn := getOperator2("*")
	result2 := fn(3, 4) //함수 타입 변수를 사용해서 함수 호출
	fmt.Println(result2)

	CaptureLoop()
	CaptureLoop2()

	funct()
	error1()
	error2()
	error3()

	readEq("123 3")
	readEq("123 abc")

	divide(9, 3)
	// divide(9, 0) // ❷ Panic 발생

	fpanic()
	fmt.Println("프로그램이 계속 실행됨") // 프로그램 실행 지속됨
}

var printAddress = func(slice []string) {
	fmt.Printf("2: %p\t%p\n", &slice, &slice[0])
}

// 복사
func slice1() {
	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"
	punnet := make([]string, 3)
	copy(punnet, fruits[1:4])
	punnet[2] = "Strawberry"
	fmt.Println(fruits)
	fmt.Println(punnet)
}

// 리슽라이스
func reslice1() {
	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	punnet := fruits[1:4]
	punnet[2] = "Strawberry"

	fmt.Println(fruits)
	fmt.Println(punnet)
}

// 크기 증가
func slicesize() {
	numSlice := make([]int, 7)
	for i := 0; i < 7; i++ {
		numSlice[i] = i * 100
	}
	// numSlice의 1번째 요소를 참조
	willBeOutdate := &numSlice[1]
	fmt.Printf("Before append:\t%p\t%p\n", &numSlice, &numSlice[0])

	// append 함수 호출 시 슬라이스의 값이 복사되어 전달되므로 반환값을 사용해야 함
	numSlice = append(numSlice, 800)
	numSlice[1]++
	fmt.Printf("After append:\t%p\t%p\n", &numSlice, &numSlice[0])
	fmt.Println("willBeOutdate:", *willBeOutdate, "numSlice[1]:", numSlice[1])
}

// 요소 삭제
func slicedel() {
	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"
	fmt.Printf("Length[%d] Capacity[%d] %v\n", len(fruits), cap(fruits), fruits)
	// Banana를 삭제
	fruits = append(fruits[:2], fruits[3:]...)
	fmt.Printf("Length[%d] Capacity[%d] %v\n", len(fruits), cap(fruits), fruits)
}

// 요소 삭제 크기도
func slicedel2() {
	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"
	fmt.Printf("Length[%d] Capacity[%d] %v\n", len(fruits), cap(fruits), fruits)
	// Banana를 삭제
	fruits = deleteAt(2, fruits)
	fmt.Printf("Length[%d] Capacity[%d] %v\n", len(fruits), cap(fruits), fruits)

}

func deleteAt(index int, slice []string) []string {
	slice = append(slice[:index], slice[index+1:]...)
	newSlice := make([]string, len(slice))
	copy(newSlice, slice)
	return newSlice
}

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) { //일반 함수 표현
	a.balance -= amount
}
func (a *account) withdrawMethod(amount int) { //메서드 표현
	a.balance -= amount
}
func method() {
	a := &account{100}   //balance가 100인 account 포인터 변수 생성
	withdrawFunc(a, 30)  //함수 형태 호출
	a.withdrawMethod(30) //메서드 형태 호출
	fmt.Printf("%d \n", a.balance)
}

// method 2
// 사용자 정의 별칭 타입
type myInt int

// myInt 별칭 타입을 리시버로 갖는 메서드
func (a myInt) add(b int) int {
	return int(a) + b
}
func method2() {
	var a myInt = 10       //myInt 타입 변수
	fmt.Println(a.add(30)) //myInt 타입의 add() 메서드 호출
	var b int = 20
	fmt.Println(myInt(b).add(50)) //int 타입을 타입변환
}

// method 3
type account2 struct {
	balance   int
	firstName string
	lastName  string
}

// 포인터 타입 메서드
func (a1 *account2) withdrawPointer(amount int) {
	a1.balance -= amount
}

// 값 타입 메서드
func (a2 account2) withdrawValue(amount int) {
	a2.balance -= amount
}

// 변경된 값을 반환하는 값 타입 메서드
func (a3 account2) withdrawReturnValue(amount int) account2 {
	a3.balance -= amount
	return a3
}

func method3() {
	var mainA *account2 = &account2{100, "Joe", "Park"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance) //70 이 출력
	mainA.withdrawValue(20)    //포인터 변수 값타입 메서드 호출
	fmt.Println(mainA.balance) //여전히 70이 출력
	var mainB account2 = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance) //50이 출력
	mainB.withdrawPointer(30)  //값 변수 포인터타입 메서드 호출
	fmt.Println(mainB.balance) //20이 출력
}

// interface

type Stringer interface { // ❶ Stringer 인터페이스 선언 String() string
	String() string
}
type Student struct {
	Name string
	Age  int
}

func (s Student) String() string { // ❷ Student의 String() 메서드
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name) // ❸ 문자열 만들기
}
func interface1() {
	student := Student{"철수", 12}          // Student 타입
	var stringer Stringer                 // Stringer 타입
	stringer = student                    // 4 stringer값으로 student 대입
	fmt.Printf("%s\n", stringer.String()) // 5 stringer의 String() 메서드 호출
}

// Fedex에서 제공한 패키지 내 전송을 담당하는 구조체입니다.
type FedexSender struct{}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends %v parcel\n", parcel)
}

//	func SendBook(name string, sender FedexSender) {
//		sender.Send(name)
//	}
func interface2() {
	// Fedex 전송 객체를 만듭니다.
	// sender := FedexSender{}
	// SendBook("어린 왕자", sender)
	// SendBook("그리스인 조르바", sender)

	// 우체국 전송 객체를 만듭니다.
	// sender2 := PostSender{} // ❶ *koreaPost.PostSender 타입
	// SendBook("어린 왕자", sender2)         // ❷ 타입이 맞지 않습니다.
	// SendBook("그리스인 조르바", sender2)

	// ❸ 우체국 전송 객체, Fedex 전송 객체 모두 SendBook 인수로 사용할 수 있습니다.
	// 우체국 전송 객체를 만듭니다.
	koreaPostSender := &PostSender{}
	SendBook("어린 왕자", koreaPostSender)
	SendBook("그리스인 조르바", koreaPostSender)
	// Fedex 전송 객체를 만듭니다.
	fedexSender := &FedexSender{}
	SendBook("어린 왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)
}

// 우체국에서 제공한 패키지 내 전송을 담당하는 구조체입니다.
type PostSender struct{}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("우체국에서 택배 %v를 보냅니다.\n", parcel)
}

// ❶ Sender 인터페이스를 만들었습니다.
type Sender interface {
	Send(parcel string)
}

// ❷ Sender 인터페이스를 입력으로 받습니다.
func SendBook(name string, sender Sender) {
	sender.Send(name)
}

// 빈 인터페이스
func PrintVal(v interface{}) { // ❶ 빈 인터페이스를 인수로 받는 함수
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default: // 그외 타입인 경우 타입과 값을 출력합니다.
		fmt.Printf("Not supported type: %T:%v\n", t, t)
	}
}

type Student2 struct {
	Age int
}

func interface3() {
	PrintVal(10)           // int
	PrintVal(3.14)         // float64
	PrintVal("Hello")      // string
	PrintVal(Student2{15}) // Student
}

func (s *Student2) String() string { // ❸ Student 타입의 String() 메서드
	return fmt.Sprintf("Student Age:%d", s.Age)
}
func PrintAge(stringer Stringer) { //
	s := stringer.(*Student2)      // *Student 타입으로 타입 변환
	fmt.Printf("Age: %d\n", s.Age) // s.Age 출력
}

func interface4() {
	s := &Student2{15} // *Student 타입 변수 s 선언 및 초기화
	PrintAge(s)        // 변수 s 를 인터페이스 인수로 PrintAge() 함수 호출
}

// 함수 고급
func sum(nums ...int) int { //가변 인수를 받는 함수
	sum := 0
	if nums != nil {
		fmt.Println("용량이", cap(nums), "길이가", len(nums), " Slice입니다.")
	}
	fmt.Printf("nums 타입: %T\n", nums) //nums 타입 출력
	for _, v := range nums {
		sum += v
	}
	return sum
}

// 함수 타입 변수
func add(a, b int) int {
	return a + b
}
func mul(a, b int) int {
	return a * b
}
func getOperator(op string) func(int, int) int { //op에 따른 함수 타입 반환
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else { //+,*가 아니면 nil반환
		return nil
	}
}

// 함수 리터럴
type opFunc func(a, b int) int

func getOperator2(op string) opFunc {
	if op == "+" {
		return func(a, b int) int { //함수 리터럴을 사용해서 더하기 함수를 정의하고 반환
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int { //함수 리터럴을 사용해서 곱하기 함수를 정의하고 반환
			return a * b
		}
	} else {
		return nil
	}
}

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("CaptureLoop1")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("CaptureLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

// 파일 핸들 내부 상태
type Writer func(string)

func writeHello(writer Writer) { //writer 함수타입 변수 호출
	writer("Hello World")
}
func funct() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()
	writeHello(func(msg string) {
		fmt.Fprintln(f, msg) //함수 리터럴 외부 변수 f 사용
	})
}

// 에러 핸들링
func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename) // ❶ 파일 열기
	if err != nil {
		return "", err // ❷ 에러 나면 에러 반환
	}
	defer file.Close()             // ❸ 함수 종료 직전 파일 닫기
	rd := bufio.NewReader(file)    // ❹ 파일 내용 읽기
	line, _ := rd.ReadString('\n') // _ <- 에러 무시
	return line, nil
}
func WriteFile(filename string, line string) error {
	file, err := os.Create(filename) // 파일 생성
	if err != nil {                  // 에러 나면 에러 반환
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line) // 파일에 문자열 쓰기, _ <- 쓴길이 무시
	return err
}

const filename string = "data.txt"

func error1() {
	line, err := ReadFile(filename) // 파일 읽기 시도
	if err != nil {
		err = WriteFile(filename, "This is WriteFile") // 파일 생성
		if err != nil {                                // ❿ 에러를처리
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}
		line, err = ReadFile(filename) // ⓫ 다시 읽기 시도
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다.", err)
			return
		}
	}
	fmt.Println("파일내용:", line) // ⓬ 파일 내용 출력
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf(
			"제곱근은 양수여야 합니다. f:%g", f) // ❶ f가 음수이면 에러 반환
	}
	return math.Sqrt(f), nil
}
func error2() {
	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("Error: %v\n", err) // ❷ 에러 출력
		return
	}
	fmt.Printf("Sqrt(-2) = %v\n", sqrt)
}

// 에러 핸들링
type PasswordError struct { // ❶ 에러 구조체 선언
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string { // ❷ Error() 메서드
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8} // ❸ error 반환
	}
	return nil
}
func error3() {
	err := RegisterAccount("myID", "myPw") // ID, PW 입력
	if err != nil {                        // 에러확인
		if errInfo, ok := err.(PasswordError); ok { // 인터페이스 변환
			fmt.Printf("%v Len:%d RequireLen:%d\n",
				errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입되었습니다.")
	}
}

// 에러 랩핑
func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) // ❶ 스캐너 생성
	scanner.Split(bufio.ScanWords)                      // ❷ 한 단어씩 끊어읽기
	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err) // 에러 감싸기
	}
	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}
	return a * b, nil
}

// 다음 단어를 읽어서 숫자로 변환하여 반환합니다.
// 변환된 숫자, 읽은 글자수, 에러를 반환합니다.
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() { // ❸ 단어 읽기
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word) // ❹ 문자열을 숫자로 변환
	if err != nil {
		return 0, 0,
			fmt.Errorf("Failed to convert word to int, word:%s err:%w", word, err) // 에러 감싸기
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError // Atoi() <- 문자가 섞인 경우 NumError를 return
		if errors.As(err, &numError) { // 감싸진 에러가 NumError인지 확인
			fmt.Println("NumberError:", numError)
		}
	}
}

// panic
func divide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다") // ❶ Panic 발생
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

// panic 전파 복구
func fpanic() {
	fmt.Println("f() 함수 시작")
	defer func() { // ❹ 패닉 복구
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}
	}()
	g() //❶g()->h()순서로 호출
	fmt.Println("f() 함수 끝")
}
func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
	// ❷ h() 함수 호출 - 패닉
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다.") // ❸ 패닉 발생!!
	}
	return a / b
}
