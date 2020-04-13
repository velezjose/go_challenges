package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

var chars = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V",
	"W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r",
	"s", "t", "u", "v", "w", "x", "y", "z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "!", "@", "#", "$",
	"%", "^", "&", "*", "(", ")", ";", ",", "'", ".", "/", "?",
}

func crack(hashString string) string {
	// your code here - change this method!
	return hashString
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	/**
	 * Password Cracking:
	 *  Your task is to design a concurrent algorithm that can crack the 3 password hashes below.
	 *  These are the equations you're dealing with (related to the variable names below):
	 *    MD5(a) = hash1
	 *    MD5(b) = hash2
	 *    MD5(c) = hash3
	 *    MD5(a + b + c) = MD5(original password) = hashOfPwd
	 *    original password = str1 + str2 + str3
	 *
	 *  Hints:
	 *   1. Use the chars slice above and the getMD5Hash(text string) function provided.
	 *   2. Make sure you can crack short passwords before trying to crack the password for this exercise, to
	 *      to make sure your algorithm works. For example, try to crack the password behind:
	 *          "0cc175b9c0f1b6a831c399e269772661" = MD5("a")   –> should retrieve "a" given that hash
	 *          "900150983cd24fb0d6963f7d28e17f72" = MD5("abc") -> should retrieve "abc" given the hash
	 *   3. You should have at least 3 goroutines running concurrently.
	 *
	 *  If your algorithm is taking more than 2 minutes, you're probably not taking full advantage of concurrency
	 *  or you might be doing repeated work.
	 *
	*/
	start := time.Now()
	hash1 := "90b4cf5918aed1be1ce9b9e3bd8200d4"
	hash2 := "74359efeda2ff9a8afe3185e6b3d4aa3"
	hash3 := "950fcab7f816eb9bdb3b8e7b557fabed"
	hashOfPwd := "d7b287291140f42c558a7dcc33ffd75b"

	////////////////////////////////////
	// your code in here, and above in the "crack" function. you can create other functions
	password := crack(hash1) + crack(hash2) + crack(hash3)
	//
	////////////////////////////////////

	if hashOfPwd == getMD5Hash(password) {
		fmt.Println("You just won a", password)
	} else {
		fmt.Println("You haven't cracked the password :(")
	}

	fmt.Println("Total process took", time.Since(start))
}
