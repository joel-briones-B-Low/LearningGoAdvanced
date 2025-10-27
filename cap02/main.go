package main

import "fmt"

func tarea(id int, canal chan<- string) {

	canal <- fmt.Sprintf("Tarea %d completada", id)

}

func main() {
	// solo puedo enviar string por este canal
	canal := make(chan string)
	for i := range 3 {
		go tarea(i, canal)
	}

	for range 3 {
		msg := <-canal
		fmt.Println(msg)
	}

}
