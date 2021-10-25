package rest

import (
	"net/http"
	"strconv"

	"github.com/dinel13/thesis-ac/course/service"
	"github.com/julienschmidt/httprouter"
)

type CourseHandlers struct {
	service service.CourseService
}

// GetCourse is handler for GET /course/{id}
func (h *CourseHandlers) GetCourse(w http.ResponseWriter, r *http.Request) {
	// get the course id from request param
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
	}

	// get the course from service
	course, err := h.service.GetCourse(id)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the course to response
	WriteJson(w, http.StatusOK, course, "course")
}
