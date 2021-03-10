package main

import (
	"log"
)

func main() {
	log.Println("starts")

	var gmux = NewRouter()
	RegisterRoutes(gmux)
	Run(gmux)

	log.Printf("ready")
	go func() {
		WaitBreak()
		log.Println("shutting down by break begin")
	}()
	WaitExit()
	log.Println("shutting down complete.")
}
