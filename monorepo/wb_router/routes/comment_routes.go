package routes

import (
	"net/http"
	"v1/v1/handlers"
)

func GetCommentRoutes(h handlers.CommentHandler) []Route {
	return []Route{
		{
			URI:         "/post/{postId}/comments",
			Method:      http.MethodPost,
			Function:    h.CreateComment,
			RequireAuth: true,
		},
		{
			URI:         "/post/{postId}/comments",
			Method:      http.MethodGet,
			Function:    h.GetComments,
			RequireAuth: true,
		},
	}
}
