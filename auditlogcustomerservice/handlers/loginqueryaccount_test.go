package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginQueryAccount(t *testing.T) {

	w := httptest.NewRecorder()
	jsonBody := []byte(`{
		"Username": "usefulUser4",
		"Password": "randompass"
	}`)
	bodyReader := bytes.NewReader(jsonBody)
	serverPort := 4000

	requestURL := fmt.Sprintf("http://localhost:%d/loginuser", serverPort)
	_, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
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
