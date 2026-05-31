package routes

import (
	"net/http"
	"v1/v1/handlers"
)

func GetSearchRoutes(h handlers.SearchHandler) []Route {
	return []Route{
		{
			URI:         "/search",
			Method:      http.MethodGet,
			Function:    h.Search,
			RequireAuth: true,
		},
	}
}
