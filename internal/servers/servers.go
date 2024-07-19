package servers

import (
	"fmt"

	"github.com/ddreval/waytogo/internal/config"
	"github.com/ddreval/waytogo/internal/controllers"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type Server struct {
	cfg	*config.Config
	router *gin.Engine
}

func New (di *do.Injector) (*Server, error) {
	cfg, err := do.Invoke [*config.Config] (di)
	if err != nil {
		return nil, err
	}	
	router := gin.New()	
	controller, err := do.Invoke [*controllers.StaticController] (di)
	if err != nil {
		return nil, err
	}
	controller.Wire(router)
	server := &Server {cfg, router}
	return server, err
}
	
func (s *Server) Run() error {
	addr := fmt.Sprintf(":%d", s.cfg.Port)
	return s.router.Run(addr);	
}