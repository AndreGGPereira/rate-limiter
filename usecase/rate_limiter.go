package usecase

import (
	"context"
	"fmt"
	"rate-limiter/config"
	"rate-limiter/infra/cache"
	"strconv"
	"time"
)

type RateLimitToken struct {
	Token      string
	Limiter    int
	Expiration int
}

type RateLimiter struct {
	Cache cache.Repository
}

func NewRateLimiter(ch cache.Repository) *RateLimiter {
	return &RateLimiter{
		Cache: ch,
	}
}

func (rl *RateLimiter) AllowRequest(ctx context.Context, key string, isToken bool) bool {

	if isToken {

		key = fmt.Sprintf("block:%s", key)

		tokenLimiter := configToken(key)
		if (RateLimitToken{} == tokenLimiter) {
			return false
		}
		if rl.IsBlocked(ctx, key) {
			return false
		}

		count, err := rl.Cache.Incr(ctx, key)
		if err != nil {
			return false
		}

		if count == 1 {
			rl.Cache.Expire(ctx, key, time.Duration(tokenLimiter.Expiration)*time.Second)
		}

		if count > int64(tokenLimiter.Limiter) {
			rl.Cache.Set(ctx, fmt.Sprintf("block:%s", key), "blocked", time.Duration(config.Limiter().BlockingTimeToken)*time.Second)
			return false
		}
		return true

	} else {

		if rl.IsBlocked(ctx, key) {
			return false
		}

		count, err := rl.Cache.Incr(ctx, key)
		if err != nil {
			return false
		}

		if count == 1 {
			rl.Cache.Expire(ctx, key, time.Duration(config.Limiter().Expiration)*time.Second)
		}
		if count > int64(config.Limiter().LimitRequestPerIp) {
			rl.Cache.Set(ctx, fmt.Sprintf("block:%s", key), "blocked", time.Duration(config.Limiter().BlockingTimeIP)*time.Second)
			return false
		}
		return true
	}

}

func (rl *RateLimiter) IsBlocked(ctx context.Context, key string) bool {
	val, err := rl.Cache.Get(ctx, fmt.Sprintf("block:%s", key))
	return err == nil && val == "blocked"
}

func (rl *RateLimiter) Reset(ctx context.Context, key string) {
	rl.Cache.Set(ctx, key, "0", 0)
	rl.Cache.Set(ctx, fmt.Sprintf("block:%s", key), "", 0)
}

func (rl *RateLimiter) Count(ctx context.Context, key string) (int, bool) {
	count, err := rl.Cache.Get(ctx, key)
	if err != nil {
		if err.Error() == "redis: nil" {
			return -1, true
		}
	}

	val, err := strconv.Atoi(count)
	if err != nil {
		return -1, false
	}

	return val, false
}

func configToken(token string) RateLimitToken {
	for _, v := range config.Token() {
		if token == v.Token {
			return RateLimitToken{
				Token:      v.Token,
				Limiter:    v.Limiter,
				Expiration: v.Expiration,
			}
		}
	}
	return RateLimitToken{}
}
