package routes

import (
	"net/http"
	"v1/v1/handlers"
)

const (
	Post          = "/post"
	PostById      = "/post/{id}"
	PostByUserId  = "/post/{userId}"
	CommunityPost = "/post/c/{communityId}"
	PostByName    = "/post/{PostId}"
	PostByTitle   = "/post/{title}"
)

// dps faço o restante
func GetPostRoutes(h handlers.PostHandler) []Route {
	return []Route{
		{
			URI:         Post,
			Method:      http.MethodPost,
			Function:    h.CreatePost,
			RequireAuth: false,
		},
		{
			URI:         CommunityPost,
			Method:      http.MethodGet,
			Function:    h.GetCommunityPosts,
			RequireAuth: true,
		},
		{
			URI:         PostByUserId,
			Method:      http.MethodGet,
			Function:    h.GetUserPosts,
			RequireAuth: true,
		},
		{
			URI:         PostByTitle,
			Method:      http.MethodGet,
			Function:    h.GetPostByTittle,
			RequireAuth: true,
		},
		{
			URI:         PostById,
			Method:      http.MethodPut,
			Function:    h.UpdatePost,
			RequireAuth: true,
		},
		{
			URI:         PostById,
			Method:      http.MethodDelete,
			Function:    h.DeletePost,
			RequireAuth: true,
		},
	}

}
