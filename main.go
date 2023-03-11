package main

import (
	"log"
	"net/http"

	console "github.com/Ajsalemo/go-logging-examples/console"
	customLoggerToConsole "github.com/Ajsalemo/go-logging-examples/customloggertoconsole"
	customLoggerToConsoleAndFile "github.com/Ajsalemo/go-logging-examples/customloggertoconsoleandfile"
	customLoggerToFile "github.com/Ajsalemo/go-logging-examples/customloggertofile"
	index "github.com/Ajsalemo/go-logging-examples/index"
	logfile "github.com/Ajsalemo/go-logging-examples/logfile"
	logrus "github.com/Ajsalemo/go-logging-examples/logrus"
	zap "github.com/Ajsalemo/go-logging-examples/zap"
	zerolog "github.com/Ajsalemo/go-logging-examples/zerolog"
)

func main() {
	http.HandleFunc("/", index.IndexController)
	http.HandleFunc("/logger/logfile", logfile.LogFileController)
	http.HandleFunc("/logger/logrus", logrus.LogrusController)
	http.HandleFunc("/logger/zerolog", zerolog.ZerologController)
	http.HandleFunc("/logger/console", console.ConsoleLogController)
	http.HandleFunc("/logger/zap", zap.ZapController)
	http.HandleFunc("/logger/customloggertofile", customLoggerToFile.CustomLoggerToFileController)
	http.HandleFunc("/logger/customloggertoconsole", customLoggerToConsole.CustomLoggerToConsoleController)
	http.HandleFunc("/logger/customloggertoconsoleandfile", customLoggerToConsoleAndFile.CustomLoggerToConsoleAndFileController)

	log.Println("Server listening on port 8080.")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
