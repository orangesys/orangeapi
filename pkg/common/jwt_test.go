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
		{
			"name",
			fields{
				"a36c3049b36249a3c9f8891cb127243c",
				"e71829c351aa4242c2719cbfbe671c09",
			},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhMzZjMzA0OWIzNjI0OWEzYzlmODg5MWNiMTI3MjQzYyJ9.U8dOyd1978lmbWNk7gXHf7krDTjYKZanrVpayA0Lhug",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Consumer{
				Iss:    tt.fields.Iss,
				Secret: tt.fields.Secret,
			}
			got, _ := c.CreateToken()
			if got != tt.want {
				t.Errorf("Consumer.CreateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
