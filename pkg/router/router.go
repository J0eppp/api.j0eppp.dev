package router

import (
	"github.com/J0eppp/api.j0eppp.dev/cmd/api/handlers/getuser"
	"github.com/J0eppp/api.j0eppp.dev/pkg/application"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/users", getuser.Do(app))
	return mux
}
