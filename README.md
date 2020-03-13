# Go Workshop

This repo contains the content discussed in our Golang workshops. There are some basic setup commands to set up go in 
your local dev environment, and some challenges to go through during the time allocated.

Let's set up our workspace before we go on to start with our projects.

## Setup:

1. Open a terminal session.
2. `brew install go` (If you have it already, then see if you can update it: `brew update go`)
3. Clone (or fork and then clone) this repo `git clone git@github.pie.apple.com:jose-v/go_workshops.git`
4. Fingers crossed it works!
5. Download Goland IDE, and open this repo in it.


## 1. FizzBuzz: 

This is a warmup. Write a short program that prints each number from 1 to 100 on a new line. 
- For each multiple of 3, print `Fizz` instead of the number. 
- For each multiple of 5, print `Buzz` instead of the number. 
- For numbers which are multiples of both 3 and 5, print "FizzBuzz" instead of the number.


## 2. Fibonacci Closure

- Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...) every time it's invoked.
- For more info: https://tour.golang.org/moretypes/26


## 3. Goroutines & Channels

- Run 2 go routines within the main method.
- Use a channel to synchronize them and print `Ping`, `Pong`, `Ping`, `Pong`, .., indefinitely every second.


## 4. Http Package

1. Use the `net/http` package to send GET requests to every url provided.
2. When the response is received, print the Status Code and the url.
3. Take the total time it takes to send and receive every request, from start to finish.
4. If you didn't do it before, make every request run on a separate goroutine. Notice the improvement in time!
Bonus: Use a mutex (`sync` package) to lock the one resource that will be shared by all goroutines (??).


## 5. Producer and Consumer

- Given is a producer-consumer scenario, where a producer reads in tweets from a mockstream and a consumer is processing the data.
- Your task is to change the code so that the producer as well as the consumer can run concurrently.


## 6. Video Processing Service

- Your video processing service has a freemium model. 
- Everyone has 10 seconds of free processing time on your service. After that, the service will kill your process, unless you are a paid premium user.
- Apply the 10s max per user (accumulated) rule.
