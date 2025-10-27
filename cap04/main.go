package main

import (
	"fmt"
	"sync"
)

var (
	completadas int
	mu          sync.Mutex
)

func tarea(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	completadas++
	mu.Unlock()

	fmt.Println("Tarea", i, "completada")
}

func main() {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go tarea(i, &wg)
	}
	wg.Wait()
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("Total de tareas completadas:", completadas)

}
