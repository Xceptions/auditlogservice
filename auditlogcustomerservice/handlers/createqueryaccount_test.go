package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateQueryAccount(t *testing.T) {

	w := httptest.NewRecorder()
	jsonBody := []byte(`{
		"Username": "usefulUser4",
		"Email": "randomemail@email.com",
		"Password": "randompass"
	}`)
	bodyReader := bytes.NewReader(jsonBody)
	serverPort := 4000

	requestURL := fmt.Sprintf("http://localhost:%d/createuser", serverPort)
	_, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
}
