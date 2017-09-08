package server

import (
	"github.com/koudaiii/sample-go-api/middleware"
	"github.com/koudaiii/sample-go-api/router"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.SetDBtoContext(db))
	router.Initialize(r)
	return r
}
