//////////////////////////////////////////////////////////////////////
//
// Note: I obtained this exercise from https://github.com/loong/go-concurrency-exercises
//
// Video Processing Service
//
// Background:
//  Your video processing service has a freemium model. Everyone has 10
//  seconds of free processing time on your service. After that, the
//  service will kill your process, unless you are a paid premium user.
//
// TODO:
//  Apply the 10s max per user (accumulated) rule.
//

package main

import (
	"time";
	"log"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// Seconds of free processing time.
const FREE_TIER_TIME_LIMIT = 10;

/* Problems
How to kill process in the middle of running? Even if we return
false early, we're still doing extra work in the background.

Concurrent requests per user messes up the timer method. We could
lock it so each user gets one processer at a time. With a userId->lock map.
*/

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	if u.IsPremium {
		return HandlePremiumRequest(process, u)
	} else {
		return HandleRegularRequest(process, u)
	}
}

func HandlePremiumRequest(process func(), u *User) bool {
	process()
	return true
}

func HandleRegularRequest(process func(), u *User) bool {

	if u.TimeUsed >= FREE_TIER_TIME_LIMIT {
		log.Printf("not running process for user %d because they maxed TimeUsed", u.ID)
		return false;
	}

	var timeLeft int64 = FREE_TIER_TIME_LIMIT - u.TimeUsed
	log.Printf("Creating a timer for user %d with limit %ds", u.ID, timeLeft)
	// TODO: add 1ms to fix the race on the equal case.
	noTimeLeftTimer := time.NewTimer(time.Duration(timeLeft) * time.Second)

	notifyProcessDone := make(chan int64)
	go ProcessNotifier(process, notifyProcessDone)

	select {
	case <- noTimeLeftTimer.C:
		log.Printf("noTimeLeftTimer ran out for user %d", u.ID)
		u.TimeUsed = 10;
		return false
	case elapsedSeconds := <- notifyProcessDone:
		u.TimeUsed += elapsedSeconds
		log.Printf("user %d TimeUsed increased to %d", u.ID, u.TimeUsed)
		return true
	}
}

/** Wraps a process, and notifys when it completes with the time taken. */
func ProcessNotifier(process func(), notifyDone chan int64) {
	start := time.Now()
	process()
	elapsed := time.Since(start)
	var elapsedSeconds int64 = int64(elapsed / time.Second)
	// log.Printf("process took %d", elapsedSeconds)
	notifyDone <- elapsedSeconds;
}

func main() {
	RunMockServer()
}
