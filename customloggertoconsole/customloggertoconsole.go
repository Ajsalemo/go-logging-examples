package customloggertoconsole

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

func CustomLoggerToConsoleController(w http.ResponseWriter, r *http.Request) {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
    WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
    ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)

    InfoLogger.Println("This is info level logging")
    WarningLogger.Println("This is warning level logging")
    ErrorLogger.Println("This is error level logging")

	m := map[string]string{"msg": "Logging to a console with a custom Go logger.."}
	e := json.NewEncoder(w)
	e.Encode(m)
}
