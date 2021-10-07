package services

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

func (ns NotificationService) SendEmailWithTemplate(to, templatePrefix, templateName string, bindingData interface{}) error {
	_, err := ns.gmailService.SendEmail(EmailParams{
		BodyTemplate:    templatePrefix + templateName + "_body" + ".txt",
		BodyData:        bindingData,
		SubjectTemplate: templatePrefix + templateName + "_sub" + ".txt",
		SubjectData:     bindingData,
		To:              to,
	})
	return err
}
