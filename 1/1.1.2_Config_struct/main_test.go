package config

import "testing"

func TestConfig(t *testing.T) {
	cfg := New()

	if cfg.Height != 50 {
		t.Errorf("expected height: %d, actual height: %d", 50, cfg.Height)
	}

	if cfg.Width != 100 {
		t.Errorf("expected width: %d, actual width: %d", 100, cfg.Width)
	}
}
