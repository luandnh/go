package api

import (
	"go-project/common/response"
	authMdw "go-project/middleware/auth"

	"github.com/gin-gonic/gin"
)

type Auth struct {
}

func NewAuth(r *gin.Engine) {
	auth := &Auth{}
	group := r.Group("v1/auth")
	{
		group.GET("/check", authMdw.AuthMiddleware(), auth.CheckAuthen)
		group.POST("/token", auth.GenerateToken)
	}
}

func (auth *Auth) GenerateToken(c *gin.Context) {
	data := map[string]interface{}{
		"id":    "abc",
		"level": "admin",
	}
	token, err := authMdw.GenerateJWT(data)
	if err != nil {
		c.JSON(response.ServiceUnavailableMsg(err.Error()))
		return
	}
	c.JSON(response.OK(map[string]interface{}{"token": token}))
}

func (auth *Auth) CheckAuthen(c *gin.Context) {

	c.JSON(response.OK(nil))
}
