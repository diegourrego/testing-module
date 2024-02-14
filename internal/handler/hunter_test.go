package handler

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testdoubles/internal/prey"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("success - case_01: Prey configured", func(t *testing.T) {
		// Arrange
		// prey
		pr := prey.NewPreyStub()
		// handler
		hd := NewHunter(nil, pr)
		hdFunc := hd.ConfigurePrey()

		// Act
		clientRequest := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(
			`{"speed": 20.0, "position": {"X": 0.0,"Y": 0.0,"Z": 0.0}}`,
		))

		clientRequest.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, clientRequest)

		// Assert
		expectedCode := http.StatusOK
		expectedBody := `{"message": "prey configured", "data": null}`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
	})

	t.Run("failure - case01: Problems with body request configuring hunter", func(t *testing.T) {
		// Arrange
		//handler
		hd := NewHunter(nil, nil)
		FuncConfHunter := hd.ConfigureHunter()

		// Act
		clientRequest := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("Invalid"))
		clientRequest.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		FuncConfHunter(response, clientRequest)

		// Assert
		expectedCode := http.StatusBadRequest
		require.Equal(t, expectedCode, response.Code)
	})
}
