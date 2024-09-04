package protocol

import (
	"github.com/AnumaNetwork/anumad/infrastructure/logger"
	"github.com/AnumaNetwork/anumad/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
