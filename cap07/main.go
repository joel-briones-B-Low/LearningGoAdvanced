package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var estado sync.Map

	var completadas int64
	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			estado.Store(fmt.Sprintf("Tarea completada: %d", id), "-- Completada --")
			atomic.AddInt64(&completadas, 1)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Total de tareas completadas: %d\n", completadas)
	estado.Range(func(k, v any) bool {
		fmt.Printf("Estado de: %s: -- %s -- sss\n ", k, v)
		return true
	})

	fmt.Printf("completadas: %d\n", completadas)
}
