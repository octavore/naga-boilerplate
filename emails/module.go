package emails

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/keighl/postmark"
	"github.com/octavore/naga/service"
	"github.com/octavore/nagax/config"
	"github.com/octavore/nagax/logger"
)

type emailClient interface {
	SendEmail(postmark.Email) (postmark.EmailResponse, error)
}

type dummyEmailClient struct {
	logger *logger.Module
}

func (d *dummyEmailClient) SendEmail(email postmark.Email) (postmark.EmailResponse, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		d.logger.Error(err)
	} else {
		fmt.Fprint(f, email.HtmlBody)
		email.HtmlBody = ""
		f.Close()
		os.Rename(f.Name(), f.Name()+".html")
		d.logger.Infof("%s %+v", f.Name()+".html", email)
	}

	return postmark.EmailResponse{}, nil
}

// PostmarkConfig contains credentials for postmark
type PostmarkConfig struct {
	Postmark struct {
		ServerToken  string `json:"server_token"`
		AccountToken string `json:"account_token"`
	} `json:"postmark"`
}

// Module emails handles email delivery for your app.
type Module struct {
	Config *config.Module
	Logger *logger.Module

	templates      map[string]*template.Template
	config         PostmarkConfig
	postmarkClient emailClient
}

func (m *Module) Init(c *service.Config) {
	c.Setup = func() error {
		err := m.Config.ReadConfig(&m.config)
		if err != nil {
			return err
		}
		if m.config.Postmark.ServerToken != "" &&
			m.config.Postmark.AccountToken != "" {
			m.postmarkClient = postmark.NewClient(
				m.config.Postmark.ServerToken,
				m.config.Postmark.AccountToken,
			)
		} else {
			m.postmarkClient = &dummyEmailClient{m.Logger}
		}

		m.templates = map[string]*template.Template{}
		for _, file := range AssetNames() {
			name := strings.TrimRight(path.Base(file), ".html")
			data, err := Asset(file)
			if err != nil {
				return err
			}
			m.templates[name], err = template.New(name).Parse(string(data))
			if err != nil {
				return err
			}
		}

		return nil
	}
}

// SendResetPasswordEmail sends a password reset email to toAddress containing
// the resetToken.
func (m *Module) SendResetPasswordEmail(toAddress, resetToken string) {
	resetPasswordLink := "https://www.example.com/reset-password?t=" + resetToken
	e := postmark.Email{
		To:      toAddress,
		Subject: "Your password reset instructions",
		TextBody: fmt.Sprintf("Please visit %s to reset your password. "+
			"If you did not request a password reset, ignore this email.",
			resetPasswordLink),
	}

	err := m.SendEmail("reset-password", e, map[string]string{
		"resetPasswordLink": resetPasswordLink,
	})
	if err != nil {
		m.Logger.Error("error sending reset password email:", err)
	}
}

func (m *Module) SendEmail(tmplName string, email postmark.Email, data interface{}) error {
	defer func() {
		if err := recover(); err != nil {
			m.Logger.Error(err)
		}
	}()
	email.From = "Support <no-reply@example.com>"
	if tmplName != "" {
		tmpl, ok := m.templates[tmplName]
		if !ok {
			return fmt.Errorf("email template %s not found", tmplName)
		}
		w := &bytes.Buffer{}
		err := tmpl.Execute(w, data)
		if err != nil {
			return err
		}
		email.HtmlBody = w.String()
	}
	_, err := m.postmarkClient.SendEmail(email)
	return err
}
