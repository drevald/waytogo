package middleware

import (
	"fmt"
	"github.com/ddreval/waytogo/internal/databases"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	logger *logrus.Logger
}

func NewAuth(di *do.Injector) (*Auth, error) {
	logger, _ := do.Invoke[*logrus.Logger](di)
	return &Auth{logger}, nil
}

func (auth *Auth) Authenticate(c *gin.Context) {

	auth.logger.Info("Authenticate() called with c.Request.URL.Path = " + c.Request.URL.Path)

	session := sessions.Default(c)
	user := session.Get("user").(*databases.User)

	if c.Request.URL.Path == "/login" || c.Request.URL.Path == "/register" || c.Request.URL.Path == "/favicon.ico" {
		auth.logger.Debug("logger page ok")
		c.Next()
	} else if user != nil { 
		auth.logger.Info(fmt.Scanf("Logged as %s", user.Username))
		c.Next()
	} else {
		auth.logger.Debug("redirecting to logger page")
		c.Redirect(301, "login")
	}	
}
