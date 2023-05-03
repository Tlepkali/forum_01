package logger

import (
	"log"
	"os"
	"sync"
)

type Level uint8

const (
	LevelInfo Level = iota
	LevelError
	LevelFatal
	LevelOff
)

type Logger struct {
	minLevel Level
	logTypes map[string]*log.Logger
}

var (
	once   sync.Once
	logger *Logger
)

func GetLoggerInstance() *Logger {
	once.Do(func() {
		logger = &Logger{}
		logger.init()
	})
	return logger
}

func (l *Logger) PrintInfo(message string) {
	l.print(LevelInfo, message)
}

func (l *Logger) PrintError(err error) {
	l.print(LevelError, err.Error())
}

func (l *Logger) PrintFatal(err error) {
	l.print(LevelFatal, err.Error())
}

func (l *Logger) print(level Level, message string) {
	switch level {
	case LevelInfo:
		l.logTypes["info"].Printf(message)
	case LevelError:
		l.logTypes["error"].Printf(message)
	case LevelFatal:
		l.logTypes["fatal"].Fatalf(message)
	default:
	}
}

func (l *Logger) init() {
	err := os.MkdirAll("logs", 0o755)

	if err != nil || os.IsExist(err) {
		panic("can't create log dir, no configured logging to files")
	}

	infoF, err := os.OpenFile("logs/info.log", os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		log.Fatal(err)
	}

	errorF, err := os.OpenFile("logs/error.log", os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		log.Fatal(err)
	}

	fatalF, err := os.OpenFile("logs/fatal.log", os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		log.Fatal(err)
	}

	logTypes := make(map[string]*log.Logger)

	infoLog := log.New(infoF, "INFO\t", log.Ldate|log.Ltime)
	logTypes["info"] = infoLog

	errorLog := log.New(errorF, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	logTypes["error"] = errorLog

	fatalLog := log.New(fatalF, "FATAL\t", log.Ldate|log.Ltime|log.Lshortfile)
	logTypes["fatal"] = fatalLog

	l.logTypes = logTypes
	l.minLevel = LevelInfo
}
