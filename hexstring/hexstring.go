package hexstring

import (
	"encoding/hex"
	"fmt"
	"os"
)

//Takes hex, returns the b64
func Decode(s string) []byte {

	raw, err := hex.DecodeString(s)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return (raw)
}

func Encode(b []byte) string {

	s := hex.EncodeToString(b)

	return (s)
}

func ScoreEnglish(b [][]byte) (score, key int) {

	// Somewhat arbitrary score assigned to each character
	// " " 	:= 13
	//	e 	:= 12
	//	t 	:= 9
	//	a 	:= 8
	//	o 	:= 7
	//	i 	:= 7
	//	n 	:= 7
	//	s 	:= 6

	// Since we are receiving an array of arrays that were bruteforce XORd,
	// we can infer that the first arrary was XORd with \x00.
	// y is keeping track of that.
	// Im sure there is a more graceful way of doing this...

	score = 0
	y := 0
	for _, i := range b {
		x := 0

		for _, j := range i {

			if j < byte('\x40') || j > byte('\x7A') {

				x = -10

			}
			if j == byte('\x20') {

				x += 20
			}

			if j == byte('\x45') || j == byte('\x65') {

				x += 10

			}

			if j == byte('\x54') || j == byte('\x74') {

				x += 9

			}

			if j == byte('\x41') || j == byte('\x61') {

				x += 8

			}

			if j == byte('\x4F') || j == byte('\x6F') {

				x += 7

			}

			if j == byte('\x49') || j == byte('\x69') {

				x += 7

			}

			if j == byte('\x4E') || j == byte('\x6E') {

				x += 7

			}

			if j == byte('\x53') || j == byte('\x73') {

				x += 6

			}

			if x > score {
				score = x
				key = y

			}

		}
		y++

	}

	return

}
