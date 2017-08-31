package wheel

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type CreateInfluxDB struct {
	Name      string
	Retention string
	PVCSize   string
	ChartURL  string `json:"chart_url"`
	Values    struct {
		Raw string `json:"raw"`
	} `json:"values"`
}

type CreateGrafana struct {
	Name     string
	ChartURL string `json:"chart_url"`
}

// WheelInfluxdb create influxdb with wheel
func (c *CreateInfluxDB) WheelInfluxdb() error {
	releaseURL := "http://wheel.kube-system:9855/tiller/v2/releases/" + c.Name + "-i/json"
	c.ChartURL = "https://github.com/orangesys/charts/raw/master/docs/influxdb-0.1.13.tgz"
	c.Values.Raw = "{\"retentionPolicy\":" + c.Retention + ",\"persistence.size\":" + c.PVCSize + "}"
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(c)
	_, err := http.Post(releaseURL, "application/json; charset=utf-8", b)
	if err != nil {
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
