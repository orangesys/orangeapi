package k8s

import (
  _ "fmt"
  "encoding/base64"
)

func GetSecret (namespace, key, data string) (string, error){
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
