package storage

import (
	"encoding/json"
	"fmt"

	"github.com/influxdata/influxdb/client/v2"
	log "github.com/rs/zerolog/log"
)

const (
	database = "_internal"
)

// InfluxDBClient init new client
func InfluxDBClient(consumerID string) client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://" + consumerID + "-i-influxdb.default",
	})
	if err != nil {
		log.Error().Msgf("can not init new influxdb client: %v", err)
	}
	return c
}

// GetStorageUsed get used storage
func GetStorageUsed(c client.Client) (int64, error) {
	q := client.Query{
		Command:  fmt.Sprintf("select sum(diskBytes) FROM tsm1_filestore WHERE time > now() - 30s GROUP BY time(2s) fill(none) limit 1;"),
		Database: database,
	}
	resp, err := c.Query(q)
	if err != nil {
		return 0, err
	}
	if resp.Error() != nil {
		return 0, err
	}

	res, err := resp.Results[0].Series[0].Values[0][1].(json.Number).Int64()
	if err != nil {
		return 0, err
	}

	return res, nil
}
