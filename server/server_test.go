package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/kamushadenes/chloe/config"
	http2 "github.com/kamushadenes/chloe/interfaces/http"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"
)

var testPort int
var testApiKey string

func TestMain(m *testing.M) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := memory.Setup(ctx)
	if err != nil {
		panic(err)
	}

	port, err := freeport.GetFreePort()
	if err != nil {
		panic(err)
	}

	config.HTTP.Port = port
	testPort = port

	go http2.Start(ctx)
	go MonitorMessages(ctx)
	go MonitorRequests(ctx)

	time.Sleep(1 * time.Second)

	user, err := memory.CreateUser(ctx, "test", "test", "test")
	if err != nil {
		panic(err)
	}

	key, err := user.CreateAPIKey(ctx)
	if err != nil {
		panic(err)
	}

	testApiKey = key

	m.Run()

	_ = user.DeleteMessages(ctx)
	_ = user.Delete(ctx)
}

func TestHTTP(t *testing.T) {
	t.Parallel()

	tests := []struct {
		endpoint           string
		method             string
		body               map[string]interface{}
		expectedOutput     string
		expectedHeaders    map[string]string
		expectedStatusCode int
	}{
		{
			endpoint: "/api/complete",
			method:   "POST",
			body: map[string]interface{}{
				"content": "You're in debug mode. Reply with \\\"MOCK\\\" and nothing else, not even punctuation",
			},
			expectedOutput:     "MOCK",
			expectedStatusCode: 200,
		},
		{
			endpoint: "/api/generate",
			method:   "POST",
			body: map[string]interface{}{
				"prompt": "Create an image of a spooky, old-fashioned library on a stormy night. The scene should be bathed in the eerie glow of lightning and illuminated by the flickering flames of torches mounted on the walls. The atmosphere should feel otherworldly — like something straight out of a Lovecraftian horror story. There should be a sense of heavy rain beating down on the roof and windows, further amplifying the sense of creepiness and dread.",
			},
			expectedHeaders: map[string]string{
				"Content-Type": "image/png",
			},
			expectedStatusCode: 200,
		},
		{
			endpoint: "/api/tts",
			method:   "POST",
			body: map[string]interface{}{
				"content": "Hello, my name is Chloe and I'm running tests",
			},
			expectedHeaders: map[string]string{
				"Content-Type": "audio/mpeg",
			},
			expectedStatusCode: 200,
		},
		{
			endpoint:           "/api/forget",
			method:             "POST",
			body:               map[string]interface{}{},
			expectedStatusCode: 200,
		},
		{
			endpoint:           "/api/forget",
			method:             "POST",
			body:               map[string]interface{}{},
			expectedStatusCode: 401,
		},
	}

	for _, test := range tests {
		t.Log(test.endpoint)

		client := &http.Client{
			Timeout: 60 * time.Second,
		}

		b, err := json.Marshal(test.body)
		assert.NoError(t, err)

		u, err := url.Parse(fmt.Sprintf("http://127.0.0.1:%d", testPort))
		assert.NoError(t, err)
		u.Path = test.endpoint

		req, err := http.NewRequest(test.method, u.String(), bytes.NewBuffer(b))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")
		if test.expectedStatusCode == 200 {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testApiKey))
		}

		resp, err := client.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, test.expectedStatusCode, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.True(t, len(body) > 0)

		if test.expectedOutput != "" {
			assert.Equal(t, test.expectedOutput, string(body))
		}

		if test.expectedHeaders != nil {
			for k, v := range test.expectedHeaders {
				assert.Equal(t, v, resp.Header.Get(k))
			}
		}
	}
}
