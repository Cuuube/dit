package cli

import (
	"testing"
	"time"
)

func TestLoading(t *testing.T) {
	var loading Loading = NewSimpleLoading("Loading", WithFrameDuration(time.Second/10))
	go loading.Play()
	go loading.Play()
	go loading.Play()
	time.Sleep(10 * time.Second)
	loading.Stop()
}
