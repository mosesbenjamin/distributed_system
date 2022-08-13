package portal

import (
	"html/template"
)

var rootTemplate *template.Template

func ImportTemplates() error {
	var err error
	rootTemplate, err = template.ParseFiles(
		"portal/students.gohtml",
		"portal/student.gohtml",
		"portal/grades.gohtml")

	if err != nil {
		return err
	}

	return nil
}
