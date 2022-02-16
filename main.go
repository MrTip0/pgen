package main

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	pass := []rune{}
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#@*%&?-_~.")

	numbers := make([]byte, 20)
	if _, err := io.ReadFull(rand.Reader, numbers); err != nil {
		fmt.Printf("An error occured: %s", err.Error())
	}

	max := len(chars)

	for i := 0; i < 20; i++ {
		pass = append(pass, chars[int(numbers[i]) % max])
	}

	if !CheckSpecialandNumber(pass) {
		el := 2
		pn, pc := numbers[0] % 20, numbers[1] % 20
		for pn == pc {
			pn = numbers[el] % 20
			el ++
			el = el % 20
		}
		pass[pn] = rune((numbers[el] % 10) + '0')
		el ++
		el = el % 20

		schars := []rune("!#@*%&?-_~.")
		pass[pc] = schars[numbers[el] % 10]
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