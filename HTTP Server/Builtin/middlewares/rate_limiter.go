package middleware

import (
	"net/http"
	"sync"
	"time"
)

func RateLimiter(limit int, window time.Duration) Middleware {
	var clients = make(map[string]*client)
	var mu sync.Mutex

	type client struct {
		requests int
		lastSeen time.Time
	}

	return func(next Handler) Handler {
		return func(w http.ResponseWriter, r *http.Request) {
			mu.Lock()
			c, exists := clients[r.RemoteAddr]
			if !exists {
				c = &client{}
				clients[r.RemoteAddr] = c
			}
			c.requests++
			c.lastSeen = time.Now()

			if c.requests > limit {
				mu.Unlock()
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
			mu.Unlock()

			next(w, r)

			// Clean up old entries
			go func() {
				mu.Lock()
				defer mu.Unlock()
				for addr, client := range clients {
					if time.Since(client.lastSeen) > window {
						delete(clients, addr)
					}
				}
			}()
		}
	}
}
