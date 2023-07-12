package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FormStatusEmail struct {
	bun.BaseModel    `bun:"form_status_emails,alias:form_status_emails"`
	ID               uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Email            string    `json:"email"`
	Theme            string    `json:"theme"`
	TemplateFileName string    `json:"templateFileName"`

	FormStatus   *FormStatus   `bun:"rel:belongs-to" json:"formStatus"`
	FormStatusID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formStatusId"`
}

type FormStatusEmails []*FormStatusEmail

type FormStatusEmailsWithCount struct {
	FormStatusEmails FormStatusEmails `json:"items"`
	Count            int              `json:"count"`
}

type EmailSender interface {
	SendEmail(to []string, subject string, body string) error
}

type TemplateParser interface {
	ParseTemplate(data interface{}, templates ...string) (string, error)
}

func (item *FormStatusEmail) Send(formValue *FormValue, sender EmailSender, templateParser TemplateParser) error {
	emailStruct := struct {
		FormValue *FormValue
	}{
		formValue,
	}
	mail, err := templateParser.ParseTemplate(emailStruct, "email/"+item.TemplateFileName)
	if err != nil {
		return err
	}
	err = sender.SendEmail([]string{item.Email}, item.Theme, mail)
	if err != nil {
		return err
	}
	return nil
}
