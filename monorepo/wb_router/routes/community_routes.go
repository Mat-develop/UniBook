package routes

import (
	"net/http"
	"v1/v1/handlers"
)

const (
	Communities              = "/c"
	CreateCommunity          = "/c/create"
	AllCommunities           = "/c/all"
	JoinedCommunities        = "/c/joined"
	CommunityByID            = "/c/{communityId}"
	CommunityFollowByID      = "/c/{communityId}/follow"
	CommunityUnfollowByID    = "/c/{communityId}/unfollow"
	CommunityFollowersByID   = "/c/{communityId}/followers"
	CommunitiesByCreator     = "/users/{userId}/communities"
)

func GetCommunitiesRoutes(c handlers.CommunityHandler) []Route {
	return []Route{
		{
			URI:         CreateCommunity,
			Method:      http.MethodPost,
			Function:    c.CreateCommunity,
			RequireAuth: true,
		},
		{
			URI:         CommunitiesByCreator,
			Method:      http.MethodGet,
			Function:    c.GetCreatedCommunities,
			RequireAuth: true,
		},
		{
			URI:         AllCommunities,
			Method:      http.MethodGet,
			Function:    c.ListCommunities,
			RequireAuth: true,
		},
		{
			URI:         JoinedCommunities,
			Method:      http.MethodGet,
			Function:    c.GetJoinedCommunities,
			RequireAuth: true,
		},
		{
			URI:         CommunityByID,
			Method:      http.MethodGet,
			Function:    c.GetCommunityByID,
			RequireAuth: true,
		},
		{
			URI:         CommunityFollowByID,
			Method:      http.MethodPost,
			Function:    c.FollowCommunity,
			RequireAuth: true,
		},
		{
			URI:         CommunityUnfollowByID,
			Method:      http.MethodPost,
			Function:    c.FollowCommunity,
			RequireAuth: true,
		},
		{
			URI:         CommunityFollowersByID,
			Method:      http.MethodGet,
			Function:    c.GetCommunityFollowers,
			RequireAuth: true,
		},
	}

}
