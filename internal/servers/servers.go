package servers

import (
	"fmt"
	"github.com/ddreval/waytogo/internal/config"
	"github.com/ddreval/waytogo/internal/controllers"
	"github.com/ddreval/waytogo/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

type Server struct {
	cfg    *config.Config
	router *gin.Engine
	logger *logrus.Logger
}

func New(di *do.Injector) (*Server, error) {
	cfg, err := do.Invoke[*config.Config](di)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	logger, err := do.Invoke[*logrus.Logger](di)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	router := gin.New()
	store := cookie.NewStore([]byte("secret"))
  	router.Use(sessions.Sessions("mysession", store))
	router.Use(gin.Recovery())
	auth, err := do.Invoke[*middleware.Auth](di)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}	
	router.Use(auth.Authenticate)
	tcontroller, err := do.Invoke[*controllers.TemplateController](di)
	if err != nil {
		logger.Info("Failed to get Template Controller")
		return nil, err
	}
	tcontroller.Wire(router)
	scontroller, err := do.Invoke[*controllers.StaticController](di)
	if err != nil {
		logger.Info("Failed to get Static Controller")
		return nil, err
	}
	scontroller.Wire(router)
	server := &Server{cfg, router, logger}
	return server, err
}

func (s *Server) Run() error {
	fmt.Println("Test")
	s.logger.Info("test")
	addr := fmt.Sprintf(":%d", s.cfg.Port)
	s.logger.Info(fmt.Sprintf("running server on port %d", s.cfg.Port))
	return s.router.Run(addr)
}
