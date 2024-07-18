package servers

import (
	"github.com/ddreval/waytogo/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"fmt"
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
	router.GET("", func(c *gin.Context){
		c.String(200, "hello")
	})
	server := &Server {cfg, router}
	return server, err
}
	
func (s *Server) Run() error {
	addr := fmt.Sprintf(":%d", s.cfg.Port)
	return s.router.Run(addr);	
}