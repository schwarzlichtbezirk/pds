package main

import (
	"log"
)

func main() {
	log.Println("starts")
	Run()
	log.Printf("ready")
	WaitBreak()
	log.Println("shutting down begin")
	Shutdown()
	log.Println("shutting down complete.")
}
