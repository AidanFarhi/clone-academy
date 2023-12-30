package service

type Course struct {
	Title       string
	Description string
}

func GetCourses() []*Course {
	return []*Course{
		{Title: "Algebra", Description: "Intro to Algebra"},
	}
}
