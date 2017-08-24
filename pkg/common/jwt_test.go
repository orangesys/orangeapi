package common

import "testing"

func TestConsumer_CreateToken(t *testing.T) {
	type fields struct {
		Iss    string
		Secret string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Consumer{
				Iss:    tt.fields.Iss,
				Secret: tt.fields.Secret,
			}
			got, err := c.CreateToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("Consumer.CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Consumer.CreateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
