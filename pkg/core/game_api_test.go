package core

import (
	"testing"
	"time"
)

func TestSecondsToServerISO(t *testing.T) {
	now := time.Unix(1765733352, 0)
	testee := SecondsToServerISO(now, "Asia/Seoul")
	expected := "20251215022912"
	if testee != expected {
		t.Errorf("Expected %s got %s", expected, testee)
	}
}
