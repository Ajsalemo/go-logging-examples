package customloggertofile

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func CustomLoggerToFileController(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
    WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime)

    InfoLogger.Println("This is info level logging")
    WarningLogger.Println("This is warning level logging")
    ErrorLogger.Println("This is error level logging")

	m := map[string]string{"msg": "Logging to a file with a custom Go logger.."}
	e := json.NewEncoder(w)
	e.Encode(m)
}
