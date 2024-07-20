package main

import (
	"fmt"
	"github.com/ddreval/waytogo/internal/config"
	"github.com/ddreval/waytogo/internal/servers"
	"github.com/ddreval/waytogo/internal/injectors"
	"github.com/sirupsen/logrus"
	"github.com/samber/do"

)

func main() {
	logger, err := do.Invoke[*logrus.Logger] (injectors.Default)
	if err != nil {
		fmt.Println(err)
	}
	cfg, err := do.Invoke[*config.Config] (injectors.Default)
	if err != nil {
		logger.Info("config not parsed")
	}
	logger.Info("config parsed")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	fmt.Println(cfg.Port)
	server, err := do.Invoke[*servers.Server] (injectors.Default )
	server.Run()
}