package main

import (
"fmt"
"time"
)

var urls = []string {
	"https://apple.com", "https://applemusic.com",	"https://google.com", "https://youtube.com", "https://wikipedia.org",
	"https://facebook.com", "https://yahoo.com", "https://airbnb.com", "https://amazon.com", "https://live.com",
	"https://reddit.com", "https://netflix.com", "https://blogspot.com", "https://github.com", "https://hulu.com",
	"https://espn.com", "https://uber.com", "https://lyft.com", "https://shopify.com", "https://ebay.com",
}

func main() {
	/*
	 * Http Package:
	 *  1. Use the `net/http` package to send GET requests to every url in the urls slice.
	 *  2. When the response is received, print the Status Code and the url.
	 *  3. Take the total time it takes to send and receive every request, from start to finish.
	 *  4. If you didn't do it before, make every request run on a separate goroutine. Notice the improvement in time!
	 *  Bonus: Use a mutex (`sync` package) to lock the one resource that will be shared by all goroutines (??).
	 */
	start := time.Now()

	fmt.Println("total process took", time.Since(start))
}
