package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func getHMAC(s, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	_, err := io.WriteString(h, s)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	fmt.Printf("enter text:\n> ")
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Printf("Hashed:\n> %v\n", getHMAC(input, "PASSWORD"))
}
