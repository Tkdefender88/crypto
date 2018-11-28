package main

import (
	"fmt"

	"github.com/Tkdefender88/crypto/caesar"
)

func main() {
	fmt.Println(caesar.Encode("hello world", 7))
}
