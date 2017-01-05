package helm

import (
	_ "fmt"
	_ "os"
	"os/exec"
)

type InfluxdbCommander struct {
	Name, Retention, Pvcsize string
}

type GrafanaCommander struct {
	Name string
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
