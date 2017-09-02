package wheel

import (
	"testing"
)

func TestCreateGrafana_WheelGrafana(t *testing.T) {
	type fields struct {
		Name     string
		ChartURL string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateGrafana{
				Name:     tt.fields.Name,
				ChartURL: tt.fields.ChartURL,
			}
			if err := c.WheelGrafana(); (err != nil) != tt.wantErr {
				t.Errorf("CreateGrafana.WheelGrafana() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateInfluxDB_WheelInfluxdb(t *testing.T) {
	type fields struct {
		Name     string
		ChartURL string
		Values   Values
	}
	type args struct {
		retention string
		pvcsize   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CreateInfluxDB{
				Name:     tt.fields.Name,
				ChartURL: tt.fields.ChartURL,
				Values:   tt.fields.Values,
			}
			if err := c.WheelInfluxdb(tt.args.retention, tt.args.pvcsize); (err != nil) != tt.wantErr {
				t.Errorf("CreateInfluxDB.WheelInfluxdb() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
