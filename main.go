package main

import (
	"e-shop/Config"
	"e-shop/Models"
	"e-shop/Web/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var err error

func main()  {
	//Creating connection to the database
	Config.DB, err = gorm.Open("mysql",Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("status: ", err)
	}

	defer Config.DB.Close()

	// Migrations
	Config.DB.AutoMigrate(&Models.User{})
	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Subscribe{})
	Config.DB.AutoMigrate(&Models.Address{})
	Config.DB.AutoMigrate(&Models.MpesaStkPush{})

	// Setup routes
	r := router.SetupRoutes()
	r.Use(gin.Recovery())
	// running
	_ = r.Run(":8000")

}
