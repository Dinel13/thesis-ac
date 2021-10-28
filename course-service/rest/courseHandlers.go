package rest

import (
	"net/http"
	"strconv"

	"github.com/dinel13/thesis-ac/course/service"
	"github.com/julienschmidt/httprouter"
)

type CourseHandlers interface {
	GetCourse(http.ResponseWriter, *http.Request)
}

type DefaultCourseHandlers struct {
	Service service.CourseService
}

// GetCourse is handler for GET /course/{id}
func (h *DefaultCourseHandlers) GetCourse(w http.ResponseWriter, r *http.Request) {
	// get the course id from request param
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
	}

	// get the course from service
	course, err := h.Service.GetCourse(id)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the course to response
	WriteJson(w, http.StatusOK, course, "course")
}

func NewCoursRestHandlers(s service.CourseService) DefaultCourseHandlers {
	return DefaultCourseHandlers{s}
}
