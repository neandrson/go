package main

import (
	"testing"
	"time"
)

func TestTimeAgo(t *testing.T) {
	pastTime := time.Now().Add(-2 * time.Hour)
	timeAgo := TimeAgo(pastTime)
	expected := "2 hours ago"
	if timeAgo != expected {
		t.Errorf("Expected time ago string to be %s, but got %s", expected, timeAgo)
	}
}
