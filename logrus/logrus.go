package logrus

import (
	log "github.com/sirupsen/logrus"
	"encoding/json"
	"net/http"
)

func LogrusController(w http.ResponseWriter, r *http.Request) {
	log.Info("Info level logging")
	log.Warn("Warn level logging")
	log.Error("Error level logging")

	m := map[string]string{"msg": "Logging to console with Logrus.."}
	e := json.NewEncoder(w)
	e.Encode(m)
}
