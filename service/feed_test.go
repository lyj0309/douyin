package service

import (
	"testing"
	"time"
)

func TestFeed(t *testing.T) {
	Feed(time.Now(), "")
}
