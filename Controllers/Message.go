package Controllers

import (
	"e-shop/Forms"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nexmo-community/nexmo-go"
	"log"
	"net/http"
)

func SendMessages(c *gin.Context) {
	var err error
	// Auth
	auth := nexmo.NewAuthSet()
	auth.SetAPISecret("5a8f4581", "knlI7i5383mfDWNM")

	// Init Nexmo
	client := nexmo.NewClient(http.DefaultClient, auth)

	// Bind our form

	var StkPushForm Forms.StkPushForm
	err = c.BindJSON(&StkPushForm)
	if err != nil {
		appError := err
		appError.Error()
		return
	}

	m := &StkPushForm
	smsContent := nexmo.SendSMSRequest{
		From: "E-Shop",
		To:   m.PartyA,
		Text: "Your payment of" + " " + m.Amount + " " + "has been received",
	}
	// SMS

	smsResponse, _, err := client.SMS.SendSMS(smsContent)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status:", smsResponse.Messages[0].Status)
}
