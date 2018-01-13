package router

import (
	"github.com/labstack/echo"
)

// Handler register
type Handler interface {
	Register(*echo.Group)
}

// Handlers is an array of handler
type Handlers []Handler

// Register an handler to the group
func (h Handlers) Register(g *echo.Group) {
	for _, handler := range h {
		handler.Register(g)
	}
}

// Route structure
type Route struct {
	Method      string
	Path        string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}

// Register a route
func (s *Route) Register(g *echo.Group) {
	g.Add(s.Method, s.Path, s.Handler, s.Middlewares...)
}

// CRUD structure
type CRUD struct {
	Path        string
	Middlewares []echo.MiddlewareFunc
	Create      Route
	Read        Route
	Update      Route
	Delete      Route
}

// Register CRUD
func (s CRUD) Register(g *echo.Group) {
	g.Use(s.Middlewares...)

	g.POST(s.Path+s.Create.Path, s.Create.Handler, s.Create.Middlewares...)
	g.GET(s.Path+s.Read.Path, s.Read.Handler, s.Read.Middlewares...)
	g.PUT(s.Path+s.Update.Path, s.Update.Handler, s.Update.Middlewares...)
	g.DELETE(s.Path+s.Delete.Path, s.Delete.Handler, s.Delete.Middlewares...)
}
