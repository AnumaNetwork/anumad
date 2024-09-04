package main

import (
	"github.com/AnumaNetwork/anumad/infrastructure/logger"
	"github.com/AnumaNetwork/anumad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("APLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
