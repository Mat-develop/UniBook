package routes

import (
	"net/http"
	"v1/v1/handlers"
)

const (
	Communities    = "/c"
	AllCommunities = "/c/all"
)

func GetCommunitiesRoutes(c handlers.CommunityHandler) []Route {
	return []Route{
		{
			URI:         Communities,
			Method:      http.MethodPost,
			Function:    c.GetCommunityByName,
			RequireAuth: false,
		},
		{
			URI:         AllCommunities,
			Method:      http.MethodGet,
			Function:    c.ListCommunities,
			RequireAuth: true,
		},
	}

}
