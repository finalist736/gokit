package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/finalist736/gokit/config"
	"github.com/finalist736/gokit/logger"
	"github.com/finalist736/gokit/mainloop"
	"github.com/finalist736/gokit/webserver"
	"github.com/gocraft/web"
)

var configPath = flag.String("config", "config.ini", "where config path is")

func main() {
	flag.Parse()

	var err = config.Init(config.NewFileProvider(configPath))
	if err != nil {
		panic(err)
	}

	logger.ReloadLogs(config.DefaultConfig())

	logger.StdOut().Debugln(config.Log())

	webserver.Start(createRouters(), config.MustString("listen"))

	mainloop.Loop(stop, grace, config.DefaultConfig())
}

type Ctx struct {
	User int
	// *User
	// *db
	// *cache...
}

func createRouters() *web.Router {
	router := web.New(Ctx{})

	router.Middleware(func(c *Ctx, rw web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {
		c.User = rand.Intn(100) + 1
		next(rw, r)
	})

	router.Get("/", func(rw web.ResponseWriter, r *web.Request) {
		fmt.Fprint(rw, "<h1>Main Page</h1><a href='/test'>test</a>")
	})
	router.Get("/test", func(c *Ctx, rw web.ResponseWriter, r *web.Request) {
		logger.StdOut().Debugf("user id: %d", c.User)
		fmt.Fprintf(rw, "<h1>User Page</h1><h3>Hello user with id: %d</h3><a href='/'>main</a>", c.User)
	})
	return router
}

func stop() {
	webserver.Stop()
}

func grace() {
	webserver.Grace(configPath)
}
