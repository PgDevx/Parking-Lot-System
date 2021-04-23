package wrapper

import (
	"encoding/json"
	"my/v1/model"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func Response(requestCTX RequestContext, w http.ResponseWriter, r *http.Request) {

	requestCTX.RequestID = uuid.NewV4().String() + "-" + uuid.NewV4().String()

	switch t := requestCTX.ResponseType; t {
	case model.HTMLResp:
		w.Header().Set("Content-Type", "text/html")
		w.Write(requestCTX.Response.Payload.([]byte))
	case model.JSONResp:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(requestCTX.Response)
	case model.ErrorResp:
		w.Header().Set("Content-Type", "application/json")
		requestCTX.Err.RequestID = &requestCTX.RequestID
		json.NewEncoder(w).Encode(&requestCTX.Err)
	}
}
