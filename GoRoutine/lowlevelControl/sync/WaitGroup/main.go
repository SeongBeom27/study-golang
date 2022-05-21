package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	i int64
	mu sync.Mutex	// 공유 데이터 i를 보호하기 위한 뮤텍스
}

func (c *counter) increment(){
	c.mu.Lock()
	c.i += 1
	c.mu.Unlock()
}

func (c *counter) display() {
	fmt.Println(c.i)
}


func main() {
	// 모든 cpu를 사용하게 함
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}
	wg := sync.WaitGroup{}
	
	// c.increment()를 실행하는 고루틴 1000개 실행
	for i := 0; i < 1000; i++ {
		// WaitGroup의 고루틴 개수 1 증가
		wg.Add(1)
		go func() {
			defer wg.Done()		// 고루틴 종료 시 Done 처리
			c.increment()		// 카운터 값을 1 증가시킴
		}()
	}

	wg.Wait()

	c.display()	// c의 값 출력
}