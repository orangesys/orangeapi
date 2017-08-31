package k8s

import "testing"

func TestGetSecret(t *testing.T) {
	type args struct {
		namespace string
		key       string
		data      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSecret(tt.args.namespace, tt.args.key, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
