package index

import (
	"encoding/json"
	"net/http"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{"msg": "go-logging-examples"}
	e := json.NewEncoder(w)
	e.Encode(m)
}
