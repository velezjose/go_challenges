//////////////////////////////////////////////////////////////////////
//
// Note: I obtained this exercise from https://github.com/loong/go-concurrency-exercises
//
// Producer and Consumer
//
// Background:
//  Given is a producer-consumer scenario, where a producer reads in
//  tweets from a mockstream and a consumer is processing the data.
//
// TODO
//  Your task is to change the code so that the producer as well as the
//  consumer can run concurrently. This will make the stream processing faster.
//
//  From about "Process took 3.608734916s" to about "Process took 1.980625118s"
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream) (tweets []*Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return tweets
		}

		tweets = append(tweets, tweet)
	}
}

func consumer(tweets []*Tweet) {
	for _, t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// Producer
	tweets := producer(stream)

	// Consumer
	consumer(tweets)

	fmt.Printf("Process took %s\n", time.Since(start))
}
