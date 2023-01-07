package main

import (
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("Error occured while running hhtp server: %s", err.Error())
	}
}
