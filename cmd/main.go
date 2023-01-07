package main

import (
	"log"

	rsapi "github.com/wanna-beat-by-bit/rest-golang"
)

func main() {
	srv := new(rsapi.Server)

	if err := srv.Run("8000"); err != nil {
		log.Fatalf("Error occured whilte running http server: %s", err)
	}
}
