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
	fmt.Println("AAA")
	val := session.Get("user")
	if c.Request.URL.Path == "/login" || c.Request.URL.Path == "/register" || c.Request.URL.Path == "/favicon.ico" {
		fmt.Println("BBB")
		auth.logger.Debug("logger page ok")
		c.Next()
	} else if val != nil { 
		fmt.Println("CCC")
		user := val.(*databases.User)
		auth.logger.Info(fmt.Scanf("Logged as %s", user.Username))
		c.Next()
	} else {
		fmt.Println("DDD")
		auth.logger.Debug("redirecting to logger page")
		c.Redirect(301, "login")
	}	
	fmt.Println("EEE")
}