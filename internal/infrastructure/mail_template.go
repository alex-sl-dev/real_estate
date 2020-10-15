package infrastructure

import (
	"errors"
	"html/template"
	"log"
	"os"
	"untitled/internal/domain"
)

type MailTemplateHandler struct {

}

func (mt *MailTemplateHandler)ProcessMailTemplate(mailTemplate domain.MailTemplate) error {

	var files []string

	pwd := os.Getenv("APP_ROOT_PATH")

	switch mailTemplate.Role {
	case "registration":
		files = append(files, pwd + "/templates/mails/registration.tmpl")
		break
	default:
		return errors.New("role for template is missing")
	}

	// Initialize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must be the *first* file in the slice.
	files = append(files, pwd + "/templates/mails/base.tmpl")

	//
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	//
	err = ts.Execute(mailTemplate.Content, mailTemplate.Variables)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}