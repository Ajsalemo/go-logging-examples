package zap

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

func ZapController(w http.ResponseWriter, r *http.Request) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() 
	sugar := logger.Sugar()
	sugar.Infow("Log contents",
		"name", "Go",
		"logger", "zap",
	)

	m := map[string]string{"msg": "Logging to console with Zap.."}
	e := json.NewEncoder(w)
	e.Encode(m)
}
