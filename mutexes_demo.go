package main

import (
	"sync"
	"math/rand"
	"sync/atomic"
	"time"
	"fmt"
)

func main(){
	var stat = make(map[int]int)
	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for r:=0;r<100;r++{
		go func() {
			total :=0
			for{
				key := rand.Intn(5)
				mutex.Lock()
				total += stat[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w:=0;w<10;w++{
		go func() {
			key := rand.Intn(5)
			val := rand.Intn(100)
			mutex.Lock()
			stat[key] += val
			mutex.Unlock()
			atomic.AddUint64(&writeOps,1)
			time.Sleep(time.Millisecond)
		}()
	}

	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:",readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:",writeOpsFinal)
	mutex.Lock()
	fmt.Println("state:",stat)
	mutex.Unlock()
}
