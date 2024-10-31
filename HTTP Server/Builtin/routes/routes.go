package router

import (
	"library/controllers"
	"library/middleware"
	"library/repositories"
	"library/services"
	"net/http"
	"strings"
	"time"
)

type Router struct {
	routes map[string]map[string]http.HandlerFunc // method -> path -> handler
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	// Try to find exact match
	if handlers, ok := r.routes[method]; ok {
		if handler, ok := handlers[path]; ok {
			handler(w, req)
			return
		}
	}

	// Handle routes with parameters (e.g., /books/{id})
	for routePath, handler := range r.routes[method] {
		if strings.Contains(routePath, "{id}") {
			basePath := strings.Split(routePath, "{id}")[0]
			if strings.HasPrefix(path, basePath) {
				handler(w, req)
				return
			}
		}
	}

	http.NotFound(w, req)
}

// AddRoute adds a route with middlewares
func (r *Router) AddRoute(method, path string, handler http.HandlerFunc, middlewares ...middleware.Middleware) {
	if _, exists := r.routes[method]; !exists {
		r.routes[method] = make(map[string]http.HandlerFunc)
	}
	finalHandler := middleware.Chain(handler, middlewares...)
	r.routes[method][path] = finalHandler
}

func NewBookController() *controllers.BookController {
	bookRepo := repositories.NewBookRepository()
	bookService := services.NewBookService(bookRepo)
	bookController := controllers.NewBookController(bookService)
	return bookController
}

func NewRouterWithRoutes() *Router {
	r := NewRouter()
	bookController := NewBookController()

	// Middlewares
	rateLimiter := middleware.RateLimiter(10, time.Minute)
	logger := middleware.Logger()

	// Routes with middlewares
	r.AddRoute("GET", "/books", bookController.GetAllBooks, logger)
	r.AddRoute("GET", "/books/{id}", bookController.GetBookByID, logger)
	r.AddRoute("POST", "/books", bookController.CreateBook, logger, rateLimiter)
	r.AddRoute("PUT", "/books/{id}", bookController.UpdateBook, logger, rateLimiter)
	r.AddRoute("DELETE", "/books/{id}", bookController.DeleteBook, logger, rateLimiter)

	return r
}
