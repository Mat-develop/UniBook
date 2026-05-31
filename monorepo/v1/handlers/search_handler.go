package handlers

import (
	"net/http"
	"strings"

	communityService "v1/community/service"
	postService "v1/post/service"
	"v1/util/authentication"
	"v1/util/response"
)

type SearchHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
}

type searchHandler struct {
	communities communityService.CommunityService
	posts       postService.PostService
}

func NewSearchHandler(cs communityService.CommunityService, ps postService.PostService) SearchHandler {
	return &searchHandler{communities: cs, posts: ps}
}

func (h *searchHandler) Search(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.URL.Query().Get("q"))

	viewerID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	communities, err := h.communities.Search(q)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	posts, err := h.posts.SearchPosts(viewerID, q)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]interface{}{
		"communities": communities,
		"posts":       posts,
	})
}
