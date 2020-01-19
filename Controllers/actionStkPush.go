package Controllers

import (
	"e-shop/Config"
	"e-shop/Models"
	"e-shop/Utils"
	"e-shop/Web/auth"
	"encoding/json"
	"github.com/AndroidStudyOpenSource/mpesa-api-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

const (
	appKey    = "ZVvVF1As0ugLQfEKfPVqK18lEfZAQIjM"
	appSecret = "88S1bO4FerQMwM6G"
)

func MpesaExpress(c *gin.Context) {
	//clear previous error if any
	errList := map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	transaction := Models.MpesaStkPush{}

	err = json.Unmarshal(body, &transaction)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	// check if the user exist:
	user := Models.User{}
	err = Config.DB.Debug().Model(Models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	transaction.AuthorID = uid //the authenticated user is the one creating the product

	paymentSend, err := transaction.SaveMpesaTrasaction(Config.DB)
	if err != nil {
		errList := Utils.FormatError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	svc, err := mpesa.New(appKey, appSecret, mpesa.SANDBOX)
	if err != nil {
		panic(err)
	}

	res, err := svc.Simulation(mpesa.Express{
		BusinessShortCode: "174379",
		Password:          "MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMTgwNDA5MDkzMDAy",
		Timestamp:         "20180409093002",
		TransactionType:   "CustomerPayBillOnline",
		Amount:            transaction.Amount,
		PartyA:            transaction.PartyA,
		PartyB:            "174379",
		PhoneNumber:       transaction.PartyA,
		CallBackURL:       "https://wamu.co.ke",
		AccountReference:  "account",
		TransactionDesc:   "test",
	})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Failed to pay",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success",
		"res":     res,
		"results":  paymentSend,
	})

}
