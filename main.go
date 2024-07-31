package main

import (
	"github.com/ddreval/waytogo/internal/injectors"
	"github.com/ddreval/waytogo/internal/servers"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
)

func main() {
	logger, err := do.Invoke[*logrus.Logger](injectors.Default)
	if err != nil {
		return
	}
	server, err := do.Invoke[*servers.Server](injectors.Default)
	if err != nil {
		logger.Error(err)
	} 
	server.Run()
}
