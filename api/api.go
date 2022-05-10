package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/husanmusa/pro-book-service-api_gateway/api/docs"
	"github.com/husanmusa/pro-book-service-api_gateway/api/handlers"
	"github.com/husanmusa/pro-book-service-api_gateway/config"
	"github.com/saidamir98/udevs_pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Cfg config.Config
	Log logger.LoggerI
}

// SetupRouter godoc
// @description This is api gateway
// @termsOfService https://t.me/husanmusa
func SetupRouter(cfg config.Config, log logger.LoggerI) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	handler := handlers.New(cfg, log)
	url := swagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))

	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.GetUser)
	r.GET("/user", handler.ListUser)
	r.GET("/user/find_book/:id", handler.FindBookFromUsers)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	r.POST("/book", handler.CreateBook)
	r.GET("/book/:id", handler.GetBook)
	r.GET("/book", handler.ListBook)
	r.GET("/book/user/:id", handler.ListBooksByUserId)
	r.PUT("/book/:id", handler.UpdateBook)
	r.DELETE("/book/:id", handler.DeleteBook)

	r.POST("/sale", handler.CreateSaleBook)
	r.GET("/sale/:id", handler.GetSaleBook)
	r.GET("/sale", handler.ListSaleBook)
	r.PUT("/sale/:id", handler.UpdateSaleBook)
	r.DELETE("/sale/:id", handler.DeleteSaleBook)

	return r
}
