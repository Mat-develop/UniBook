package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"v1/community/model"
	"v1/community/service"
	"v1/util/response"
)

type CommunityHandler interface {
	CreateCommunity(w http.ResponseWriter, r *http.Request)
	GetCommunityByID(w http.ResponseWriter, r *http.Request)
	GetCommunityByName(w http.ResponseWriter, r *http.Request)
	ListCommunities(w http.ResponseWriter, r *http.Request)
	DeleteCommunity(w http.ResponseWriter, r *http.Request)
	FollowCommunity(w http.ResponseWriter, r *http.Request)
	GetCommunityFollowers(w http.ResponseWriter, r *http.Request)
}

type communityHandler struct {
	service service.CommunityService
}

func NewCommunityHandler(service service.CommunityService) CommunityHandler {
	return &communityHandler{service: service}
}

func (c *communityHandler) CreateCommunity(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var community model.Community
	if err = json.Unmarshal(body, &community); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	id, err := c.service.Create(community)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusAccepted, id)

}

func (c *communityHandler) GetCommunityByID(w http.ResponseWriter, r *http.Request) {}

func (c *communityHandler) GetCommunityByName(w http.ResponseWriter, r *http.Request) {}
func (c *communityHandler) ListCommunities(w http.ResponseWriter, r *http.Request) {
	communities, err := c.service.ListCommunities()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, communities)
}

func (c *communityHandler) DeleteCommunity(w http.ResponseWriter, r *http.Request)       {}
func (c *communityHandler) FollowCommunity(w http.ResponseWriter, r *http.Request)       {}
func (c *communityHandler) GetCommunityFollowers(w http.ResponseWriter, r *http.Request) {}
