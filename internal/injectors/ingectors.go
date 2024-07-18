package injectors

import (
	"github.com/samber/do"
	"github.com/ddreval/waytogo/internal/config"
	"github.com/ddreval/waytogo/internal/servers"
)

var Default *do.Injector

func init() {
   Default = do.New()
   do.Provide(Default, config.New)
   do.Provide(Default, server.New)
}