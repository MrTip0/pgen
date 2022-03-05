package main

import (
	"crypto/rand"
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	pass := []rune{}
	alfabeth := string("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#@*%&?-_~.")
	plen := 20

	flag.StringVar(&alfabeth, "alfa", alfabeth, "the alfabeth")
	flag.IntVar(&plen, "l", plen, "the lenght of the password")

	flag.Parse()

	numbers := make([]byte, plen)
	if _, err := io.ReadFull(rand.Reader, numbers); err != nil {
		fmt.Printf("An error occured: %s", err.Error())
		os.Exit(1)
	}

	chars := []rune(alfabeth)

	max := len(chars)

	for i := 0; i < plen; i++ {
		pass = append(pass, chars[int(numbers[i]) % max])
	}

	if !CheckSpecialandNumber(pass) {
		el := 2
		plenb := byte(plen)
		pn, pc := numbers[0] % plenb, numbers[1] % plenb
		for pn == pc {
			pn = numbers[el] % plenb
			el ++
			el = el % plen
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