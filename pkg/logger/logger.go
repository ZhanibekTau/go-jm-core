package logger

import (
	"go-jm-core/pkg/exception"
	"log"
	"os"
)

func Info(format string, v ...any) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf(format, v)
}

func Error(format string, v ...any) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf(format, v)
}

func LogError(err error) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog.Println(err)
}

func LogAppException(appException *exception.AppException) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog.Println(appException.Error.Error())
}
