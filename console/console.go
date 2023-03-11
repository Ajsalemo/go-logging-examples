package console

import (
	"encoding/json"
	"log"
	"net/http"
)

func ConsoleLogController(w http.ResponseWriter, r *http.Request) {
	log.Println("This is from the ConsoleLogController..")
	m := map[string]string{"msg": "Logging to console.."}
	e := json.NewEncoder(w)
	e.Encode(m)
}
