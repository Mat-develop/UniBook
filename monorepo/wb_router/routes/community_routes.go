package routes

import (
	"net/http"
	"v1/v1/handlers"
)

const (
	Communities             = "/c"
	AllCommunities          = "/c/all"
	CommunityByID           = "/c/{communityId}"
	CommunityFollowByID     = "/c/{communityId}/follow"
	CommunityUnfollowByID   = "/c/{communityId}/unfollow"
	CommunityFollowersByID  = "/c/{communityId}/followers"
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
