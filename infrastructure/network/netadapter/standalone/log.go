package standalone

import (
	"github.com/AnumaNetwork/anumad/infrastructure/logger"
	"github.com/AnumaNetwork/anumad/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
