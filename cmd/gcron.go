package main

import "time"

func main() {
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			//TODO
		default:
			// TODO

		}
	}
}
