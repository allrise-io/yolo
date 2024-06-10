package config

var LOG_TPL = `package config

import (
	"io"
	"os"
	"path"
	"runtime"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func InitLogger(logFile string) {
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	writers := []io.Writer{f, os.Stdout}
	allWriters := io.MultiWriter(writers...)

	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return frame.Function, filename + ":" + strconv.Itoa(frame.Line)
		},
	})
	log.SetOutput(allWriters)
	log.SetLevel(log.InfoLevel)
}
`
