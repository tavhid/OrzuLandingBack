package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

// InitLogger ...
func InitLogger() (log *logrus.Logger) {
	f, err := os.OpenFile("../log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	log = &logrus.Logger{
		// Log into f file handler and on os.Stdout
		Out:   io.MultiWriter(f, os.Stdout),
		Level: logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			FullTimestamp: true,
			// ForceColors:   true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				return fmt.Sprintf("-> %s:%d", f.File, f.Line), fmt.Sprintf("[%s]", f.Function)
			},
		},
	}
	log.SetReportCaller(true)

	return
}
