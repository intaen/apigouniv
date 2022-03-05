package main

import (
	mc "apigouniv/src/masterdata/controller"
	mdr "apigouniv/src/masterdata/repo"
	mds "apigouniv/src/masterdata/service"

	"apigouniv/src/util"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "apigouniv/src/docs"
)

func init() {
	viper.SetConfigFile("config/Config.json")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(err)
	}
}

// @title GOUNIV
// @version 1.0
// @description This page is API documentation for master data of university in the world
// @schemes http
// @host localhost:3040
// @contact.name Developer
// @contact.email intanmarsjaf@outlook.com
func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Handle wrong method
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		util.HandleError(c, http.StatusMethodNotAllowed, "045", "Method Not Allowed", "Method Not Allowed", "Method Not Allowed")
	})
	// Handle no route
	r.NoRoute(func(c *gin.Context) {
		util.HandleError(c, http.StatusNotFound, "044", "Page Not Found", "Page Not Found", "Page Not Found")
	})

	// Swagger
	url := ginSwagger.URL(viper.GetString("host") + ":" + viper.GetString("port") + "/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Initiate Repo
	mdRepo := mdr.CreateMasterDataRepoImpl()

	// Initiate Service
	mdService := mds.CreateMasterDataServiceImpl(mdRepo)

	// Initiate Controller
	mc.CreateMasterDataController(r, mdService)

	r.Run(":" + viper.GetString("port"))
	// r.Run() // Heroku will supply automatically
}
