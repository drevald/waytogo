package middleware

import (
	"encoding/json"
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
	c.Next()
	session := sessions.Default(c)
		val := session.Get("user")
	if c.Request.URL.Path == "/login" || c.Request.URL.Path == "/register" || c.Request.URL.Path == "/favicon.ico" {
		auth.logger.Warn("logger page ok")
		c.Next()
	} else if val != nil {
		var user databases.User 
		userBytes := val.([]byte)
		userString := string(userBytes)
		json.Unmarshal([]byte(userString), &user)
		auth.logger.Infof("Logged as %s", user.Username)
		c.Next()
	} else {
		auth.logger.Error("Unauthorized access - redirecting to login")
		c.Redirect(301, "login")
	}	
}