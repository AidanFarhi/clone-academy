package service

import (
	"clone-academy/app/repository"
	"clone-academy/app/repository/models"
	"fmt"
	"net/url"
)

func GetCourses() []*models.Course {
	return repository.GetCourses()
}

func AddCourse(formValues *url.Values) {
	fmt.Println(formValues.Get("course-title"))
	fmt.Println(formValues.Get("course-description"))
}
