package logger

import (
	"log"
	"os"
)

type Logger struct {
	Info *log.Logger
	Err  *log.Logger
}

func Init() *Logger {
	logger := Logger{}
	logger.Info = log.New(os.Stdout, "[ИНФО]   ", log.Ldate|log.Ltime)
	logger.Err = log.New(os.Stderr, "[ОШИБКА] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Info.Println("Логгер инициализирован")
	return &logger
}
