package injectors

import (
	"github.com/ddreval/waytogo/internal/config"
	"github.com/ddreval/waytogo/internal/controllers"
	"github.com/ddreval/waytogo/internal/servers"
	"github.com/ddreval/waytogo/internal/loggers"
	"github.com/samber/do"
)

var Default *do.Injector

func init() {
   Default = do.New()
   do.Provide(Default, config.New)
   do.Provide(Default, servers.New)
   do.Provide(Default, controllers.NewStatic)
   do.Provide(Default, loggers.New)
}