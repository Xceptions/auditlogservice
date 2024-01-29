package handlers

import (
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
	}
}
