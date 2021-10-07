package services

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"go-fx-test/lib"
	"go-fx-test/utils"
	"log"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type GmailService struct {
	gmailService *gmail.Service
	env          lib.Env
}

func NewGmailService(
	env lib.Env,
) GmailService {
	ctx := context.Background()
	oauthConfig := oauth2.Config{
		ClientID:     env.MailClientId,
		ClientSecret: env.MailClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
		Scopes:       []string{"https://www.googleapis.com/auth/gmail.send"},
	}

	expiry, _ := time.Parse("2006-01-02", "2018-04-16")
	token := oauth2.Token{
		RefreshToken: env.MailRefreshToken,
		TokenType:    "Bearer",
		Expiry:       expiry,
	}

	var tokenSource = oauthConfig.TokenSource(ctx, &token)

	srv, err := gmail.NewService(ctx, option.WithTokenSource(tokenSource))

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

	emailTo := "To: " + params.To + "\r\n"
	subject := "Subject: " + params.Subject + "\r\n"

	msgStr := emailTo + subject + "\n" + emailBody
	var msg []byte

	msg = []byte(msgStr)
	message.Raw = base64.URLEncoding.EncodeToString(msg)

	// Send the message
	_, err = g.gmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Printf("Error: %v", err)
		return false, err
	}

	fmt.Println("Message sent!")
	return true, nil
}
