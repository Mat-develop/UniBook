package routes

import (
	"net/http"
	"v1/v1/handlers"
)

const (
	Tags = "/tags"
)

func GetTagRoutes(h handlers.TagHandler) []Route {
	return []Route{
		{
			URI:         Tags,
			Method:      http.MethodPost,
			Function:    h.CreateTag,
			RequireAuth: true,
		},
		{
			URI:         Tags,
			Method:      http.MethodGet,
			Function:    h.ListTags,
			RequireAuth: true,
		},
	}
}
