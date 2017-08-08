package helm

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os/exec"
)

type InfluxdbCommander struct {
	Name, Retention, Pvcsize string
}

type GrafanaCommander struct {
	Name string
}

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
	ChartURL string
}

// WheelInfluxdb create influxdb with wheel
func (c *CreateInfluxDB) WheelInfluxdb() error {
	releaseURL := "http://wheel.kube-system:9855/tiller/v2/releases/" + c.Name + "-i/json"
	c.ChartURL = "https://github.com/orangesys/charts/raw/master/docs/influxdb-0.1.12.tgz"
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
	c.ChartURL = "https://github.com/orangesys/charts/raw/master/docs/grafana-0.1.18.tgz"
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(c)
	_, err := http.Post(releaseURL, "application/json; charset=utf-8", b)
	if err != nil {
		return err
	}
	return nil
}

func (i *InfluxdbCommander) InstallInfluxdb() error {
	releasename := "--name=" + i.Name + "-i"
	opt := "retentionPolicy=" + i.Retention + ",persistence.size=" + i.Pvcsize
	cmdName := "helm"
	cmdArgs := []string{
		"install",
		"--namespace=default",
		releasename,
		"or-charts/influxdb",
		"--set",
		opt,
	}
	return exec.Command(cmdName, cmdArgs...).Run()
}

func (g *GrafanaCommander) InstallGrafana() error {
	releasename := "--name=" + g.Name + "-g"
	cmdName := "helm"
	cmdArgs := []string{
		"install",
		"--namespace=default",
		releasename,
		"or-charts/grafana",
	}
	return exec.Command(cmdName, cmdArgs...).Run()
}

//func main()  {
//  name := "rlxebz"
//  var err error
//
//  influxdb := InfluxdbCommander{
//    Name: name,
//    Retention: "10d",
//    Pvcsize: "10Gi",
//  }
//  err = influxdb.InstallInfluxdb()
//  if err != nil {
//    fmt.Printf("can not create influxdb%+v", err)
//    os.Exit(98)
//  }
//
//  grafana := GrafanaCommander{ Name: name }
//  err = grafana.InstallGrafana()
//  if err != nil {
//    fmt.Printf("can not create grafana%+v", err)
//    os.Exit(99)
//  }
//}
