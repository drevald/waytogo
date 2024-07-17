package injectors

import (
	"github.com/samber/do"
	"config"
)

var Default *do.Injector

func init() {
   Default = do.New()
   do.Provide(Default, config.New)
}