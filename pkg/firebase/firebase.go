package firebase

import (
	"fmt"
	_ "os"

	"github.com/JustinTulloss/firebase"
	"github.com/orangesys/orangeapi/pkg/config"
)

type FirebaseConfiguration struct {
	Config     *config.FirebaseConfiguration
	UUID       string
	ConsumerID string
	Token      string
}

func (f *FirebaseConfiguration) CheckUser() error {
	c := firebase.NewClient(f.Config.FirebaseURL+"/users/"+f.UUID, f.Config.FirebaseAuth, nil)
	var r map[string]interface{}
	err := c.Value(&r)
	if err != nil {
		return err
	}
	if r == nil {
		return fmt.Errorf("%s %s", "can not get consumer", f.UUID)
	}
	return nil
}

func (f *FirebaseConfiguration) SaveToken() error {
	c := firebase.NewClient(f.Config.FirebaseURL+"/users/"+f.UUID, f.Config.FirebaseAuth, nil)
	tf := map[string]interface{}{"consumerId": f.ConsumerID, "token": f.Token}
	_, err := c.Set("telegraf", tf, nil)

	if err != nil {
		return err
	}
	return nil
}

//func main() {
//    config, err := config.LoadFirebaseConfig()
//    if err != nil {
//        fmt.Println(err)
//	os.Exit(1)
//    }
//    user := FirebaseConfiguration{
//	Config: config,
//	UUID: "iGzNX6QzfudVlwKtR8CQCj0itIU2",
//	ConsumerID: "test",
//	Token: "testtest",
//    }
//    err = user.CheckUser()
//    if err !=nil {
//        fmt.Println(err)
//	os.Exit(1)
//    }
//    err = user.SaveToken()
//    if err !=nil {
//        fmt.Println(err)
//	os.Exit(2)
//    }
//}
