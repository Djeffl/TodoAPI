package routing

import (
	"TodoAPI/routing/api"
	"TodoAPI/routing/helper"
	"TodoAPI/routing/models"
	"net/http"

	"github.com/gorilla/mux"
)

var apiRoutes models.Routes = api.GetAPIRoutes()

var routes models.Routes = apiRoutes

func InvokeRouter() {
	// Collect routes in variable
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = helper.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	http.Handle("/", router)
}
