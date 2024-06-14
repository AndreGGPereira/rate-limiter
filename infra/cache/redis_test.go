package cache

import (
	"context"
	"testing"
	"time"
)

func TestIncr(t *testing.T) {

	ctx := context.Background()

	cache, err := NewConnection(
		ctx,
		"localhost",
		"6379",
		"",
		1,
	)
	if err != nil {
		t.Errorf("failed connection")
	}

	count, err := cache.Incr(ctx, "testKey")
	if err != nil {
		t.Errorf("failed incr")
	}
	if count == 0 {
		t.Errorf("failed incr")
	}
}

func TestSet(t *testing.T) {

	ctx := context.Background()

	cache, err := NewConnection(
		ctx,
		"localhost",
		"6379",
		"",
		1,
	)
	if err != nil {
		t.Errorf("failed connection")
	}

	if err := cache.Set(ctx, "testKey", "val", 0); err != nil {
		t.Errorf("failed Set")
	}
}

func TestExpire(t *testing.T) {

	ctx := context.Background()

	cache, err := NewConnection(
		ctx,
		"localhost",
		"6379",
		"",
		1,
	)
	if err != nil {
		t.Errorf("failed connection")
	}

	ok, err := cache.Expire(ctx, "testKey", time.Duration(10))
	if err != nil {
		t.Errorf("failed expire")
	}
	if !ok {
		t.Errorf("failed expire")
	}
}

func TestGet(t *testing.T) {

	ctx := context.Background()

	cache, err := NewConnection(
		ctx,
		"localhost",
		"6379",
		"",
		1,
	)
	if err != nil {
		t.Errorf("failed connection")
	}

	if err := cache.Set(ctx, "testKey", "val", 0); err != nil {
		t.Errorf("failed Set")
	}

	val, err := cache.Get(ctx, "testKey")
	if err != nil {
		t.Errorf("failed get")
	}
	if val == "" {
		t.Errorf("failed get")
	}
}

func TestClose(t *testing.T) {

	ctx := context.Background()

	cache, err := NewConnection(
		ctx,
		"localhost",
		"6379",
		"",
		1,
	)
	if err != nil {
		t.Errorf("failed connection")
	}
	if err := cache.Close(); err != nil {
		t.Errorf("faild close")
	}
}
