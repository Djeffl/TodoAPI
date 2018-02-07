package api

import (
	"TodoAPI/routing/models"
)

var apiRoutes = []models.Routes{
	todoRoutes,
}

//Collect all api's
func GetAPIRoutes() models.Routes {
	var rts models.Routes = make([]models.Route, 0, 5)

	for _, routes := range apiRoutes {
		for _, route := range routes {
			rts = append(rts, route)
		}
	}
	return rts
}
