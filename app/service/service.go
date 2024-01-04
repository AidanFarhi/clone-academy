package service

import (
	"clone-academy/app/repository"
	"clone-academy/app/repository/models"
)

func GetCourses() []*models.Course {
	return repository.GetCourses()
}
