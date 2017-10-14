package common

import "testing"

func TestReleaseName(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"name", 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReleaseName(); len(got) != tt.want {
				t.Errorf("ReleaseName() = %v, want %v", got, tt.want)
			}
		})
	}
}
