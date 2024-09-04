package grpcclient

import (
	"github.com/AnumaNetwork/anumad/infrastructure/logger"
	"github.com/AnumaNetwork/anumad/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
