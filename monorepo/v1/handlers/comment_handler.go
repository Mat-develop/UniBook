package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"v1/comment/model"
	"v1/comment/service"
	"v1/util/authentication"
	"v1/util/response"

	"github.com/gorilla/mux"
)

type CommentHandler interface {
	CreateComment(w http.ResponseWriter, r *http.Request)
	GetComments(w http.ResponseWriter, r *http.Request)
}

type commentHandler struct {
	service service.CommentService
}

func NewCommentHandler(svc service.CommentService) CommentHandler {
	return &commentHandler{service: svc}
}

func (h *commentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var dto model.CommentDTO
	if err = json.Unmarshal(body, &dto); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = h.service.CreateComment(postID, userID, dto); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, nil)
}

func (h *commentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	comments, err := h.service.GetComments(postID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, comments)
}
