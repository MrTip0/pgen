package main

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	pass := []rune{}
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#@*%&?-_~")

	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	max := len(chars)

	for i := 0; i < 20; i++ {
		pass = append(pass, chars[rand.Int() % max])
	}

	if !CheckSpecialandNumber(pass) {
		pn, pc := rand.Int() % 20, rand.Int() % 20
		for pn == pc {
			pn = rand.Int() % 20
		}
		pass[pn] = rune((rand.Int() % 10) + '0')

		schars := []rune("!#@*%&?-_~")
		pass[pc] = schars[rand.Int() % 10]
	}

	spass := string(pass)
	fmt.Printf("Password: %s\nSHA1: %x\n", spass, sha1.Sum([]byte(spass)))
}

func CheckSpecialandNumber(s []rune) bool {
	n, c := false, false
	l := len(s)

	schars := []rune("!#@*%&?-_~")
	ll := len(schars)

	for i := 0; i < l && (!n || !c); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			n = true
		} else {
			for j := 0; j < ll && !c; j++ {
				if schars[j] == s[i] {
					c = true
				}
			}
		}
	}
	return n && c
}