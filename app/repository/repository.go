package repository

import (
	models "clone-academy/app/repository/models"
	"database/sql"

	_ "modernc.org/sqlite"
)

func GetCourses() []*models.Course {

	allCourses := make([]*models.Course, 0)

	conn, _ := sql.Open("sqlite", "./clone_academy.db")
	rows, _ := conn.Query("select * from course")

	for rows.Next() {
		var course models.Course
		rows.Scan(&course.Title, &course.Description, &course.Completion)
		allCourses = append(allCourses, &course)
	}

	return allCourses
}
