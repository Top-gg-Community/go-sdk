package dbl

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testToken = "wblAV@d!Od9uL761Rz23BEQC$#YCJdQ0nDlZUEfnDxY"
)

var (
	testPayload  = []byte(`{"bot":"441751906428256277","user":"105122038586286080","type":"upvote","isWeekend":false,"query":""}`)
	testListener = NewListener(testToken, func(p *WebhookPayload) {})
)

func TestHookMethod(t *testing.T) {
	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(testPayload))

	testListener.handlePayload(rec, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code, "GET method should not be allowed")
}

func TestHookAuthentication(t *testing.T) {
	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(testPayload))

	testListener.handlePayload(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code, "Unauthorized request should not be processed")
}

func TestWebhookProcessing(t *testing.T) {
	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(testPayload))
	req.Header.Set("Authorization", testToken)

	testListener.handlePayload(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code, "Request should succeed")
}
