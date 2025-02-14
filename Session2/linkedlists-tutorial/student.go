package linkedlist

import "fmt"

type Student struct {
	ID    int64
	Name  string
	Level EducationLevel
}

func NewStudent(id int64, name string, level EducationLevel) Student {
	return Student{
		ID:    id,
		Name:  name,
		Level: level,
	}
}

func StudentToString(student Student) string {
	return fmt.Sprintf("{id: %d, name: %q, level: %s}",
		student.ID,
		student.Name,
		LevelToString(student.Level))
}
