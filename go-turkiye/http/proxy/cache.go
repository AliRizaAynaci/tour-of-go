package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"strings"
	"sync"
	"time"
)

// DONE: Burada ki mapi thread safe hale getirebilirsiniz.
// 102-concurrency egitimindeki mutex orneklerine bakabilirsiniz.
// Ref: https://pmihaylov.com/thread-safety-concerns-go/
// Ref: https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c

// Thread-safe cache structure using CacheStore with mutex for concurrency control.
var cache = &CacheStore{data: map[string]Cache{}}

// CacheStore struct holds the cache data and includes a sync.Mutex for thread-safe operations.
type CacheStore struct {
	sync.Mutex
	data map[string]Cache
}

type Cache struct {
	body []byte
	ttl  time.Time
}

type CacheProxy struct {
	key string
	ttl time.Duration
}

// EvictCacheHandler deletes a cache entry for a specific path.
func EvictCacheHandler(c *fiber.Ctx) error {
	// DONE: [DELETE] /cache/:key/* pathine istek atildiginda memorydeki cachei temizleyen handleri implement edebilirsiniz.
	// Done: implement me!

	key := c.Params("key")
	path := "/" + key + c.Params("*") // Combine key and remaining path to form the full cache path.

	cache.Lock()
	defer cache.Unlock()

	// If the key doesn't exist in the cache, return NotFound error
	if _, ok := cache.data[path]; ok {
		return fiber.ErrNotFound
	}

	// Safely delete the cache entry
	delete(cache.data, path)

	c.Response().SetStatusCode(fiber.StatusNoContent)
	return nil
}

func NewCacheProxy(key string, ttl time.Duration) CacheProxy {
	return CacheProxy{
		key: key,
		ttl: ttl,
	}
}

func (p CacheProxy) Accept(key string) bool {
	return p.key == key
}

// Proxy handles caching and redirects requests to the original URL if needed.
func (p CacheProxy) Proxy(c *fiber.Ctx) error {
	path := c.Path()
	key := c.Params("key")

	cache.Lock()
	defer cache.Unlock()

	// Cache if cache exists and is still valid
	if r, ok := cache.data[path]; ok && r.ttl.After(time.Now()) {
		c.Response().SetBody(r.body)
		c.Response().Header.Add("cache-control", fmt.Sprintf("max-age:%d", p.ttl/time.Second))
		return nil
	}

	// Construct the URL by triming the key prefix from the path
	url := "https://mocki.io/" + strings.TrimPrefix(path, "/"+key+"/")
	fmt.Printf("HTTP request redirecting to %s \n", url)

	// Forward request to original URL using Fiber's proxy middleware
	if err := proxy.Do(c, url); err != nil {
		return err
	}

	respStatusCode := c.Response().StatusCode()

	if respStatusCode != fiber.StatusOK {
		return fiber.NewError(respStatusCode, "Check your request")
	}

	// Update cache with the new response
	ch := Cache{
		ttl:  time.Now().Add(p.ttl),
		body: c.Response().Body(),
	}
	cache.data[path] = ch

	// Remove the Server header for consistency
	c.Response().Header.Del(fiber.HeaderServer)
	return nil
}
