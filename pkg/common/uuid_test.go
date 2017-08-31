package common

import "testing"

func TestUUID(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UUID()
			if (err != nil) != tt.wantErr {
				t.Errorf("UUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
