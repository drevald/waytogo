package injectors

import (
	"github.com/samber/do"
	"github.com/ddreval/waytogo/internal/config"
)

var Default *do.Injector

func init() {
   Default = do.New()
   do.Provide(Default, config.New())
}