package jLog

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func PrintInfoJLog(serviceName string, method string, uri string, status int, requestId string, message string) {
	PrintJLog("INFO", serviceName, method, uri, status, requestId, message)
}

func PrintErrorJLog(serviceName string, method string, uri string, status int, requestId string, message string) {
	PrintJLog("ERROR", serviceName, method, uri, status, requestId, message)
}

func PrintJLog(level string, serviceName string, method string, uri string, status int, requestId string, message string) {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	messageBuilder := strings.Builder{}
	messageBuilder.WriteString(time.Now().Format("2006-01-02 15:04:05.345"))
	messageBuilder.WriteString(" " + level + " ")
	messageBuilder.WriteString("[" + serviceName + "," + requestId + "]")
	messageBuilder.WriteString("[" + method + "," + uri + "," + strconv.Itoa(status) + "]")
	messageBuilder.WriteString(" " + message)

	log.Println(messageBuilder.String())
	log.SetFlags(log.Ldate | log.Ltime)
}
