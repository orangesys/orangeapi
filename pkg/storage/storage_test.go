package storage

import (
	"reflect"
	"testing"

	"github.com/influxdata/influxdb/client/v2"
)

func TestInfluxDBClient(t *testing.T) {
	type args struct {
		consumerID string
	}
	tests := []struct {
		name string
		args args
		want client.Client
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InfluxDBClient(tt.args.consumerID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InfluxDBClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStorageUsed(t *testing.T) {
	type args struct {
		c client.Client
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStorageUsed(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStorageUsed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetStorageUsed() = %v, want %v", got, tt.want)
			}
		})
	}
}
