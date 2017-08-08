package common

import (
	"crypto/rand"
)

func ReleaseName() string {
	const alphanum = "abcdefghijkmnopqrstuvwxyz"
	var bytes = make([]byte, 6)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
