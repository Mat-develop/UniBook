package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"v1/community/model"
	"v1/community/service"
	"v1/util/authentication"
	"v1/util/response"

	"github.com/gorilla/mux"
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
func (c *communityHandler) FollowCommunity(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	communityID, err := strconv.ParseUint(params["communityId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if strings.Contains(r.URL.Path, "/follow") {
		err = c.service.Follow(userID, communityID)
	} else {
		err = c.service.Unfollow(userID, communityID)
	}

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func (c *communityHandler) GetCommunityFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	communityID, err := strconv.ParseUint(params["communityId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	followers, err := c.service.GetFollowers(communityID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}
