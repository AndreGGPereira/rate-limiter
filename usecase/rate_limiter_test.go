package usecase

import (
	"context"
	"log"
	"rate-limiter/config"
	"rate-limiter/infra/cache"

	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
)

func TestRate(t *testing.T) {

	ctx := context.Background()
	s, err := miniredis.Run()
	if err != nil {
		t.Fatalf("Failed to start miniredis: %v", err)
	}
	defer s.Close()

	if err := config.Load("./../config/config.yaml"); err != nil {
		log.Fatalf("Failed get config [%s]", err)
	}

	conn, err := cache.NewConnection(ctx, "localhost", "6379", "", 1)
	if err != nil {
		t.Errorf("failed opem connection %s", err.Error())
	}

	rl := NewRateLimiter(conn)

	ip := "127.0.0.1"

	for i := 0; i < 5; i++ {
		if !rl.AllowRequest(ctx, ip, false) {
			t.Fatalf("Request %d should have been allowed (IP)", i+1)
		}
	}

	if rl.AllowRequest(ctx, ip, false) {
		t.Fatalf("6th request should have been blocked (IP)")
	}

	if !rl.IsBlocked(ctx, ip) {
		t.Fatalf("IP should be blocked")
	}

	if count, ok := rl.Count(ctx, ip); ok {
		t.Logf("IP Request count before sleep: %d, %v", count, ok)
	}

	time.Sleep(10 * time.Second)

	if count, ok := rl.Count(ctx, ip); ok {
		t.Logf("IP Request count after sleep: %d, %v", count, ok)
	}
}
