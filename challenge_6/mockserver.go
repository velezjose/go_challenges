//////////////////////////////////////////////////////////////////////
//
// Note: I obtained this exercise from https://github.com/loong/go-concurrency-exercises
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go` file.
//

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// RunMockServer pretends to be a video processing service. It
// simulates a user interacting with the Server.
func RunMockServer() {
	u1 := User{ ID: 0, IsPremium: false }
	u2 := User{ ID: 1, IsPremium: true }

	wg.Add(5)

	go createMockRequest(1, shortProcess, &u1) // user 1:  current: 6s   total: 6s
	time.Sleep(1 * time.Second)
	time.Sleep(7 * time.Second)

	go createMockRequest(2, longProcess, &u2)  // user 2:  current: 11s  total: 11s
	time.Sleep(2 * time.Second)

	go createMockRequest(3, shortProcess, &u1) // user 1:  current: 6s   total: 12s
	time.Sleep(1 * time.Second)
	time.Sleep(5 * time.Second)

	go createMockRequest(4, longProcess, &u1)  // user 1:  current: 11s  total: 23s

	go createMockRequest(5, shortProcess, &u2) // user 2:  current: 6s   total: 17s

	wg.Wait()
}

func createMockRequest(pid int, fn func(), u *User) {
	fmt.Println("UserID:", u.ID, "\tProcess", pid, "started.")
	res := HandleRequest(fn, u)

	if res {
		fmt.Println("UserID:", u.ID, "\tProcess", pid, "DONE.")
	} else {
		fmt.Println("UserID:", u.ID, "\tProcess", pid, "killed. (No quota left)")
	}

	wg.Done()
}

func shortProcess() {
	time.Sleep(6 * time.Second)
}

func longProcess() {
	time.Sleep(11 * time.Second)
}
