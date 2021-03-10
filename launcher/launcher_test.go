package launcher

import "testing"

func TestNewLauncherDefault(t *testing.T) {
	l, err := NewLauncherDefault()
	if err != nil {
		t.Error(err)
	}
	_ = l
}
