package storage

import (
    "encoding/json"
    "fmt"
    "github.com/influxdata/influxdb/client/v2"
    "log"
)

const (
    database = "_internal"
)

func InfluxDBClient(consumerId string) client.Client {
    c, err := client.NewHTTPClient(client.HTTPConfig{
        Addr:     "http://" + consumerId + "-i-influxdb.default",
    })
    if err != nil {
        log.Fatalln("Error: ", err)
    }
    return c
}

func GetStorageUsed(c client.Client) (int64, error) {
    q := client.Query{
        Command: fmt.Sprintf("select sum(diskBytes) FROM tsm1_filestore WHERE time > now() - 19s"),
        Database: database,
    }
    resp, err := c.Query(q)
    if err != nil {
        return "", err
    }
    if resp.Error() != nil {
        return "", err
    }

    res, err := resp.Results[0].Series[0].Values[0][1].(json.Number).Int64()
    if err != nil {
        return "", err
    }

    return res, nil
}

//func main() {
//    consumerId := "test1"
//    c := InfluxDBClient(consumerId)
//    log.Printf("Mean values: diskBytes='%d'", GetStorageUsed(c))
//}
