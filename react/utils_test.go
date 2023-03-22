package react

import (
	"errors"
	"testing"
	"time"

	"github.com/kamushadenes/chloe/structs"
	"github.com/stretchr/testify/assert"
)

func TestStartAndWait(t *testing.T) {
	timeout := time.After(1 * time.Second)

	req := structs.CompletionRequest{}

	// Start Channel not nil
	startChannel := make(chan bool)
	req.StartChannel = startChannel
	continueChannel := make(chan bool)
	req.ContinueChannel = continueChannel
	go func() {
		<-startChannel
		continueChannel <- true
	}()

	done := make(chan bool)
	go func() {
		StartAndWait(&req)
		done <- true
	}()

	select {
	case <-timeout:
		t.Fatal("timeout")
	case <-done:
	}
}

func TestNotifyError(t *testing.T) {
	req := structs.CompletionRequest{}
	err := errors.New("some error")

	// Error Channel not nil
	errorChannel := make(chan error)
	req.ErrorChannel = errorChannel
	go func() {
		<-errorChannel
	}()
	assert.Equal(t, err, NotifyError(&req, err))

	// Error Channel nil
	req.ErrorChannel = nil
	assert.Equal(t, err, NotifyError(&req, err))

	assert.Equal(t, nil, NotifyError(&req, nil))
}

func TestWriteResult(t *testing.T) {
	req := structs.CompletionRequest{}
	result := "some result"

	// Result Channel not nil
	resultChannel := make(chan interface{})
	req.ResultChannel = resultChannel
	go func() {
		WriteResult(&req, result)
	}()
	assert.Equal(t, result, <-resultChannel)
}

func TestTruncate(t *testing.T) {
	assert.Equal(t, "abcd", Truncate("abcdefg", 4))
	assert.Equal(t, "abcdefgh", Truncate("abcdefgh", 10))
}
