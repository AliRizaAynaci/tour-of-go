package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// DONE: Burada ki mapi thread safe hale getirebilirsiniz.
// 102-concurrency egitimindeki mutex orneklerine bakabilirsiniz.
// Ref: https://pmihaylov.com/thread-safety-concerns-go/
// Ref: https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c

// Thread-safe coutner structure using LimitCounter struct
var counter = &LimitCounter{v: map[string]*Limit{}}

type LimitCounter struct {
	sync.Mutex
	v map[string]*Limit
}

type Limit struct {
	count int
	ttl   time.Time
}

type LimitProxy struct {
	key   string
	limit int
	ttl   time.Duration
}

// ResetLimitHandler resets the limit for a given key
func ResetLimitHandler(c *fiber.Ctx) error {
	// Done: [DELETE] /limit/:key/* pathine istek atildiginda limiti sifirlayan handleri implement edebilirsiniz.
	// Done: implement me!

	key := strings.TrimPrefix(c.Path(), "/limit")

	// If the key does not exist, return NotFound error
	counter.Lock()
	defer counter.Unlock()
	if _, ok := counter.v[key]; ok {
		return fiber.ErrNotFound
	}

	delete(counter.v, key) // Safely deletes the key
	c.Response().SetStatusCode(fiber.StatusNoContent)
	return nil
}

func NewLimitProxy(key string, limit int, ttl time.Duration) LimitProxy {
	return LimitProxy{
		key:   key,
		limit: limit,
		ttl:   ttl,
	}
}

func (p LimitProxy) Accept(key string) bool {
	return p.key == key
}

func (p LimitProxy) Proxy(c *fiber.Ctx) error {
	path := c.Path()

	counter.Lock()
	defer counter.Unlock()

	// If the path was added before, check the limit and TTL
	if r, ok := counter.v[path]; ok {
		// If the limit is reached and TTL has not expired, return an error
		if r.count >= p.limit && r.ttl.After(time.Now()) {
			c.Response().SetStatusCode(fiber.StatusTooManyRequests)

			fmt.Printf("request limit reached for %s\n", path)
			return fiber.ErrTooManyRequests
		}

		// If the TTL has expired, reset the counter and renew the TTL
		if r.ttl.Before(time.Now()) {
			r.count = 0
			r.ttl = time.Now().Add(p.ttl)
		}
	} else {
		// For a new path, create a `Limit` instance
		counter.v[path] = &Limit{
			count: 0,
			ttl:   time.Now().Add(p.ttl),
		}
	}

	// Increment the counter when the request is accepted
	counter.v[path].count++

	if err := c.SendString("Go Turkiye - 103 Http Package"); err != nil {
		return err
	}

	return nil
}
