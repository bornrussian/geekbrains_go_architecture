package email

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"shop/config"
	"shop/models"
	"strconv"
	"text/template"
)

type ShopEmailServer struct {
	SMTPServer 			string
	Port				int64
	From				string
	User				string
	Password			string
	MsgTmplFile			string
}

func NewShopEmailServer (conf config.EmailConfig) (*ShopEmailServer, error) {
	return &ShopEmailServer{
		SMTPServer: conf.SMTPServer,
		Port: conf.Port,
		User: conf.User,
		From: conf.From,
		Password: conf.Password,
		MsgTmplFile: conf.MsgTmplFile,
	}, nil
}

func (s *ShopEmailServer) SendOrderNotification (order *models.Order) error {
	t, err := template.ParseFiles(s.MsgTmplFile)
	if err != nil {
		return err
	}

	var text bytes.Buffer

	text.WriteString(fmt.Sprintf("From: '%v' <%v>\r\n", s.From, s.User))

	if err := t.Execute(&text, order); err != nil {
		return err
	}

	log.Println(text.String())

	serverWithPort := s.SMTPServer + ":" + strconv.Itoa(int(s.Port))
	auth := smtp.PlainAuth("", s.User, s.Password, s.SMTPServer)
	if err := smtp.SendMail(serverWithPort, auth, s.User, []string{order.Email}, []byte(text.String())); err != nil {
		return err
	}

	return nil
}
