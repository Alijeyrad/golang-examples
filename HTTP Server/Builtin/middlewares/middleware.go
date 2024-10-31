package middleware

import (
	"net/http"
)

// Handler is our custom handler type that includes the next handler in the chain.
type Handler func(w http.ResponseWriter, r *http.Request)

// Middleware is a function that takes a Handler and returns a Handler.
type Middleware func(Handler) Handler

// Chain builds the middleware chain recursively.
func Chain(f Handler, m ...Middleware) Handler {
	// If there is no middleware left, return the original handler
	if len(m) == 0 {
		return f
	}
	// Wrap the next middleware
	return m[0](Chain(f, m[1:]...))
}
