package firebase

import (
    "github.com/JustinTulloss/firebase"
    "github.com/orangesys/orangeapi/config"
)

type Date struct {
    updatedAt string
}

func checkuser(endpoint, auth, uuid string) bool {
    c := firebase.NewClient(endpoint + "/users/" + uuid, auth, nil)
    var r map[string]interface{}
    err := c.Value(&r)
    if err != nil {
        return false
    }
    if r == nil {
        return false
    }
    return true
}

func savetoken(endpoint, auth, uuid, consumerId, token string) bool {
    c := firebase.NewClient(endpoint + "/users/" + uuid, auth, nil)
    tf := map[string]interface{}{ "consumerID": consumerId, "token": token }
    _, err := c.Set("telegraf", tf, nil)

    if err != nil {
      return false
    }
    return true
}
