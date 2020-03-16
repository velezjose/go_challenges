package main

import "fmt"

func main() {
	/**
	 * Goroutines & Channels:
	 *  1. Run 2 goroutines within the main goroutine.
	 *  2. Use two channels to share "Ping" and "Pong" messages between the goroutines so they can print:
	 *     Ping
	 *     Pong
	 *     Ping
	 *     ..., indefinitely every second.
	 *  Hint: Use a WaitGroup to avoid main goroutine exiting.
	 */

		fmt.Println("Hello world")
}
