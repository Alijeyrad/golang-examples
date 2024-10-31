package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var clients = make(map[string]*client)
var mu sync.Mutex

type client struct {
	requests int
	lastSeen time.Time
}

func RateLimiter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		mu.Lock()
		defer mu.Unlock()
		addr := c.IP()
		cl, exists := clients[addr]
		if !exists {
			cl = &client{}
			clients[addr] = cl
		}
		cl.requests++
		cl.lastSeen = time.Now()

		if cl.requests > 10 {
			return c.Status(fiber.StatusTooManyRequests).SendString("Too Many Requests")
		}

		// Clean up old entries
		go func() {
			mu.Lock()
			defer mu.Unlock()
			for addr, client := range clients {
				if time.Since(client.lastSeen) > time.Minute {
					delete(clients, addr)
				}
			}
		}()

		return c.Next()
	}
}
