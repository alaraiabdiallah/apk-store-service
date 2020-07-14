package tests

import (
	"github.com/alaraiabdiallah/apk-store-service/app"
	"testing"
)

func TestGetTransport(t *testing.T) {

	res, err := app.GetTransport("transporter-1")
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}

func TestAddTransport(t *testing.T) {
	res, err := app.NewTransport(map[string]interface{}{
		"name": "transport 2",
		"type": "SMTP",
		"config": map[string]interface{}{},
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}