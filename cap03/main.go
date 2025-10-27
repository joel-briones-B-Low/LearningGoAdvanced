package main

import (
	"fmt"
	"sync"
)

var completadas int

var mu sync.Mutex

func tarea(id int, canal chan<- string) {
	mu.Lock()
	completadas++
	mu.Unlock()
	canal <- fmt.Sprintf("Tarea completada: %d", id)

}

func main() {
	canal := make(chan string)
	for i := range 3 {
		go tarea(i, canal)
	}

	for range 3 {
		fmt.Println(<-canal)
	}
	mu.Lock()
	fmt.Printf("Total de tareas %d", completadas)
	mu.Unlock()

}
