package main

import (
	"fmt"
	"github.com/AnumaNetwork/anumad/version"
	"os"
	"path/filepath"
	"strings"
)

func showVersion() {
	appName := filepath.Base(os.Args[0])
	appName = strings.TrimSuffix(appName, filepath.Ext(appName))
	fmt.Println(appName, "version", version.Version())
}
