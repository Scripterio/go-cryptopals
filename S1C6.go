package main

import (
	"fmt"
	"github.com/scripterio/go-cryptopals/base64"
	"github.com/scripterio/go-cryptopals/bitwise"
	"github.com/scripterio/go-cryptopals/hexstring"
	"github.com/scripterio/go-cryptopals/io"
)

func main() {

	file := io.Read("./data/6.txt")

	raw := base64.Decode(file)

	//fmt.Printf(hexstring.Encode(raw))

	KEYSIZE := 2
	bestkey := 0
	ham := 1000000 // Arbitaray large value
	for KEYSIZE < 40 {

		slice1 := raw[0:KEYSIZE]
		slice2 := raw[KEYSIZE:(KEYSIZE * 2)]

		distance := (bitwise.HAM(slice1, slice2) / KEYSIZE)

		if distance <= ham {
			ham = distance
			bestkey = KEYSIZE

		}
		KEYSIZE++
	}
	print("KEYSIZE: ", bestkey, " HAM: ", ham, "\n")

	KEYSIZE = bestkey

	k := 0

	var finalkey []byte

	for k < KEYSIZE {
		singlekey := 0
		l := 0
		var m []byte

		for l < (len(raw) / KEYSIZE) {

			if (l * KEYSIZE) < (len(raw) - 1) {
				m = append(m, raw[(l*KEYSIZE)+k])
			}

			l++

		}

		bruteout := bitwise.BRUTEXOR(m)

		_, insidekey := hexstring.ScoreEnglish(bruteout)
		singlekey = insidekey
		finalkey = append(finalkey, byte(singlekey))
		//fmt.Printf("%s", bitwise.SBXOR(m, byte(singlekey)))

		m = nil
		k++
	}
	fmt.Printf("-----------------------------------------------------\n")
	fmt.Printf("%s\n", finalkey)
	fmt.Printf("-----------------------------------------------------")
	print("\n")

	//finaloutput := bitwise.RXOR(raw, finalkey)

	//print(string(finaloutput))

	// Hamming Test
	input1 := "@ @ @ @ @ @ @ @ @ !"
	input2 := " @ @ @ @ @ @ @ @ @ "

	testhamming := bitwise.HAM([]byte(input1), []byte(input2))

	if testhamming == 37 {
		print("\n\nPass Hamming Test\n\n")
	} else {
		print("\n\n###FAIL###\n\n")
	}

	//Test
	//fmt.Printf("Set 1   Challenge 6 \n")

	testinput := "f0f0f0"

	testraw := hexstring.Decode(testinput)

	testoutput := base64.Encode(testraw)

	testreverse := base64.Decode(testoutput)

	testcooked := hexstring.Encode(testreverse)

	if testcooked == testinput {

		//fmt.Printf("DING!!!!\n\n")
	}

}
