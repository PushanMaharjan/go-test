package services

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"

	"go-fx-test/lib"
	"go-fx-test/utils"
)

// GmailService -> GmailService structure
type GmailService struct {
	gmailService *gmail.Service
	lib          lib.Env
}

// NewGmailService -> Constructor Function
func NewGmailService(
	env lib.Env,
) GmailService {

	oauthConfig := oauth2.Config{
		ClientID:     env.MailClientId,
		ClientSecret: env.MailClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
	}

	token := oauth2.Token{
		// AccessToken:  env.MailAccessToken,
		RefreshToken: env.MailRefreshToken,
		TokenType:    "Bearer",
		Expiry:       time.Now(),
	}

	var tokenSource = oauthConfig.TokenSource(context.Background(), &token)

	srv, err := gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))

	if err != nil {
		log.Println("failed to retrieve gmail client", err.Error())
	}

	log.Println("gmail client created successfully")

	return GmailService{
		gmailService: srv,
	}
}

type EmailParams struct {
	To              string
	Subject         string
	SubjectData     interface{}
	SubjectTemplate string
	Body            string
	BodyData        interface{}
	BodyTemplate    string
}

// Send Email -> sends email with provided data
func (g GmailService) SendEmail(params EmailParams) (bool, error) {
	isEmailSubjectEmpty := utils.IsInterfaceEmpty(params.SubjectData)
	var err error

	if !isEmailSubjectEmpty {
		params.Subject, err = utils.ParseTemplate(params.SubjectTemplate, params.SubjectData)
		if err != nil {
			return false, errors.New("unable to parse email subject template")
		}
	}

	emailBody, err := utils.ParseTemplate(params.BodyTemplate, params.BodyData)
	if err != nil {
		return false, errors.New("unable to parse email body template")
	}

	var message gmail.Message
	var msg []byte

	emailTo := "To: " + params.To + "\r\n"
	subject := "Subject: " + params.Subject + "\r\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"

	msgStr := emailTo + subject + mime + "\n" + emailBody
	msg = []byte(msgStr)

	message.Raw = base64.URLEncoding.EncodeToString(msg)

	fmt.Println("here")

	// Send the message
	_, err = g.gmailService.Users.Messages.Send("me", &message).Do()

	fmt.Println("herererererer", err)
	if err != nil {
		log.Printf("Err: %v", err)
		return false, err
	}

	fmt.Println("Message sent!")
	return true, nil
}
