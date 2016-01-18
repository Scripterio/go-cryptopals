package main

import (
	"fmt"
	"github.com/scripterio/go-cryptopals/base64"
	"github.com/scripterio/go-cryptopals/hexstring"
)

func main() {

	input := "f0f0f0"

	raw := hexstring.Decode(input)

	output := base64.Encode(raw)

	//Test
	fmt.Printf("Set 1   Challenge 1 \n")
	if output == "8PDw" {

		fmt.Printf("PASS!!\n\n")

	} else {
		fmt.Printf("##FAIL##\n\n")
	}

}
