package config

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var createTableStatement string = `
	CREATE TABLE IF NOT EXISTS course (
		title TEXT,
		description TEXT,
		completion DECIMAL(3, 2)
	)
`

var insertStatement string = `
	INSERT INTO 
		course (title, description, completion)
	VALUES
		('7th Grade Math', 'Intro to Algebra, Geometry, and others.', 89),
		('7th Grade English', 'Essays, poetry, and other stuff.', 78.3),
		('7th Grade Biology', 'Cells, photosynthesis, and others.', 96.4)
`

func InitDatabase() {
	// Open or create the SQLite database file
	conn, err := sql.Open("sqlite", "./clone_academy.db")

	// Create table
	conn.Exec(createTableStatement)

	// Insert some data
	conn.Exec(insertStatement)

	if err != nil {
		panic(err)
	}
	fmt.Println("Succefully connected/created db")
}
