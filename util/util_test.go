package util

import (
	"testing"
)

func TestGetUpTime(t *testing.T) {
	uptime := GetUpTime()
	if uptime == "" {
		t.Errorf("Expected a value, but got empty")
	}
}
