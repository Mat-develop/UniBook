package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"v1/post/model"
	"v1/post/service"
	"v1/util/authentication"
	"v1/util/response"

	"github.com/gorilla/mux"
)

type PostHandler interface {
	GetCommunityPosts(w http.ResponseWriter, r *http.Request)
	GetUserPosts(w http.ResponseWriter, r *http.Request)
	GetPostByTittle(w http.ResponseWriter, r *http.Request)
	GetFeed(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
}

type postHandler struct {
	service service.PostService
}

func NewPostHandler(handlerService service.PostService) PostHandler {
	return &postHandler{service: handlerService}
}

func (p *postHandler) GetCommunityPosts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	communityId, err := strconv.ParseUint(params["communityId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
	}

	posts, err := p.service.GetPosts(0, communityId)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
	}

	response.JSON(w, http.StatusOK, posts)
}

func (p *postHandler) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	posts, err := p.service.GetPosts(userId, 0)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, posts)
}

func (p *postHandler) GetPostByTittle(w http.ResponseWriter, r *http.Request) {

}

func (p *postHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var postBody model.PostDTO
	if err = json.Unmarshal(body, &postBody); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	err = p.service.CreatePost(userId, postBody)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, nil)
}

func (p *postHandler) GetFeed(w http.ResponseWriter, r *http.Request) {

}

func (p *postHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {

}

func (p *postHandler) DeletePost(w http.ResponseWriter, r *http.Request) {

}
