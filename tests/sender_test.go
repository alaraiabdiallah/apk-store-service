package tests

import (
	"github.com/alaraiabdiallah/apk-store-service/app"
	"testing"
)

func TestSendSMTP(t *testing.T) {
	trans, err := app.GetTransport("mailstrap")
	if err != nil { t.Error(err) }
	dataTrans := trans.(map[string]interface{})
	err = app.SendSMTP(dataTrans["config"].(map[string]interface{}), app.SMTPBody{
		To: "spensa2010alarai@gmail.com",
		From: "alaraiabdiallah@gmail.com",
		Subject: "Test Email",
		Body: "<h2>Test email</h2>",
	})
	if err != nil { t.Error(err) }
}
