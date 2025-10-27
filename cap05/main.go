package main

import (
	"fmt"
	"time"
)

func tarea(id int, canal chan<- string) {
	time.Sleep(time.Duration(id) * 1000 * time.Millisecond)
	canal <- fmt.Sprintf("Tarea %d completada", id)
}

func main() {

	canal := make(chan string, 3)

	for i := range 6 {
		go tarea(i, canal)

	}

	timeout := time.After(4 * time.Second)

	for range 6 {
		select {
		case msg := <-canal:
			fmt.Println(msg)
		case <-timeout:
			fmt.Println("Tiempo de espera agotado")
			return
		}
	}

}
