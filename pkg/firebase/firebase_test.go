package firebase

import (
	_ "os"
	"testing"

	"github.com/orangesys/orangeapi/pkg/config"
)

func TestFirebaseConfiguration_CheckUser(t *testing.T) {
	type fields struct {
		Config     *config.FirebaseConfiguration
		UUID       string
		ConsumerID string
		Token      string
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
			f := &FirebaseConfiguration{
				Config:     tt.fields.Config,
				UUID:       tt.fields.UUID,
				ConsumerID: tt.fields.ConsumerID,
				Token:      tt.fields.Token,
			}
			if err := f.CheckUser(); (err != nil) != tt.wantErr {
				t.Errorf("FirebaseConfiguration.CheckUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
