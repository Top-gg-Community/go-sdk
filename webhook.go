package dbl

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ListenerFunc func(*WebhookPayload)

type WebhookListener struct {
	token   string
	handler ListenerFunc
	mux     *http.ServeMux
}

type WebhookPayload struct {
	// ID of the bot that received a vote
	Bot string

	// ID of the user who voted
	User string

	// The type of the vote (should always be "upvote" except when using the test button it's "test")
	Type string

	// Whether the weekend multiplier is in effect, meaning users votes count as two
	IsWeekend bool

	// Query string params found on the /bot/:ID/vote page. Example: ?a=1&b=2
	Query url.Values
}

type wPayload struct {
	// ID of the bot that received a vote
	Bot string `json:"bot"`

	// ID of the user who voted
	User string `json:"user"`

	// The type of the vote (should always be "upvote" except when using the test button it's "test")
	Type string `json:"type"`

	// Whether the weekend multiplier is in effect, meaning users votes count as two
	IsWeekend bool `json:"isWeekend"`

	// Query string params found on the /bot/:ID/vote page. Example: ?a=1&b=2
	Query string `json:"query"`
}

// Create a new webhook listener
func NewListener(token string, handler func(*WebhookPayload)) *WebhookListener {
	return &WebhookListener{
		token:   token,
		handler: ListenerFunc(handler),
	}
}

// Starts listening on specific address. A Blocking call.
// Returns non-nil error from ListenAndServe
func (wl *WebhookListener) Serve(addr string) error {
	wl.mux = http.NewServeMux()

	wl.mux.HandleFunc("/", wl.handlePayload)

	return http.ListenAndServe(addr, wl.mux)
}

func (wl *WebhookListener) handlePayload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	if r.Header.Get("Authorization") != wl.token {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return
	}

	p := &wPayload{}

	if err = json.Unmarshal(body, p); err != nil {
		return
	}

	m, err := url.ParseQuery(p.Query)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusNoContent)

	wl.handler(&WebhookPayload{
		Bot:       p.Bot,
		User:      p.User,
		Type:      p.Type,
		IsWeekend: p.IsWeekend,
		Query:     m,
	})
}
