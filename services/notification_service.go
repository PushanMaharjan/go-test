package services

import "fmt"

type NotificationService struct {
	gmailService GmailService
}

func NewNotificationService(
	gmailService GmailService,

) NotificationService {
	return NotificationService{
		gmailService: gmailService,
	}
}

func (ns NotificationService) WithTemplate(to, templatePrefix, templateName string, bindingData interface{}) error {

	fmt.Println(to, templatePrefix, templateName, bindingData)

	_, err := ns.gmailService.SendEmail(EmailParams{
		BodyTemplate:    templatePrefix + templateName + "_body.txt",
		BodyData:        bindingData,
		SubjectTemplate: templatePrefix + templateName + "_sub.txt",
		SubjectData:     bindingData,
		To:              to,
	})
	return err
}
