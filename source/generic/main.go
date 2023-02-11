package main

import (
	"errors"
	"fmt"
	"reflect"
)

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"a", "b", "c"})
	PrintSlice([]float64{1.2, -2.33, 4.55})

	fmt.Println("4 = 3 is", Same(4, 3))
	fmt.Println("aa = aa is", Same("aa", "aa"))
	fmt.Println("4.1 = 4.15 is", Same(4.1, 4.15))

	fmt.Println("4 + 3 =", Add(4, 3))
	fmt.Println("4.1 + 3.2 =", Add(4.1, 3.2))

	tempStr := TreeLast[string]{"aa", "bb"}
	fmt.Println(tempStr)
	tempStr.replaceLast("cc")
	fmt.Println(tempStr)
	tempInt := TreeLast[int]{12, -3}
	fmt.Println(tempInt)
	tempInt.replaceLast(0)
	fmt.Println(tempInt)

	//
	var myList list[int]
	fmt.Println(myList)
	myList.add(12)
	myList.add(9)
	myList.add(3)
	myList.add(9)
	// Print all elements
	for {
		fmt.Println("*", myList.start)
		if myList.start == nil {
			break
		}
		myList.start = myList.start.next
	}

	//
	Print(12)
	Print(-1.23)
	Print("Hi!")
	PrintGenerics(1)
	PrintGenerics("a")
	PrintGenerics(-2.33)
	PrintNumeric(1)
	PrintNumeric(-2.33)

	// 리플렉션
	PrintReflection([]int{1, 2, 3})
	PrintReflection([]string{"a", "b", "c"})
	PrintReflection([]float64{1.2, -2.33, 4.55})
}

// 제약 조건
func Same[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

// 제약 조건 만들기
type Numeric interface {
	int | int8 | int16 | int32 | int64 | float64
}

func Add[T Numeric](a, b T) T {
	return a + b
}

// 데이터 타입 정의
type TreeLast[T any] []T

func (t TreeLast[T]) replaceLast(element T) (TreeLast[T], error) {
	if len(t) == 0 {
		return t, errors.New("This is empty!")
	}

	t[len(t)-1] = element
	return t, nil
}

// node list
type node[T any] struct {
	Data T
	next *node[T]
}

type list[T any] struct {
	start *node[T]
}

func (l *list[T]) add(data T) {
	n := node[T]{
		Data: data,
		next: nil,
	}

	if l.start == nil {
		l.start = &n
		return
	}

	if l.start.next == nil {
		l.start.next = &n
		return
	}

	temp := l.start
	l.start = l.start.next
	l.add(data)
	l.start = temp
}

// 인터페이스 비교
func Print(s interface{}) {
	// type switch
	switch s.(type) {
	case int:
		fmt.Println(s.(int) + 1)
	case float64:
		fmt.Println(s.(float64) + 1)
	default:
		fmt.Println("Unknown data type!")
	}
}

func PrintGenerics[T any](s T) {
	fmt.Println(s)
}
func PrintNumeric[T Numeric](s T) {
	fmt.Println(s + 1)
}

// 리플렉션
func PrintReflection(s interface{}) {
	fmt.Println("** Reflection")
	val := reflect.ValueOf(s)

	if val.Kind() != reflect.Slice {
		return
	}

	for i := 0; i < val.Len(); i++ {
		fmt.Print(val.Index(i).Interface(), " ")
	}
	fmt.Println()
}
