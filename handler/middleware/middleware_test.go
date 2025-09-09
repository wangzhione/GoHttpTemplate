// go
package middleware

import (
	"net/http/httptest"
	"testing"
)

func TestResponseWriterPanicError(t *testing.T) {
	rec := httptest.NewRecorder()
	ResponseWriterPanicError(rec)

	if rec.Code != 589 {
		t.Errorf("expected status code 589, got %d", rec.Code)
	}

	expectedBody := `{"code":"589", "message":"Internal Server Panic Error"}`
	if rec.Body.String() != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, rec.Body.String())
	}
}
