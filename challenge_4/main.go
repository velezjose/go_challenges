package main

import (
		"fmt"
		"time"
)

var urls = []string {
		"https://apple.com", "https://applemusic.com",	"https://google.com", "https://youtube.com", "https://wikipedia.org",
		"https://facebook.com", "https://yahoo.com", "https://airbnb.com", "https://amazon.com", "https://live.com",
}

func main() {
	/*
	 * Http Package:
	 *  1. Use the `net/http` package to send GET requests to every url in the urls slice.
	 *  2. When the response is received, print the Status Code and the url.
	 *  3. If you didn't do it before, make every request run on a separate goroutine. Notice the improvement in time!
	 *  Bonus: Use a mutex (`sync` package) to lock the one resource that will be shared by all goroutines (??).
	 */
	start := time.Now()

	// your code here

	fmt.Println("total process took", time.Since(start))
}
