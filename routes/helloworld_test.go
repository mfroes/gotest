package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"mfroes.com/main/routes"
)

type httpResponse struct {
	statusCode int
	body       string
}

func TestHelloWorldRoute(t *testing.T) {
	expected := httpResponse{
		statusCode: http.StatusOK,
		body:       "Hello World",
	}

	// Create a request to pass to our route/handler.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.HelloWorld)

	handler.ServeHTTP(response, req)

	// Check the status code is what we expect.
	if status := response.Code; status != expected.statusCode {
		t.Errorf("returned wrong status code:\n got  [%v]\n want [%v]",
			status, expected.statusCode)
	}

	// Check the response body is what we expect.
	if response.Body.String() != expected.body {
		t.Errorf("returned unexpected body:\n got  [%v]\n want [%v]",
			response.Body.String(), expected.body)
	}
}
