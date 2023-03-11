package zerolog

import (
    "github.com/rs/zerolog/log"
	"encoding/json"
	"net/http"
)

func ZerologController(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello with Zerolog")
	log.Debug().
        Str("Name", "Go").
        Send()

	m := map[string]string{"msg": "Logging to console with Zerolog.."}
	e := json.NewEncoder(w)
	e.Encode(m)
}
