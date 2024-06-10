package router

var URL_TPL = `package router

import (
	"github.com/gin-gonic/gin"

	"{{.ProjectName}}/pkg/controller"
)

func URLs(r *gin.Engine) {

	read := r.Group("/v1", DefaultAuth)
	{
		read.GET("/liveness", controller.Liveness)
	}

	write := r.Group("/v1", RoleAuth)
	{
		write.POST("/liveness", controller.Liveness)
	}
}
`
