package router

var MID_TPL = `package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"{{.ProjectName}}/pkg/model"
	"{{.ProjectName}}/pkg/util/cm"
)

func checkAuth(c *gin.Context) {
	c.Set("username", "")
	c.Set("roles", "")
}

func DefaultAuth(c *gin.Context) {
	checkAuth(c)
	c.Next()
}

func RoleAuth(c *gin.Context) {
	checkAuth(c)
	roles := c.MustGet("roles").(string)
	roleList := strings.Split(roles, ",")
	fmt.Println("roles: ", roleList)
	if !cm.In(model.ROLE_ADMIN, roleList) && !cm.In(model.ROLE_OP, roleList) && !cm.In(model.ROLE_RD, roleList) {
		c.JSON(http.StatusForbidden, "403 forrbidden!")
		c.Abort()
		return
	}
	c.Next()
}
`
