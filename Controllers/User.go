package Controllers

import (
	"e-shop/Config"
	"e-shop/Models"
	"e-shop/Utils"
	"e-shop/Web/auth"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

func CreateUser(c *gin.Context)  {

	//clear previous erro if any
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

	user := Models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		errList["Unmarshal_error"] = "Unable to unmarshal the body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	user.Prepare()
	errorMessages := user.Validate("")
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	userCreated, err := user.SaveUser(Config.DB)
	if err != nil {
		formattedError := Utils.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error": errList,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": userCreated,
	})
}

func Login(c *gin.Context) {
	//check rprevious errors if any
	errList := map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to read body"
		c.JSON(http.StatusUnprocessableEntity, gin.H {
			"status": http.StatusUnprocessableEntity,
			"error": errList,
		})
		return
	}
	// Unmarshal body
	user := Models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		errList["Unmarshal_error"] ="Unable to unmarchal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H {
			"status": http.StatusUnprocessableEntity,
			"error": errList,
		})
		return
	}
	user.Prepare()
	errorMessages := user.Validate("login")
	if len(errorMessages) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H {
			"status": http.StatusUnprocessableEntity,
			"error": errorMessages,
		})
		return
	}
	userData, err := SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := Utils.FormatError(err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
		"status": http.StatusUnprocessableEntity,
		"error":  formattedError,
	})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": userData,
	})

}

func SignIn(email, password string) (map[string]interface{}, error) {
	var err error

	userData := make(map[string]interface{})

	user := Models.User{}

	err = Config.DB.Debug().Model(user).Where("email = ?", email).Take(&user).Error
	if err != nil {
		fmt.Println("this is the error getting the user: ", err)
		return nil, err
	}
	err = Utils.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("this is the error hashing the password: ", err)
		return nil, err
	}
	token, err := auth.CreateToken(user.ID)
	if err != nil {
		fmt.Println("this is the error creating the token: ", err)
		return nil, err
	}
	userData["token"] = token
	userData["id"] = user.ID
	userData["email"] = user.Email
	userData["username"] = user.Username

	return userData, nil
}