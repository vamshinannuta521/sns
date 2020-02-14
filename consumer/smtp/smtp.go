package smtp

import (
	// "fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"

	"consumer/models"

	"github.com/sirupsen/logrus"
)

var logger = logrus.NewEntry(logrus.New())

var smtpAddr string

func init() {
	curDir, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}
	fileByte, err := ioutil.ReadFile(curDir + "/smtp/ip.txt")
	if err != nil {
		logger.Fatal(err)
	}

	smtpAddr = string(fileByte)
}

func Send(smtpObj *models.EmailActionSpec) error {
	from := smtpObj.From
	to := strings.Split(smtpObj.To, ",")

	msg := "Subject: " + smtpObj.Subject + "\r\n" +
		"\r\n" +
		smtpObj.Body

	err := smtp.SendMail(smtpAddr, nil, from, to, []byte(msg))

	if err != nil {
		logger.Errorf("smtp error: %s", err)
		return err
	}
	return nil

}
