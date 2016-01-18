package base64

import b64 "encoding/base64"

//Takes bytes, returns the b64
func Encode(b []byte) string {

	sEnc := b64.StdEncoding.EncodeToString([]byte(b))

	return string(sEnc)
}
