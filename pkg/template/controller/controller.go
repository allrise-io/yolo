package controller

var CONTROLLER_TPL = `package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"{{.ProjectName}}/pkg/service"
)

func Liveness(c *gin.Context) {
	pr := service.NewProbe()
	if err := pr.Liveness(); err != nil {
		Response(c, http.StatusNotAcceptable, err.Error())
		return
	}
	Response(c, http.StatusOK, "ok")
}
`
