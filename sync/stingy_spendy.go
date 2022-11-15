package main

import (
	"fmt"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var (
	b             = 100
	threadProfile = pprof.Lookup("threadcreate")
	lock          = sync.Mutex{}
)

func stingy() {
	for i := 0; i < 10; i++ {
		lock.Lock()
		b += 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Stingy finished")
}

func spendy() {
	for i := 0; i < 10; i++ {
		lock.Lock()
		b -= 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Spendy finished")
}

func main() {

	fmt.Println("running...")

	fmt.Println(runtime.NumCPU())
	fmt.Printf("threads in starting: %d\n", threadProfile.Count())

	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	print(b)

	fmt.Println(runtime.NumCPU())
	fmt.Printf("threads in starting: %d\n", threadProfile.Count())

}
