// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/fd1az/go_service/foundation/web"
)

func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App {

	app := web.NewApp()

	check := check{
		log: log,
	}
	app.Handle(http.MethodGet, "/readiness", check.readiness)

	return app
}
