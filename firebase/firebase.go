package main

import (
    "fmt"
    "os"

    "github.com/JustinTulloss/firebase"
    "github.com/orangesys/orangeapi/config"
)

func CheckUser(uuid string, config *config.FirebaseConfiguration) error {
    c := firebase.NewClient(config.FirebaseURL + "/users/" + uuid, config.FirebaseAuth, nil)
    var r map[string]interface{}
    err := c.Value(&r)
    if err != nil {
        return err
    }
    if r == nil {
        return fmt.Errorf("%s %s", "can not get consumer", uuid)
    }
    return nil
}

func SaveToken(uuid, consumerId, token string, config *config.FirebaseConfiguration) error {
    c := firebase.NewClient(config.FirebaseURL + "/users/" + uuid, config.FirebaseAuth, nil)
    tf := map[string]interface{}{ "consumerID": consumerId, "token": token }
    _, err := c.Set("telegraf", tf, nil)

    if err != nil {
	return err
    }
    return nil
}

func main() {
    config, _ := config.LoadFirebaseConfig()
    fmt.Println(config)
    err := CheckUser("iGzNX6QzfudVlwKtR8CQCj0itIU2", config)
    if err !=nil {
        fmt.Println(err)
	os.Exit(1)
    }
    err = SaveToken("iGzNX6QzfudVlwKtR8CQCj0itIU2", "test", "testtest", config)
    if err !=nil {
        fmt.Println(err)
	os.Exit(2)
    }
}
