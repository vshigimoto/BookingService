package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/test", test)

	err := r.Run(":8085")
	if err != nil {
		return
	}
}

// CreateTags		godoc
// @Summary			Create tags
// @Description		Save tags data in Db.
// @Param			tags body string true "Create tags"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} string
// @Router			/tags [post]
func test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}
