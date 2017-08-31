package config

import (
	_ "fmt"
	"reflect"
	"testing"
)

func TestLoadKongConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *KongConfiguration
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadKongConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadKongConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadKongConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
