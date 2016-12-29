package k8s

import (
  _ "fmt"
  "errors"
  "encoding/base64"
)

func GetSecret (namespace, key, data string) (string, error){
  s := &Secret{}

  s, err := getSecret(namespace, key)
  if err != nil {
    return "", errors.New("missing key" + key)
  }
  v := s.Data[data]
  d, err := base64.StdEncoding.DecodeString(v)
  if err != nil {
    return "", err
  }
  return string(d), nil
}
