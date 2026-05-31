package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"v1/tag/service"
	"v1/util/authentication"
	"v1/util/response"
)

type TagHandler interface {
	CreateTag(w http.ResponseWriter, r *http.Request)
	ListTags(w http.ResponseWriter, r *http.Request)
}

type tagHandler struct {
	service service.TagService
}

func NewTagHandler(service service.TagService) TagHandler {
	return &tagHandler{service: service}
}

func (h *tagHandler) CreateTag(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var payload struct {
		Name string `json:"name"`
	}
	if err = json.Unmarshal(body, &payload); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	id, err := h.service.Create(payload.Name, userID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, id)
}

func (h *tagHandler) ListTags(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	tags, err := h.service.Search(name)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, tags)
}
