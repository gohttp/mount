package main

import "github.com/gohttp/logger"
import "github.com/gohttp/serve"
import "github.com/gohttp/mount"
import "github.com/gohttp/app"
import "net/http"

func main() {
	a := app.New()

	a.Use(logger.New())
	a.Use(mount.New("/examples", serve.New("examples")))
	a.Use(mount.New("/blog", blog()))
	a.Use(mount.New("/hello", hello))
	a.Get("/", http.HandlerFunc(hello))

	a.Listen(":3000")
}

func blog() *app.App {
	app := app.New()
	app.Get("", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("posts\n"))
	})
	return app
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello\n"))
}
