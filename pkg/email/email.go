package email

import (
	"net/smtp"
)

type EmailAuth struct {
	UserName string
	Password string
	EmailHost
}

type EmailHost struct {
	Host  string
	Port  int
	IsSSL bool
	From  string
}

func genLoginAuth(username, password string) smtp.Auth {
	return &EmailAuth{UserName: username, Password: password}
}

func NewEmailC(username, password string, v ...interface{}) *EmailAuth {
	return &EmailAuth{
		UserName: username,
		Password: password,
		EmailHost: EmailHost{
			From: username,
		},
	}
}

func (a *EmailAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.UserName), nil
}

func (a *EmailAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username":
			return []byte(a.UserName), nil
		case "Password":
			return []byte(fromServer), nil
		}
	}
	return nil, nil
}

func (a *EmailAuth) SendMail(target string, body string, subject string) error {
	auth := genLoginAuth(a.UserName, a.Password)

	contentType := "Content-Type: text/plain" + "; charset=UTF-8"
	msg := []byte("To: " + target +
		"\r\nFrom: " + a.UserName +
		"\r\nSubject: " + subject +
		"\r\n" + contentType + "\r\n\r\n" +
		body)
	err := smtp.SendMail(a.Host, auth, a.UserName, []string{target}, msg)
	if err != nil {
		return err
	}
	return nil
}
