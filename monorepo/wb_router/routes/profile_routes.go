package routes

import (
	"net/http"
	"v1/v1/handlers"
)

func GetProfileRoutes(h handlers.ProfileHandler) []Route {
	return []Route{
		{URI: "/users/{userId}/profile", Method: http.MethodGet, Function: h.GetProfile, RequireAuth: true},

		{URI: "/profile/education", Method: http.MethodPost, Function: h.AddEducation, RequireAuth: true},
		{URI: "/profile/education/{id}", Method: http.MethodPut, Function: h.UpdateEducation, RequireAuth: true},
		{URI: "/profile/education/{id}", Method: http.MethodDelete, Function: h.DeleteEducation, RequireAuth: true},

		{URI: "/profile/projects", Method: http.MethodPost, Function: h.AddProject, RequireAuth: true},
		{URI: "/profile/projects/{id}", Method: http.MethodPut, Function: h.UpdateProject, RequireAuth: true},
		{URI: "/profile/projects/{id}", Method: http.MethodDelete, Function: h.DeleteProject, RequireAuth: true},

		{URI: "/profile/courses", Method: http.MethodPost, Function: h.AddCourse, RequireAuth: true},
		{URI: "/profile/courses/{id}", Method: http.MethodPut, Function: h.UpdateCourse, RequireAuth: true},
		{URI: "/profile/courses/{id}", Method: http.MethodDelete, Function: h.DeleteCourse, RequireAuth: true},
	}
}
