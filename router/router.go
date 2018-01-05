package router

import (
	"github.com/labstack/echo"
)

// Handler interface
type Handler interface {
	Register(*echo.Echo)
}

// Route represent an api route
type Route struct {
	Method      string
	Path        string
	Handler     func(echo.Context) error
	Middlewares []echo.MiddlewareFunc
}

// Register a route
func (r Route) Register(server *echo.Echo) {
	server.Add(r.Method, r.Path, r.Handler, r.Middlewares...)
}

// Routes is an array of route
type Routes []Route

// Register all routes
func (r Routes) Register(server *echo.Echo) {
	for _, route := range r {
		route.Register(server)
	}
}

// Router represent the application router
type Router struct {
	Server *echo.Echo
}

// NewRouter create a new instance of router
func NewRouter() *Router {
	return &Router{
		Server: echo.New(),
	}
}

// Register a route
func (r *Router) Register(handler Handler) {
	handler.Register(r.Server)
}
