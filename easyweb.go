package main

import "log"

func on_start() {
	log.Printf("Seeking connection")
	if wait_for_connection() != nil {
		restart()
	}
	log.Printf("Found connection")
	open_browser()
	take_hdmi()
	log.Printf("Ran startup commands.")
}

func main() {
	log.Printf("Starting up..")
	on_start()
	serve()
}
