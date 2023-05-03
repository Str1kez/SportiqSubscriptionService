package logger

import (
	"os"
	"path"
	"runtime"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func callerPrettyfier(frame *runtime.Frame) (function string, file string) {
	fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
	return "", fileName
}

func InitLogger() {
	log.SetOutput(os.Stdout)
	if os.Getenv("GIN_MODE") == "release" {
		log.SetFormatter(&log.JSONFormatter{CallerPrettyfier: callerPrettyfier})
		log.SetLevel(log.WarnLevel)
	} else {
		log.SetFormatter(&log.TextFormatter{FullTimestamp: true, CallerPrettyfier: callerPrettyfier})
		log.SetLevel(log.DebugLevel)
	}
}
