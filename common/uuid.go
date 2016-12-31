package common

import (
    "fmt"
    "io"
    "crypto/rand"
)

// newUUID generates a random UUID according to RFC 4122
func UUID() (string, error) {
    uuid := make([]byte, 16)
    n, err := io.ReadFull(rand.Reader, uuid)
    if n != len(uuid) || err != nil {
        return "", err
    }
    // variant bits; see section 4.1.1
    uuid[8] = uuid[8]&^0xc0 | 0x80
    // version 4 (pseudo-random); see section 4.1.3
    uuid[6] = uuid[6]&^0xf0 | 0x40
    //return fmt.Sprintf("%x", uuid), nil
    return fmt.Sprintf("%x", uuid), nil
}
