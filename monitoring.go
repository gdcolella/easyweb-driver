package main

import "log"
import "net"
import "time"

var TRIES = 10

func wait_for_connection() error {
	timeout := time.Duration(30) * time.Second
	i := 0
	var e error = nil
	for i < TRIES {
		conn, err := net.DialTimeout("tcp", SERVER, timeout)
		e = err
		if e == nil {
			conn.Close()
			return nil
		}
		time.Sleep(timeout)
		log.Printf("Failed to connect %d %s", i, e)
		i++
	}
	if e != nil {
		log.Printf("Failed to connect over all retries.")
	}
	return e
}
