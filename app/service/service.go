package service

type Course struct {
	Title       string
	Description string
	Completion  string
}

func GetCourses() []*Course {
	return []*Course{
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
