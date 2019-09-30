package response

import (
	"encoding/json"
	"net/http"
)

func Send(w http.ResponseWriter, v interface{}, message string, code int) {
	resp, err := json.Marshal(&Response{message, v})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpStatusCode := http.StatusOK
	if code != 0 {
		httpStatusCode = code
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write(resp)
}
