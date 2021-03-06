package k8s

import (
	"encoding/base64"
)

// GetSecret with secret from k8s
func GetSecret(namespace, key, data string) (string, error) {
	s, err := getSecret(namespace, key)
	if err != nil {
		return "", err
	}
	v := s.Data[data]
	d, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return "", err
	}
	return string(d), nil
}
