package customloggertoconsoleandfile

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"io"
)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func CustomLoggerToConsoleAndFileController(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("consoleandfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

	InfoLogger = log.New(io.MultiWriter(os.Stdout, file), "INFO: ", log.Ldate|log.Ltime)
    WarningLogger = log.New(io.MultiWriter(os.Stdout, file), "WARNING: ", log.Ldate|log.Ltime)
    ErrorLogger = log.New(io.MultiWriter(os.Stdout, file), "ERROR: ", log.Ldate|log.Ltime)

    InfoLogger.Println("This is info level logging")
    WarningLogger.Println("This is warning level logging")
    ErrorLogger.Println("This is error level logging")

	m := map[string]string{"msg": "Logging to a file and console with a custom Go logger.."}
	e := json.NewEncoder(w)
	e.Encode(m)
}
