package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"v1/profile/model"
	"v1/profile/service"
	"v1/util/authentication"
	"v1/util/response"

	"github.com/gorilla/mux"
)

type ProfileHandler interface {
	GetProfile(w http.ResponseWriter, r *http.Request)
	AddEducation(w http.ResponseWriter, r *http.Request)
	UpdateEducation(w http.ResponseWriter, r *http.Request)
	DeleteEducation(w http.ResponseWriter, r *http.Request)
	AddProject(w http.ResponseWriter, r *http.Request)
	UpdateProject(w http.ResponseWriter, r *http.Request)
	DeleteProject(w http.ResponseWriter, r *http.Request)
	AddCourse(w http.ResponseWriter, r *http.Request)
	UpdateCourse(w http.ResponseWriter, r *http.Request)
	DeleteCourse(w http.ResponseWriter, r *http.Request)
}

type profileHandler struct {
	service service.ProfileService
}

func NewProfileHandler(s service.ProfileService) ProfileHandler {
	return &profileHandler{service: s}
}

func (h *profileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(mux.Vars(r)["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	profile, err := h.service.GetProfile(userID)
	if err != nil {
		response.Erro(w, http.StatusNotFound, err)
		return
	}
	response.JSON(w, http.StatusOK, profile)
}

func decodeBody[T any](r *http.Request) (T, error) {
	var v T
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return v, err
	}
	err = json.Unmarshal(body, &v)
	return v, err
}

func (h *profileHandler) AddEducation(w http.ResponseWriter, r *http.Request) {
	callerID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}
	e, err := decodeBody[model.Education](r)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.AddEducation(callerID, e)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusCreated, map[string]uint64{"id": id})
}

func (h *profileHandler) UpdateEducation(w http.ResponseWriter, r *http.Request) {
	callerID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	e, err := decodeBody[model.Education](r)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = h.service.UpdateEducation(callerID, id, e); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, nil)
}

func (h *profileHandler) DeleteEducation(w http.ResponseWriter, r *http.Request) {
	callerID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = h.service.DeleteEducation(callerID, id); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, nil)
}

func (h *profileHandler) AddProject(w http.ResponseWriter, r *http.Request) {
	callerID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}
	p, err := decodeBody[model.Project](r)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.AddProject(callerID, p)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusCreated, map[string]uint64{"id": id})
}

func (h *profileHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	callerID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	p, err := decodeBody[model.Project](r)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = h.service.UpdateProject(callerID, id, p); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, nil)
}

func (h *profileHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	callerID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = h.service.DeleteProject(callerID, id); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, nil)
}

func (h *profileHandler) AddCourse(w http.ResponseWriter, r *http.Request) {
	callerID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}
	c, err := decodeBody[model.Course](r)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.AddCourse(callerID, c)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusCreated, map[string]uint64{"id": id})
}

func (h *profileHandler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	callerID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	c, err := decodeBody[model.Course](r)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = h.service.UpdateCourse(callerID, id, c); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, nil)
}

func (h *profileHandler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	callerID, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = h.service.DeleteCourse(callerID, id); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, nil)
}
