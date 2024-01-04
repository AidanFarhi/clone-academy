package repository

import (
	models "clone-academy/app/repository/models"
)

func GetCourses() []*models.Course {
	return []*models.Course{
		{
			Title:       "7th Grade Math",
			Description: "Intro to Algebra, Geometry, and others.",
			Completion:  "89%",
		},
		{
			Title:       "7th Grade English",
			Description: "Essays, poetry, and other stuff.",
			Completion:  "78.3%",
		},
		{
			Title:       "7th Grade Biology",
			Description: "Cells, photosynthesis, and others.",
			Completion:  "96.4%",
		},
	}
}
