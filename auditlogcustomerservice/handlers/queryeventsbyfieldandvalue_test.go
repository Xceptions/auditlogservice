package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryEventsByFieldAndValue(t *testing.T) {

	w := httptest.NewRecorder()
	serverPort := 4000

	requestURL := fmt.Sprintf("http://localhost:%d/getevents/Customer/John", serverPort)
	_, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Println(err)
	}
	// data, err := io.ReadAll(w.Result().Body)
	// if err != nil {
	// 	t.Errorf("expected error to be nil got %v", err)
	// }

	// fmt.Println(string(data))
	// fmt.Println(reflect.TypeOf(data))

	assert.Equal(t, http.StatusOK, w.Code)
}
