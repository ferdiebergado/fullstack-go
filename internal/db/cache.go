package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	redisClient *redis.Client
	dbQueries   *Queries // Your generated sqlc queries
}

// NewCache initializes the Cache struct with a Redis client and the sqlc queries.
func NewCache(redisClient *redis.Client, dbQueries *Queries) *Cache {
	return &Cache{
		redisClient: redisClient,
		dbQueries:   dbQueries,
	}
}

// cacheFetch is the generic function to handle caching logic.
// - fetchFunc: The database query function (e.g., `GetActivityByID`).
// - cacheKey: The Redis key under which the data will be cached.
// - expiration: Time for cache expiration.
func cacheFetch[T any](ctx context.Context, redisClient *redis.Client, cacheKey string, expiration time.Duration, fetchFunc func() (T, error)) (T, error) {
	var result T

	// Try to get from cache
	cached, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil { // Cache hit
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return result, nil
		}
	}

	// Cache miss, call the database query function
	result, err = fetchFunc()
	if err != nil {
		return result, err
	}

	// Cache the result
	data, err := json.Marshal(result)
	if err != nil {
		return result, err
	}

	// Set the cache with an expiration time
	err = redisClient.Set(ctx, cacheKey, data, expiration).Err()
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetActivityByID uses the generic cacheFetch to handle caching logic.
func (c *Cache) GetActivityByID(ctx context.Context, id int64) (ActiveActivityDetail, error) {
	cacheKey := getActivityCacheKey(id)
	expiration := 10 * time.Minute

	return cacheFetch(ctx, c.redisClient, cacheKey, expiration, func() (ActiveActivityDetail, error) {
		return c.dbQueries.FindActiveActivityDetails(ctx, id)
	})
}

// InvalidateActivityCache invalidates the cache for a given activity ID (e.g., after an update).
func (c *Cache) InvalidateActivityCache(ctx context.Context, id int64) error {
	cacheKey := getActivityCacheKey(id)
	return c.redisClient.Del(ctx, cacheKey).Err()
}

// getActivityCacheKey generates a Redis key for activity cache.
func getActivityCacheKey(id int64) string {
	return "activity:" + string(rune(id))
}
