package main

import (
	"fmt"
	"log"
	"github.com/ddreval/waytogo/internal/config"
	"github.com/ddreval/waytogo/internal/servers"
	"github.com/ddreval/waytogo/internal/injectors"
	"github.com/samber/do"

)

func main() {
	cfg, err := do.Invoke[*config.Config] (injectors.Default)
	if err != nil {
		log.Fatalln(err)		
	}	
	fmt.Println(cfg.Port)
	server, err := do.Invoke[*servers.Server] (injectors.Default )
	server.Run()
}