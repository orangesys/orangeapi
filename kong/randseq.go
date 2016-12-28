package kong

import (
    "crypto/rand"
)

func randSeq(n int) string {
    const alphanum = "abcdefghijkmnopqrstuvwxyz"
    var bytes = make([]byte, n)
    rand.Read(bytes)
    for i,b := range bytes {
        bytes[i] = alphanum[b % byte(len(alphanum))]
    }
    return string(bytes)
}
