package getuser

import (
	"fmt"
	"github.com/J0eppp/api.j0eppp.dev/pkg/application"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Do(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "hello")
	}
}
