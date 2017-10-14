package common

import "testing"

func TestUUID(t *testing.T) {
	want := 32
	got, _ := UUID()
	if len(got) == want {
		return
	}
	t.Errorf("UUID() = %v, want %v", got, want)
}
