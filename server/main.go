package main

import (
	"log"
)

func main() {
	log.Println("starts")
	Run()
	log.Printf("ready")
	go func() {
		WaitBreak()
		log.Println("shutting down by break begin")
	}()
	WaitExit()
	log.Println("shutting down complete.")
}
