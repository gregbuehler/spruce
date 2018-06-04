package status

import (
	"encoding/json"
	"net/http"
)

var StatusStates = [...]string{
	"ok",     // normal operation
	"frozen", // read only mode - nominally indicates maintenance
	"down",   // system is down. do not expect endpoints to function
}

type StatusPayload struct {
	status string `json:"Status,required" description:""`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := StatusPayload{ status: "ok" }
	json.NewEncoder(w).Encode(p)
}
