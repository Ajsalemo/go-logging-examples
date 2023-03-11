package logfile

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func LogFileController(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println("This is from the LogFileController..")
	m := map[string]string{"msg": "Logging to file.."}
	e := json.NewEncoder(w)
	e.Encode(m)
}
