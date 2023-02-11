package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}
func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go PrintHangul()
	go PrintNumbers()
	time.Sleep(3 * time.Second)

	// 종료 기다리기
	wg.Add(10) // ❷ 총 작업 개수 설정
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000000)
	}
	wg.Wait() // ❹ 모든 작업이 완료되길 기다림
	fmt.Println("모든 계산이 완료되었습니다.")

	// 채널과 컨텍스트
	routine1()
	routine2()
	routine3()
	routine4()
}

var wg sync.WaitGroup // ❶ waitGroup 객체
func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)
	wg.Done() // ❸ 작업이 완료됨을 표시
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch               // ➎ 데이터를 빼온다
	time.Sleep(time.Second) // 1초 대기
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}

func routine1() {
	var wg sync.WaitGroup
	ch := make(chan int) // ❶ 채널 생성
	wg.Add(1)
	go square(&wg, ch) // ❷ Go 루틴 생성
	ch <- 9            // ❸ 채널에 데이터를 넣는다.
	wg.Wait()          // ❹ 작업이 완료되길 기다린다.

}

// close
func square2(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { // ❷ 채널이 닫히면 종료
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func routine2() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go square2(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) // ←- ❶ 채널을 닫는다.
	wg.Wait()
}

// select
func square3(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select { // ❷ ch와 quit 양쪽을 모두 기다린다.
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func routine3() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool) // ❶ 종료 채널
	wg.Add(1)
	go square3(&wg, ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	quit <- true
	wg.Wait()
}

// tick
func square4(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)            // ❶ 1초 간격 시그널
	terminate := time.After(10 * time.Second) // ❷ 10초 이후 시그널
	for {
		select { // ❸ tick, terminate, ch 순서로 처리
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated!")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func routine4() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go square(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}
