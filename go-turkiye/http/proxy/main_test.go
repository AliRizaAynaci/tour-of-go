package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
	"time"
)

// initializeApp initializes the fiber app and returns it for testing purposes
func initializeApp() *fiber.App {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/:key/*", ProxyHandler)
	app.Delete("cache/:key/*", EvictCacheHandler)
	app.Delete("limit/:key/*", ResetLimitHandler)
	return app
}

// resetGlobals clears cache and counter before each test to ensure isolation.
func resetGlobals() {
	cache = &CacheStore{data: map[string]Cache{}}
	counter = &LimitCounter{v: map[string]*Limit{}}
}

// TestRootRoute tests the root route "/"
func TestRootRoute(t *testing.T) {
	resetGlobals()
	app := initializeApp()
	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req)

	// Check if response code is 200
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Check if response body matches expected output
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "Hello, World!", string(body))
}

// TestProxyHandler tests the ProxyHandler functionality for caching and limiting
func TestProxyHandler(t *testing.T) {
	resetGlobals() // Ensure counter and cache are reset
	app := initializeApp()

	// First request, should be successful (200 OK)
	req1 := httptest.NewRequest("GET", "/user/testpath", nil)
	resp1, _ := app.Test(req1)

	assert.Equal(t, fiber.StatusOK, resp1.StatusCode)

	// Log the state after the first request
	t.Logf("First request - Path /user/testpath - Counter: %d", counter.v["/user/testpath"].count)

	// Send requests within the limit
	for i := 0; i < 2; i++ {
		req := httptest.NewRequest("GET", "/user/testpath", nil)
		resp, _ := app.Test(req)

		// Log each successful request
		t.Logf("Request #%d - Path /user/testpath - Counter: %d", i+2, counter.v["/user/testpath"].count)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	}

	// Exceed the limit with one more request
	req := httptest.NewRequest("GET", "/user/testpath", nil)
	resp, _ := app.Test(req)

	// Log the counter state before the final assertion
	t.Logf("Final request - Path /user/testpath - Expected to fail with 429 - Counter: %d",
		counter.v["/user/testpath"].count)

	// Expect 429 Too Many Requests as the limit is reached
	assert.Equal(t, fiber.StatusTooManyRequests, resp.StatusCode)
}

// TestEvictCacheHandler tests the cache eviction route
func TestEvictCacheHandler(t *testing.T) {
	resetGlobals()
	app := initializeApp()

	// Pre-fill cache with an entry
	cache.Lock()
	cache.data["/user/testpath"] = Cache{body: []byte("Cached response"), ttl: time.Now().Add(1 * time.Minute)}
	cache.Unlock()

	// Confirm cache entry exists
	_, exists := cache.data["/user/testpath"]
	assert.True(t, exists)

	// Send DELETE request to evict cache
	req := httptest.NewRequest("DELETE", "/cache/user/testpath", nil)
	resp, _ := app.Test(req)

	// Confirm deletion was successful
	assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)

	// Verify cache entry is gone
	_, exists = cache.data["/user/testpath"]
	assert.False(t, exists)
}

// TestResetLimitHandler tests the rate limit reset route
func TestResetLimitHandler(t *testing.T) {
	resetGlobals()
	app := initializeApp()

	// Manually set a limit count for testing reset
	counter.Lock()
	counter.v["/user/testpath"] = &Limit{count: 3, ttl: time.Now().Add(1 * time.Minute)}
	counter.Unlock()

	// Confirm the limit is set
	assert.Equal(t, 3, counter.v["/user/testpath"].count)

	// Send DELETE request to reset limit
	req := httptest.NewRequest("DELETE", "/limit/user/testpath", nil)
	resp, _ := app.Test(req)

	// Confirm reset was successful
	assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)

	// Verify limit is removed from counter
	_, exists := counter.v["/user/testpath"]
	assert.False(t, exists)
}
