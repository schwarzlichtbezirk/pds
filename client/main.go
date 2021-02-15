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
	WaitBreak()
	log.Println("shutting down begin")
	Shutdown()
	log.Println("shutting down complete.")
}
