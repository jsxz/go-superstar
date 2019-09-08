package main

import (
	"github.com/jsxz/go-superstar/bootstrap"
	"github.com/jsxz/go-superstar/routes"
	"github.com/jsxz/go-superstar/web/middleware/identity"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Superstar database", "anjun")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8888")
}
