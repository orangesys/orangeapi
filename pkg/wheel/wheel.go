package wheel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateInfluxDB struct {
	Name     string
	ChartURL string `json:"chart_url"`
	Values   Values `json:"values,omitempty"`
}

type Values struct {
	Values string `json:"raw"`
}

type Raw struct {
	Persistence     Persistence `json:"persistence"`
	RetentionPolicy string      `json:"retentionPolicy"`
}

type Persistence struct {
	Size string `json:"size"`
}

type CreateGrafana struct {
	Name     string
	ChartURL string `json:"chart_url"`
}

// WheelInfluxdb create influxdb with wheel
func (c *CreateInfluxDB) WheelInfluxdb(retention, pvcsize string) error {
	releaseURL := "http://wheel.kube-system:9855/tiller/v2/releases/" + c.Name + "-i/json"

	p := Persistence{Size: pvcsize}
	r := Raw{Persistence: p, RetentionPolicy: retention}
	rb, _ := json.Marshal(r)
	c.Values.Values = string(rb)

	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(data))

	req, err := http.NewRequest("POST", releaseURL, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

// WheelGrafana create grafana with wheel
func (c CreateGrafana) WheelGrafana() error {
	releaseURL := "http://wheel.kube-system:9855/tiller/v2/releases/" + c.Name + "-g/json"
	c.ChartURL = "https://github.com/orangesys/charts/raw/master/docs/grafana-0.1.21.tgz"
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(c)
	_, err := http.Post(releaseURL, "application/json; charset=utf-8", b)
	if err != nil {
		return err
	}
	return nil
}
