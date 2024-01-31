package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/xceptions/auditlogservice/auditlogcustomerservice/helpers"
)

func (h handler) PushEventToBuffer(bufferedChannel chan []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		helpers.HandleErr(err)

		// will have to comment out token authentication
		// useful during actual deployment
		// userToken := r.Header.Get("Authorization")

		// helpers.IsAuthorized(w, userToken)

		bufferedChannel <- []byte(body)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"result": "event received!"})
	}
}
