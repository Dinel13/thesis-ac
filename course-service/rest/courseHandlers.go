package rest

import (
	"net/http"
	"strconv"

	"github.com/dinel13/thesis-ac/course/domain"
	"github.com/julienschmidt/httprouter"
)

func NewCoursRestHandlers(s domain.CourseService) domain.CourseRestHandlers {
	return courseHandlers{s}
}

type courseHandlers struct {
	service domain.CourseService
}

// GetCourse is handler for GET /course/{id}
func (h courseHandlers) Read(w http.ResponseWriter, r *http.Request) {
	// get the course id from request param
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
	}

	// get the course from service
	course, err := h.service.Read(id)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the course to response
	WriteJson(w, http.StatusOK, course, "course")
}

// Create is handler for POST /course to create COurse
func (h courseHandlers) Create(w http.ResponseWriter, r *http.Request) {
	// get the course from request body
	course := &domain.Course{}
	err := ReadJson(r, course)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// create the course
	c, err := h.service.Create(course)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the course to response
	WriteJson(w, http.StatusCreated, c, "course")
}

// Update is handler for PUT /course to update Course
func (h courseHandlers) Update(w http.ResponseWriter, r *http.Request) {
	// get the course from request body
	course := &domain.Course{}
	err := ReadJson(r, course)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// update the course
	c, err := h.service.Update(course)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the course to response
	WriteJson(w, http.StatusOK, c, "course")
}

// Delete is handler for DELETE /course/{id} to delete Course
func (h courseHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	// get the course id from request param
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// delete the course
	err = h.service.Delete(id)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the course to response
	WriteJson(w, http.StatusOK, nil, "course")
}
