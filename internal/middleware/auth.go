package middleware

import (
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
	auth.logger.Info("Authenticate() called")
	c.Next()
}
