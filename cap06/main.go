package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	buf := NewBuffer()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		buf.mu.Lock()
		for len(buf.data) == 0 {
			buf.cond.Wait()
		}
		// Process the data
		datas := buf.data[0]
		buf.data = buf.data[1:]
		buf.mu.Unlock()

		defer fmt.Println("Procesado:", datas)

	}()

	time.Sleep(500 * time.Millisecond)
	buf.mu.Lock()
	buf.data = append(buf.data, 42)
	buf.cond.Signal() // notify waiting goroutine
	buf.mu.Unlock()
	wg.Wait()
}
