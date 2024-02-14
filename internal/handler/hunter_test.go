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
	t.Run("case_01 - success: Prey configured", func(t *testing.T) {
		// Arrange
		// prey
		pr := prey.NewPreyStub()
		// handler
		hd := NewHunter(nil, pr)
		hdFunc := hd.ConfigurePrey()

		// Act
		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(
			`{"speed": 20.0, "position": {"X": 0.0,"Y": 0.0,"Z": 0.0}}`,
		))

		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// Assert
		expectedCode := http.StatusOK
		expectedBody := `{"message": "prey configured", "data": null}`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())

	})
}
