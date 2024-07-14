package injectors

import (
	"github.com/samber/do"
)

var Default *do.Injector

func init() {
   Default = do.New()
}